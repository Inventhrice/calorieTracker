<script>
import { api_get, api_call } from '../../js/auth';
import EntriesDialog from '../entries/entries-dialog.vue';
import { getLocalDate } from '../../js/datefn.js'

export async function getAllTemplates() {
    let response = await api_get("/api/template/all")
    if (response.ok) {
        let allTemplates = await response.json()
        if(!allTemplates){
            return []
        }

        let allFoods = await fetchAllFoods()
        if(allFoods){
            allTemplates.forEach((el) => el.foodname = allFoods[el.foodID])
        }

        return allTemplates
    } else {
        return []
    }
}

async function fetchAllFoods() {
    const response = await api_get("/api/foodDB/all")
    let allFoods = await response.json()
    let foodDB = {}
    for(let i = 0; i < allFoods.length; i++){
        foodDB[allFoods[i].id] = allFoods[i].name
    }
    return foodDB
}

export default {
    components: { EntriesDialog },
    data() {
        return {
            listTemplates: [],
            selected: {
                    daterecord: new Date(getLocalDate(undefined) + "T00:00:00"), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: "", meal: "Breakfast"
                },
            showTemplateDialog: false
        }
    },
    methods: {
        async fetchData() {
            this.listTemplates = await getAllTemplates()
        },
        async editTemplate() {
            if (this.selected) {
                let index = this.listTemplates.findIndex((el) => this.selected.id == el.id)
                let response = await api_call("/api/template/" + this.selected.id, "PATCH", JSON.stringify(this.selected))
                if (!response.ok) {
                    console.error("Unable to edit the template.")
                }
                this.listTemplates[index] = this.selected
            }
        },
        async deleteTemplate() {
            if (this.selected) {
                let response = await api_call("/api/template/" + this.selected.id, "DELETE")
                if (!response.ok) {
                    console.error("Unable to delete the template.")
                }
                this.listTemplates.splice(this.listTemplates.findIndex((el) => this.selected.id == el.id), 1)
                this.showTemplateDialog = false
            }
        },
        async updateTemplate() {
            if (!this.selected.hasOwnProperty('id')) {
                this.addTemplate()
            } else{
                this.editTemplate()
            }
            this.showTemplateDialog = false
        },
        async addTemplate() {
            if (this.selected) {
                let response = await api_call("/api/template/", "POST", JSON.stringify(this.selected))
                if (!response.ok) {
                    console.error("Unable to create the template")
                }
                let data = await (response).json()
                this.selected.id = data.addedID
                this.listTemplates.push(this.selected)
            }
        },
        async showTemplate(index = undefined) {
            if (index === undefined) {
                this.selected = {
                    daterecord: new Date(getLocalDate(undefined) + "T00:00:00"), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: "", meal: "Breakfast"
                }
            } else {
                this.selected = this.listTemplates[index]
            }
            this.showTemplateDialog = true
        }
    },
    created() {
        this.fetchData()
    }
}

</script>

<template>
    <div
        class="flex flex-col md:flex-row overflow-hidden overflow-y-scroll md:overflow-x-scroll my-3 mr-2 py-3 items-center rounded-2xl module-background w-full">
        <div class="min-w-fit border-blue-500 border-2 m-2 border-dashed rounded-2xl grid grid-cols-1 place-content-center"
            @click="showTemplate()">
            <span class="iconify text-blue-500 hover:text-blue-600 active:text-blue-700 m-3" width="4em" height="4em"
                data-icon="mdi-plus-thick"></span>
        </div>
        <div class="min-w-fit border-2 border-blue-500 rounded-2xl h-full flex flex-col align-middle p-3 mx-2"
            v-for="(templates, index) in listTemplates" :key="index" @click="showTemplate(index)">
            <span class="font-bold">{{ templates.name }}</span>
            <span> Food: {{ templates.foodname }}</span>
            <span> Meal: {{ templates.meal }}</span>
            <span> Quantity: {{ templates.quantity }}</span>
        </div>
    </div>
    <EntriesDialog class="z-10" v-if="showTemplateDialog" @close-dialog="showTemplateDialog = false" @delete-dialog="deleteTemplate"
        @confirm-dialog="updateTemplate" :selected></EntriesDialog>
</template>