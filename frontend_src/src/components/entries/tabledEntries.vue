<script lang="ts">
import { Entry, MealTimes } from "./entry.ts";
import type { NutrientStats } from "./entry.ts"
import { defineComponent } from 'vue'

export default defineComponent({
    props: {
        entries: Array<Entry>,
        goalinfo: Object
    },
    methods: {
        addStats(totals: NutrientStats, entry: NutrientStats) {
            totals.cal += entry.cal
            totals.protein += entry.protein
            totals.fat += entry.fat
            totals.carbs += entry.carbs
            return totals
        },
        marginErrorPopup(stats: NutrientStats, tolerances: NutrientStats){
            // end result is +/-, amount off the goal
        },
        sortEntries(first: string, second: string): number {
            return new Date(first).valueOf() - new Date(second).valueOf()
        }
    },
    computed: {
        tabled_entries(): any {
            let cloned_entries = this.entries!
            let tabled_entries: any = {}

            if (cloned_entries.length > 0) {

                cloned_entries.forEach((el, index, _) => {
                    let datestr = el.daterecord.toISOString()
                    if (!(datestr in tabled_entries)) tabled_entries[datestr] = {}
                    if (!(el.meal in tabled_entries[datestr])) tabled_entries[datestr][el.meal] = [] as Array<any>
                    tabled_entries[datestr][el.meal].push(index)
                });

                for (const daterecord in tabled_entries) {
                    tabled_entries[daterecord]["totals"] = {}

                    let day_totals = { cal: 0, protein: 0, fat: 0, carbs: 0 }
                    for (const meal of MealTimes) {
                        if (meal in tabled_entries[daterecord]) {
                            let meal_totals = { cal: 0, protein: 0, fat: 0, carbs: 0 }
                            for (const entry_index of tabled_entries[daterecord][meal]) {
                                meal_totals = this.addStats(meal_totals, cloned_entries[entry_index])
                            }

                            tabled_entries[daterecord]["totals"][meal] = meal_totals

                            day_totals = this.addStats(day_totals, meal_totals)
                        }
                    }
                    tabled_entries[daterecord]["totals"]["Day"] = day_totals;
                }
            }
            return tabled_entries
        },
        formatted_entries(): any[] {
            let makeRow = (daterecord: string, meal: string, foodname: string, quantity: string, cal: number, protein: number, fat: number, carbs: number, notes: string) => {
                return { daterecord: daterecord, meal: meal, foodname: foodname, quantity: quantity, cal: cal.toFixed(2), protein: protein.toFixed(2), fat: fat.toFixed(2), carbs: carbs.toFixed(2), notes: notes }
            }

            let formatted_entries: any[] = [];
            let sorted_dates = Object.keys(this.tabled_entries).sort(this.sortEntries)
            
            for (const daterecord of sorted_dates) {
                let day_entries = this.tabled_entries[daterecord]
                let day_total = day_entries["totals"]["Day"] as NutrientStats
                formatted_entries.push(makeRow(new Date(daterecord).toDateString(), "Total", "", "", day_total.cal, day_total.protein, day_total.fat, day_total.carbs, ""))
                for (const meal of MealTimes) {
                    if(meal in day_entries){
                        let meal_entries = day_entries[meal] 
                        let meal_total = day_entries["totals"][meal]
                        formatted_entries.push(makeRow("", meal, "", "", meal_total.cal, meal_total.protein, meal_total.fat, meal_total.carbs, ""))
                        for(const entry_index of meal_entries){
                            const entry = this.entries![entry_index]
                            formatted_entries.push(makeRow("", "", entry.foodname, entry.quantity.toString(), entry.cal, entry.protein, entry.fat, entry.carbs, entry.notes))
                        }
                    }
                    
                }
            }
            return formatted_entries
        }
    }
})
</script>

<template>
        <table class="text table-border table-auto w-full">
        <thead class="module-background table-header">
            <tr>
                <th>Date</th>
                <th>Meal</th>
                <th>Food</th>
                <th>Quantity</th>
                <th>Calorie </th>
                <th>Protein (g)</th>
                <th>Fat (g)</th>
                <th>Carb (g)</th>
                <th>Notes</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(entry, index) in formatted_entries" :key="index" class="table-border">
                <td class="font-semibold">{{ entry.daterecord }}</td>
                <td class="font-semibold">{{ entry.meal }}</td>
                <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{ entry.foodname }}</a></td>
                <td class="text-right">{{ entry.quantity }}</td>
                <td class="text-right">{{ entry.cal }}</td>
                <td class="text-right">{{ entry.protein }}</td>
                <td class="text-right">{{ entry.fat }}</td>
                <td class="text-right">{{ entry.carbs }}</td>
                <td class="text-right">{{ entry.notes }}</td>
            </tr>
        </tbody>
    </table>
</template>

