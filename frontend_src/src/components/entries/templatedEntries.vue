<script>
import { api_get } from '../../js/auth';
import entriesDialog from './entries-dialog.vue';
export default {
    components: {entriesDialog},
    data() {
        return {
            listTemplates: [],
            selected: {},
            showEntriesDialog: false
        }
    },
    methods: {
        async fetchData() {

        },
        async editTemplate() {

        },
        async deleteTemplate(){

        },
        async showTemplate(index = undefined){
            if (index === undefined){
                this.selected = {daterecord: new Date(getLocalDate(undefined) + "T00:00:00"), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: ""}
            } else{
                this.selected = this.listTemplates[index]
            }
            this.showEntriesDialog = true
        }
    },
    created() {
        this.fetchData()
    }
}

</script>

<template>
    <div class="flex flex-col overflow-y-auto md:flex-row md:overflow-x-auto my-3 mr-2 rounded-2xl module-background w-full">
        <div class="border-blue-500 border-2 m-2 border-dashed rounded-2xl grid grid-cols-1 place-content-center">
            <span class="iconify text-blue-500 hover:text-blue-600 active:text-blue-700 m-3" width="4em" height="4em" data-icon="mdi-plus-thick"></span>
        </div>
        <div class="border-2 rounded-2xl h-full flex flex-col" v-for="(templates, index) in listTemplates" :key="index" @click="showTemplate(index)">
            <span class="text-bold">{{templates.name}}</span>
            <span> Food: {{templates.foodName}}</span>
            <span> Mealtime: {{ templates.meal }}</span>
            <span> Quantity: {{ templates.quantity }}</span>
        </div>
    </div>
    <entries-dialog v-if="showEntriesDialog" @close-dialog="this.showEntriesDialog = false" :selected
        @confirm-dialog="editTemplate" @delete-dialog="deleteTemplate"></entries-dialog>
</template>