<script>
import tabledEntries from './tabledEntries.vue'
import entriesDatePicker from './entriesDatePicker.vue'
import weightEntry from './weightEntry.vue'

import entriesDialog from '../template/entries-dialog.vue'
import listTemplates from '../template/listTemplates.vue'

import { getLocalDate, getToday } from '../../js/datefn.ts'
import { clone, api_call, api_get } from '../../js/api.js'
import { Entry } from "./entry.ts";
import { defineComponent } from 'vue'
import { parse_goals } from '../settings/goals.ts'

export default defineComponent({
    components: { weightEntry, entriesDatePicker, tabledEntries, entriesDialog, listTemplates },
    data() {
        return {
            title: "Entries", // Title of this page
            entries: [], // All the entries fetched by GET /api/entries
            goalinfo: {},
            start: null,
            showEntriesDialog: false,
            showConfirmDeleteDialog: false,
            selected: new Entry()
        }
    },
    methods: {
        showEntriesDialogFn(index = undefined) {
            if (index === undefined) {
                this.selected = new Entry();
            } else {
                let found = this.entries.find((el) => el.id == index)
                this.selected = clone(found)
                if (found.foodID === undefined) this.selected.foodID = undefined
                this.selected.daterecord = new Date(this.selected.daterecord)
            }
            this.showEntriesDialog = true
        },
        makeTemplateEntry(selected) {
            this.selected = clone(selected)
            delete this.selected['id']
            this.selected.daterecord = getToday();
            this.showEntriesDialog = true
        },
        async editEntry() {
            if (this.selected) {
                let selectedCopy = clone(this.selected)
                selectedCopy.foodID = { Int32: (selectedCopy.foodID === undefined ? 0 : selectedCopy.foodID), Valid: (selectedCopy.foodID !== undefined) }
                selectedCopy.daterecord = getLocalDate(new Date(selectedCopy.daterecord))

                // This checks if this is NOT an pre-existing entry
                if (!this.selected.hasOwnProperty('id')) {
                    let response = await api_call("/api/entries/", "POST", JSON.stringify(selectedCopy))

                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.id = data.addedID
                        this.entries.push(this.selected)
                    } else {
                        console.log((response).Error)
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
                    console.log((response).Error)
                }
                this.showEntriesDialog = false
            }
        },
        async fetchGoalInfo() {
            let response = await api_get("/api/goals")
            if (response.ok) {
                let obj = await response.json()
                this.goalinfo = parse_goals(obj)
            } else {
                let msg = await response.text()
                console.log(msg)
            }

        },
        async fetchEntries(currentWeek) {
            if (currentWeek) {
                this.fetchGoalInfo()
                this.start = currentWeek.start
                let response = await api_get("/api/entries/" + currentWeek.start + "/" + currentWeek.end)
                if (response.ok) {
                    this.entries = []
                    let rawEntries = await response.json()
                    for(let el of rawEntries) {
                        el.daterecord = new Date((new Date(el.daterecord)).setUTCHours(8));
                        this.entries.push(Entry.from(el))
                        
                    }
                }
            }
        }
    }
})
</script>

<template>
    <div id="main-content" class="content-list">
        <div class="text font-semibold flex flex-col justify-around items-end lg:w-full lg:flex-row lg:justify-between">
            <entries-date-picker class="my-1 p-1" @fetch-entries="fetchEntries"></entries-date-picker>
            <weight-entry class="my-1 p-1" :start></weight-entry>
            <span class="my-1 p-1">
                <button class="flex btn btn-confirm" @click="showEntriesDialogFn(undefined)">
                    <span class="icon mdi--pencil-add btn-icon text-2xl"></span>
                    <span>Add entry</span>
                </button>
            </span>
        </div>
        <div class="flex w-full">
            <list-templates @show-dialog="makeTemplateEntry"></list-templates>
        </div>
        <div class="flex w-full">
            <tabled-entries @show-dialog="showEntriesDialogFn" :goalinfo="goalinfo" :entries></tabled-entries>
        </div>
    </div>
    <entries-dialog v-if="showEntriesDialog" @close-dialog="this.showEntriesDialog = false" :selected
        @confirm-dialog="editEntry" @delete-dialog="deleteEntry"></entries-dialog>
</template>
