import { createApp, toRaw } from './vue.esm-browser.prod.js'
import deleteDialog from "../components/confirmDialog.js"
import entriesDialog from '../components/entries-dialog.js'
import sidebaritem from '../components/sidebarItem.js'
import contentHeader from '../components/headerComponent.js'
createApp({
    components: { sidebaritem, contentHeader, deleteDialog, entriesDialog },
    data() {
        return {
            title: "Entries", // Title of this page
            entries: [], // All the entries fetched by GET /api/entries
            start: this.getLastMon(new Date()), //The current displaying week, beginning on Monday and ending on Sunday new Date("2024-09-09T00:00-04:00")
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            showEntriesDialog: false,
            showConfirmDeleteDialog: false,
            selected: null,
            weightRecorded: 0
        }
    },
    computed: {
        inputStartDate : {
            get() {
                return this.getDateOnly(this.currentWeek.start)
            },
            set(dateStr){
                this.currentWeek = new Date(dateStr + "T00:00:00")
            }
        },
        currentWeek: {
            get() {
                // Make a local copy of the start date
                let startLocal = this.getLastMon(this.start)
                let endLocal = this.addDate(startLocal, 6)
                return { start: startLocal, end: endLocal }
            },
            set(startDate) {
                this.start = this.getLastMon(startDate)
            }
        },
        graftTable: {
            get() {
                let graftTable = JSON.parse(JSON.stringify(this.entries))

                for (let index = graftTable.length - 1; index > 0; index--) {
                    let entry = graftTable[index]

                    if (entry.daterecord === graftTable[index - 1].daterecord) {
                        entry.daterecord = ""
                        if (entry.meal === graftTable[index - 1].meal) {
                            entry.meal = ""
                        }
                    } else {
                        entry.daterecord = (new Date(entry.daterecord)).toDateString()
                    }
                }
                let entry = graftTable[0]
                if (entry) {
                    entry.daterecord = (new Date(entry.daterecord)).toDateString()
                }
                return graftTable
            }
        }
    },
    methods: {
        getDateOnly(date) {
            return date.toISOString().split('T')[0]
        },
        getLastMon(date) {
            // Find the difference of dates till the monday, and get the last Monday that has occured in the week. If the day is the same (1==1), then -1*0 is 9, nothing is added.
            const MONDAY = 1
            let diffStartToMonday = date.getDay() - MONDAY
            return this.addDate(date, -1 * diffStartToMonday)
        },
        addDate(date, numDays) {
            // ret is needed for SetDate
            let ret = new Date(date)
            ret.setDate(ret.getDate() + numDays)
            return ret
        },
        sortEntries(first, second) {
            let diffDate = new Date(first.daterecord) - new Date(second.daterecord)
            if (diffDate == 0) {
                return this.mealTimes.indexOf(first.meal) - this.mealTimes.indexOf(second.meal)
            } else {
                return diffDate
            }
        },
        showDeleteDialog(index) {
            this.selected = JSON.parse(JSON.stringify(this.entries[index]))
            this.showConfirmDeleteDialog = true
        },
        showEntriesDialogFn(index = undefined) {
            if (index === undefined) {
                this.selected = {
                    daterecord: this.getDateOnly(this.currentWeek.start), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carb: 0, notes: ""
                }
            } else {
                this.selected = JSON.parse(JSON.stringify(this.entries[index]))
                if(this.entries[index].foodID === undefined) this.selected.foodID = undefined 
                this.selected.daterecord = this.getDateOnly(new Date(this.selected.daterecord))
            }
            this.showEntriesDialog = true
        },
        async editEntry() {
            if (this.selected) {

                let copy_foodID = this.selected.foodID
                this.selected.foodID = { Int32: (this.selected.foodID === undefined ? 0 : this.selected.foodID), Valid: (this.selected.foodID !== undefined) }

                if (!this.selected.hasOwnProperty('id')) {
                    let response = (await fetch("/api/entries/", { method: "POST", body: JSON.stringify(this.selected) }))
                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.foodID = copy_foodID
                        this.selected.id = data.addedID
                        this.entries.push(this.selected)
                    } else {
                        console.log(response.Error)
                    }
                } else {
                    let index = this.entries.findIndex((el) => this.selected.id == el.id)
                    let response = await fetch("/api/entries/" + this.selected.id, { method: "PATCH", body: JSON.stringify(this.selected) })
                    if (response.ok) {
                        this.selected.foodID = copy_foodID
                        this.entries[index] = this.selected
                    } else {
                        let errmsg = await (response.json()).Error
                        console.log(errmsg)
                    }
                }

                this.showEntriesDialog = false

            }
        },
        async editWeight() {
            let response = await fetch("/api/weight/", {
                method: "POST", body:
                    JSON.stringify({ daterecord: this.getDateOnly(this.currentWeek.start), kg: this.weightRecorded })
            })

            if (!response.ok) {
                console.log(response.Error)
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
        async fetchEntries() {
            this.entries = await (await fetch("/api/entries/" + this.getDateOnly(this.currentWeek.start) + "/" + this.getDateOnly(this.currentWeek.end))).json()
            this.entries.forEach((el) => {
                el.foodID = (el.foodID.Valid) ? el.foodID.Int32 : undefined;
                el.daterecord = el.daterecord.split('Z')[0];
            })
            this.entries.sort(this.sortEntries)

            let response = await (await fetch("/api/weight/" + this.getDateOnly(this.currentWeek.start))).json()
            this.weightRecorded = response.kg
        }
    },
    created() {
        this.fetchEntries()
    }

}).mount('#entries')
