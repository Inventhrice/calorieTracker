<script>
import { api_get } from '../../js/auth';
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