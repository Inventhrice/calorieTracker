import { createApp, toRaw } from './vue.esm-browser.js'
import deleteDialog from "../components/confirmDialog.js"
import entriesDialog from '../components/entries-dialog.js'
import sidebaritem from '../components/sidebarItem.js'
import contentHeader from '../components/headerComponent.js'
import tabledEntries from '../components/tabledEntries.js'
import entriesDatePicker from '../components/entriesDatePicker.js'
import weightEntry from '../components/weightEntry.js'

createApp({
    components: { sidebaritem, contentHeader, deleteDialog, entriesDialog, tabledEntries, entriesDatePicker, weightEntry },
    data() {
        return {
            title: "Entries", // Title of this page
            entries: [], // All the entries fetched by GET /api/entries
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            start: null,
            showEntriesDialog: false,
            showConfirmDeleteDialog: false,
            selected: null
        }
    },
    methods: {
        showDeleteDialog(index) {
            this.selected = JSON.parse(JSON.stringify(this.entries[index]))
            this.showConfirmDeleteDialog = true
        },
        showEntriesDialogFn(index = undefined) {
            if (index === undefined) {
                this.selected = {
                    daterecord: new Date(), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: ""
                }
            } else {
                this.selected = JSON.parse(JSON.stringify(this.entries[index]))
                if (this.entries[index].foodID === undefined) this.selected.foodID = undefined
                this.selected.daterecord = this.getLocalDate(this.selected.daterecord)
            }
            this.showEntriesDialog = true
        },
        async editEntry() {
            if (this.selected) {
                let selectedCopy = JSON.parse(JSON.stringify(selected))
                selectedCopy.foodID = { Int32: (selectedCopy.foodID === undefined ? 0 : selectedCopy.foodID), Valid: (selectedCopy.foodID !== undefined) }
                selectedCopy.daterecord = this.getLocalDate(selectedCopy.daterecord)

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
            if(currentWeek){
                this.start = currentWeek.start
                this.entries = await (await fetch("/api/entries/" + currentWeek.start + "/" + currentWeek.end)).json()
                this.entries.forEach((el) => {
                    el.foodID = (el.foodID.Valid) ? el.foodID.Int32 : undefined;
                    el.daterecord = new Date((new Date(el.daterecord)).setUTCHours(8));
                })
            }
        }
    }
}).mount('#entries')
