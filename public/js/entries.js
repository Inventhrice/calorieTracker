import { createApp, toRaw } from './vue.esm-browser.js'
import deleteDialog from "../components/confirmDialog.js"
import entriesDialog from '../components/entries-dialog.js'
import sidebaritem from '../components/sidebarItem.js'
import contentHeader from '../components/headerComponent.js'
import tabledEntries from '../components/tabledEntries.js'
import entriesDatePicker from '../components/entriesDatePicker.js'
import weightEntry from '../components/weightEntry.js'
import { getLocalDate } from '../components/dateFunctions.js'
createApp({
    components: { sidebaritem, contentHeader, deleteDialog, entriesDialog, tabledEntries, entriesDatePicker, weightEntry },
    data() {
        return {
            title: "Entries", // Title of this page
            entries: [], // All the entries fetched by GET /api/entries
            goalsinfo: {percentAllowed: 0.1, Total: 1900, Breakfast: 380, Lunch: 665, Dinner: 665, Snacks: 190},
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            start: null,
            showEntriesDialog: false,
            showConfirmDeleteDialog: false,
            selected: null
        }
    },
    methods: {
        showDeleteDialog(select) {
			this.selected = select
			this.showEntriesDialog = false
			this.deleteEntry()
        },
        showEntriesDialogFn(index = undefined) {
            if (index === undefined) {
                this.selected = {
                    daterecord:  new Date(getLocalDate(undefined)+"T00:00:00"), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: ""
                }
            } else {
                let found = this.entries.find((el) => el.id == index)
                this.selected = JSON.parse(JSON.stringify(found))
                if (found.foodID === undefined) this.selected.foodID = undefined
                this.selected.daterecord = new Date(this.selected.daterecord)
            }
            this.showEntriesDialog = true
        },
        async editEntry() {
            if (this.selected) {
                let selectedCopy = JSON.parse(JSON.stringify(this.selected))
                selectedCopy.foodID = { Int32: (selectedCopy.foodID === undefined ? 0 : selectedCopy.foodID), Valid: (selectedCopy.foodID !== undefined) }
                selectedCopy.daterecord = getLocalDate(new Date(selectedCopy.daterecord))

                if (!this.selected.hasOwnProperty('id')) {
                    let response = (await fetch("/api/entries/", { method: "POST", body: JSON.stringify(selectedCopy) }))
                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.id = data.addedID
                        this.entries.push(this.selected)
                    } else {
                        console.log(response.Error)
                    }
                } else {
                    let index = this.entries.findIndex((el) => selectedCopy.id == el.id)
                    let response = await fetch("/api/entries/" + selectedCopy.id, { method: "PATCH", body: JSON.stringify(selectedCopy) })
                    if (response.ok) {
                        this.entries[index] = this.selected
                    } else {
                        let errmsg = await (response.json()).Error
                        console.log(errmsg)
                    }
                }

                this.showEntriesDialog = false

            }
        },
        async deleteEntry() {
            if (this.selected) {
                let response = await fetch("/api/entries/" + this.selected.id, { method: "DELETE" })
                if (response.ok) {
                    this.entries.splice(this.entries.findIndex((el) => this.selected.id == el.id), 1)
                } else {
                    console.log(response.Error)
                }
                this.showConfirmDeleteDialog = false
            }
        },
        async fetchEntries(currentWeek) {
            //this.goalsInfo = await(await fetch("/api/settings/goals")).json()
            if(currentWeek){
                this.start = currentWeek.start
				let response = await fetch("/api/entries/" + currentWeek.start + "/" + currentWeek.end)
				if(response.ok){
					this.entries = await response.json()
					this.entries.forEach((el) => {
						el.foodID = (el.foodID.Valid) ? el.foodID.Int32 : undefined;
						el.daterecord = new Date((new Date(el.daterecord)).setUTCHours(8));
					})
				} else{
					if(response.status === 401){
						window.location.href = "/app/login.html"
					}
				}
            }
        }
    }
}).mount('#entries')
