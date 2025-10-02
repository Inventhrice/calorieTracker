<script>
import { getAllTemplates } from "../foodDB/templatedEntries.vue"
export default {
    computed: {
        listTemplates:{
            get(){
                let temp = await getAllTemplates()
                return temp
            }
        }
    }
}

</script>

<template>
    <div class="flex flex-col md:flex-row overflow-hidden overflow-y-scroll md:overflow-x-scroll my-3 mr-2 py-3 items-center rounded-2xl module-background w-full">
        <div class="min-w-fit border-2 border-blue-500 rounded-2xl h-full flex flex-col align-middle p-3 mx-2" v-for="(templates, index) in listTemplates" :key="index" @click="$emit('showDialog', this.listTemplates[index])">
            <span class="font-bold">{{templates.name}}</span>
            <span> Food: {{templates.foodName}}</span>
            <span> Meal: {{ templates.meal }}</span>
            <span> Quantity: {{ templates.quantity }}</span>
        </div>
    </div>
</template>