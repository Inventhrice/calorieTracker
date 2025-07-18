<script>
import entriesDialog from './entries-dialog.vue'
import tabledEntries from './tabledEntries.vue'
import entriesDatePicker from './entriesDatePicker.vue'
import weightEntry from './weightEntry.vue'
import { getLocalDate } from '../../js/datefn.js'
import { api_call, api_get } from '../../js/auth.js'

export default {
    components: { weightEntry, entriesDatePicker, tabledEntries, entriesDialog },
    data() {
        return {
            title: "Entries", // Title of this page
            entries: [], // All the entries fetched by GET /api/entries
            goalsinfo: {},
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            start: null,
            showEntriesDialog: false,
            showConfirmDeleteDialog: false,
            selected: null
        }
    },
    methods: {
        showEntriesDialogFn(index = undefined) {
            if (index === undefined) {
                this.selected = {
                    daterecord: new Date(getLocalDate(undefined) + "T00:00:00"), foodname: "", foodID: undefined,
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
                    let response = await api_call("/api/entries/", "POST", JSON.stringify(selectedCopy))
                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.id = data.addedID
                        this.entries.push(this.selected)
                    } else {
                        console.log(response.Error)
                    }
                } else {
                    let index = this.entries.findIndex((el) => selectedCopy.id == el.id)
                    let response = await api_call("/api/entries/" + selectedCopy.id, "PATCH", JSON.stringify(selectedCopy))
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
                let response = await api_call("/api/entries/" + this.selected.id, "DELETE")
                if (response.ok) {
                    this.entries.splice(this.entries.findIndex((el) => this.selected.id == el.id), 1)
                } else {
                    console.log(response.Error)
                }
                this.showEntriesDialog = false
            }
        },
        async fetchGoalInfo(currentWeek = "") {
            let response = await api_get("/api/goals")
            if (response.ok) {
                let obj = await response.json()
                obj.goalsPerMeal = JSON.parse(obj.goalsPerMeal)
                let returnobj = { percentAllowed: obj.acceptablePercent, Total: obj.goalLbs * obj.multiplier }
                for (let index in this.mealTimes) {
                    let meals = this.mealTimes[index]
                    let mealGoal = obj.goalsPerMeal[index]
                    returnobj[meals] = mealGoal * returnobj.Total
                }
                this.goalsinfo = returnobj
            } else {
                let msg = await response.text()
                console.log(msg)
            }

        },
        async fetchEntries(currentWeek) {
            if (currentWeek) {
                this.fetchGoalInfo(currentWeek)
                this.start = currentWeek.start
                let response = await api_get("/api/entries/" + currentWeek.start + "/" + currentWeek.end)
                if (response.ok) {
                    this.entries = await response.json()
                    this.entries.forEach((el) => {
                        el.foodID = (el.foodID.Valid) ? el.foodID.Int32 : undefined;
                        el.daterecord = new Date((new Date(el.daterecord)).setUTCHours(8));
                    })
                } else {
                    if (response.status === 401) {
                        window.location.href = "/app/login.html"
                    }
                }
            }
        }
    }
}
</script>

<template>
    <div id="main-content" class="content-list">
        <div class="text font-semibold flex flex-col justify-around items-end lg:w-full lg:flex-row lg:justify-between">
            <entries-date-picker class="my-1 p-1" @fetch-entries="fetchEntries"></entries-date-picker>
            <weight-entry class="my-1 p-1" :start></weight-entry>
            <span class="my-1 p-1">
                <button class="flex btn btn-confirm" @click="showEntriesDialogFn(undefined)">
                    <span class="iconify btn-icon text-2xl" data-icon="mdi-pencil-add"></span>
                    <span>Add entry</span>
                </button>
            </span>
        </div>
        <tabled-entries @show-dialog="showEntriesDialogFn" :goalsinfo :entries></tabled-entries>
    </div>
    <entries-dialog v-if="showEntriesDialog" @close-dialog="this.showEntriesDialog = false" :selected
        @confirm-dialog="editEntry" @delete-dialog="deleteEntry"></entries-dialog>
</template>
