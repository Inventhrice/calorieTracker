<script>
import { api_get } from '../../js/auth';

export async function getAllTemplates(){
    let response = await api_get("/api/templates")
    if(response.ok){
        let allTemplates = await response.json()
        return allTemplates
    } else{
        return []
    }
}

export default {
    data() {
        return {
            listTemplates: [
                {name: "Tea (Breakfast)", foodName: "Tea", quantity: 1, meal: "Breakfast"},
                {name: "Tea (Dinner)", foodName: "Tea", quantity: 1, meal: "Dinner"},
                {name: "Rice", foodName: "Rice", quantity: 250, meal: "Lunch"},
                {name: "Tea (Breakfast)", foodName: "Tea", quantity: 1, meal: "Breakfast"},
                {name: "Tea (Dinner)", foodName: "Tea", quantity: 1, meal: "Dinner"},
                {name: "Rice", foodName: "Rice", quantity: 250, meal: "Lunch"},
                {name: "Tea (Breakfast)", foodName: "Tea", quantity: 1, meal: "Breakfast"},
                {name: "Tea (Dinner)", foodName: "Tea", quantity: 1, meal: "Dinner"},
                {name: "Rice", foodName: "Rice", quantity: 250, meal: "Lunch"},
                {name: "Tea (Breakfast)", foodName: "Tea", quantity: 1, meal: "Breakfast"},
                {name: "Tea (Dinner)", foodName: "Tea", quantity: 1, meal: "Dinner"},
                {name: "Rice", foodName: "Rice", quantity: 250, meal: "Lunch"}
            ],
            selected: {},
            showEntriesDialog: false
        }
    },
    methods: {
        async fetchData() {
            this.listTemplates = await getAllTemplates()
        },
        async editTemplate() {
            let response = await api_call("/api/template/" + this.selected.id, "PATCH", this.selected)
            if(!response.ok){
                console.error("Unable to edit the template.")
            }
            this.listTemplates[this.selected.id] = this.selected
        },
        async deleteTemplate(){
            let response = await api_call("/api/template/" + this.selected.id, "DELETE")
            if(!response.ok){
                console.error("Unable to delete the template.")
            }
            this.listTemplates.splice(this.listTemplates.findIndex((el) => this.selected.id == el.id), 1)
        },
        async addTemplate(){
            let response = await api_call("/api/template/", "POST", this.selected)
            if(!response.ok){
                console.error("Unable to create the template")
            }
            this.listTemplates.push(this.selected)
        },
        async showTemplate(index = undefined){
            if (index === undefined){
                this.selected = {daterecord: new Date(getLocalDate(undefined) + "T00:00:00"), foodname: "", foodID: undefined,
                    quantity: 0, cal: 0, protein: 0, fat: 0, carbs: 0, notes: "", meal: "Breakfast"}
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
    <div class="flex flex-col md:flex-row overflow-hidden overflow-y-scroll md:overflow-x-scroll my-3 mr-2 py-3 items-center rounded-2xl module-background w-full">
        <div class="min-w-fit border-blue-500 border-2 m-2 border-dashed rounded-2xl grid grid-cols-1 place-content-center">
            <span class="iconify text-blue-500 hover:text-blue-600 active:text-blue-700 m-3" width="4em" height="4em" data-icon="mdi-plus-thick"></span>
        </div>
        <div class="min-w-fit border-2 border-blue-500 rounded-2xl h-full flex flex-col align-middle p-3 mx-2" v-for="(templates, index) in listTemplates" :key="index" @click="showTemplate(index)">
            <span class="font-bold">{{templates.name}}</span>
            <span> Food: {{templates.foodName}}</span>
            <span> Meal: {{ templates.meal }}</span>
            <span> Quantity: {{ templates.quantity }}</span>
        </div>
    </div>
</template>