<script lang="ts">
import { Entry, MealTimes } from "./entry.ts";
import type { NutrientStats } from "./entry.ts"
import { defineComponent } from 'vue'
import { getLocalDate } from "../../js/datefn.ts";

type Row = { id?: number, daterecord: string, meal: string, foodname: string, quantity: string, cal: string, protein: string, fat: string, carbs: string, notes: string }

function makeRow(daterecord: string, meal: string, foodname: string, quantity: string, cal: number, protein: number, fat: number, carbs: number, notes: string, id?: number): Row {
    return { daterecord: daterecord, meal: meal, foodname: foodname, quantity: quantity, cal: cal.toFixed(2), protein: protein.toFixed(2), fat: fat.toFixed(2), carbs: carbs.toFixed(2), notes: notes, id: id }
}

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
        sortEntries(first: string, second: string): number {
            return new Date(first).valueOf() - new Date(second).valueOf()
        }
    },
    computed: {
        tabled_entries(): any {
            let entries = this.entries!
            let tabled_entries: any = {}

            if (entries.length > 0) {
                entries.forEach((el) => {
                    // This is extremely important, as it ignores any wonky timestamp stuff and only gives us the date
                    let datestr = el.daterecord.toDateString()
                    if (!(datestr in tabled_entries)) tabled_entries[datestr] = {}
                    if (!(el.meal in tabled_entries[datestr])) tabled_entries[datestr][el.meal] = [] as Array<any>
                    tabled_entries[datestr][el.meal].push(el)
                });
            }
            return tabled_entries
        },
        tabled_totals(): any {
            let tabled_totals: any = {}
            let tabled_entries = this.tabled_entries

            for (const daterecord in tabled_entries) {
                tabled_totals[daterecord] = {}

                let day_totals: NutrientStats = { cal: 0, protein: 0, fat: 0, carbs: 0 }
                for (const meal of MealTimes) {
                    if (meal in tabled_entries[daterecord]) {
                        let meal_totals: NutrientStats = { cal: 0, protein: 0, fat: 0, carbs: 0 }

                        for (const entry of tabled_entries[daterecord][meal]) {
                            meal_totals = this.addStats(meal_totals, entry)
                        }

                        tabled_totals[daterecord][meal] = meal_totals

                        day_totals = this.addStats(day_totals, meal_totals)
                    }
                }
                tabled_totals[daterecord]["Day"] = day_totals;

            }
            return tabled_totals
        },
        formatted_entries(): Row[] {
            let tabled_entries = this.tabled_entries
            let formatted_entries = [] as Row[];
            let sorted_dates = Object.keys(tabled_entries).sort(this.sortEntries)

            for (const daterecord of sorted_dates) {
                let day_entries = tabled_entries[daterecord]
                let day_total = day_entries["totals"]["Day"] as NutrientStats
                formatted_entries.push(makeRow(new Date(daterecord).toDateString(), "Total", "", "", day_total.cal, day_total.protein, day_total.fat, day_total.carbs, ""))
                for (const meal of MealTimes) {
                    if (meal in day_entries) {
                        let meal_entries = day_entries[meal]
                        let meal_total = day_entries["totals"][meal]
                        formatted_entries.push(makeRow("", meal, "", "", meal_total.cal, meal_total.protein, meal_total.fat, meal_total.carbs, ""))
                        for (const entry_index of meal_entries) {
                            const entry = this.entries![entry_index]
                            formatted_entries.push(makeRow("", "", entry.foodname, entry.quantity.toString(), entry.cal, entry.protein, entry.fat, entry.carbs, entry.notes, entry_index))
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
            <template v-for="(meal, daterecord) in tabled_entries" :key="daterecord">
                <tr class="table-border">
                    <td class="font-semibold">{{ daterecord }}</td>
                    <td class="font-semibold">Total</td>
                    <td></td>
                    <td></td>
                    <td class="text-right">{{ tabled_totals[daterecord]["Day"].cal.toFixed(2) }}</td>
                    <td class="text-right">{{ tabled_totals[daterecord]["Day"].protein.toFixed(2) }}</td>
                    <td class="text-right">{{ tabled_totals[daterecord]["Day"].fat.toFixed(2) }}</td>
                    <td class="text-right">{{ tabled_totals[daterecord]["Day"].carbs.toFixed(2) }}</td>
                    <td></td>

                </tr>
                <template v-for="(listEntries, mealname) in meal" :key="mealname">
                    <tr class="table-border">
                        <td></td>
                        <td class="font-semibold"> {{ mealname }}</td>
                        <td></td>
                        <td></td>
                        <td class="text-right">{{ tabled_totals[daterecord][mealname].cal.toFixed(2) }}</td>
                        <td class="text-right">{{ tabled_totals[daterecord][mealname].protein.toFixed(2) }}</td>
                        <td class="text-right">{{ tabled_totals[daterecord][mealname].fat.toFixed(2) }}</td>
                        <td class="text-right">{{ tabled_totals[daterecord][mealname].carbs.toFixed(2) }}</td>
                        <td></td>
                    </tr>
                    <tr v-for="(entry, index) in listEntries" :key="index" class="table-border">
                        <td></td>
                        <td></td>
                        <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{ entry.foodname }}</a>
                        </td>
                        <td class="text-right">{{ entry.quantity }}</td>
                        <td class="text-right">{{ entry.cal }}</td>
                        <td class="text-right">{{ entry.protein }}</td>
                        <td class="text-right">{{ entry.fat }}</td>
                        <td class="text-right">{{ entry.carbs }}</td>
                        <td class="text-right">{{ entry.notes }}</td>
                    </tr>
                </template>
            </template>
            <!-- <tr v-for="(entry, index) in formatted_entries" :key="index" class="table-border">
                <td class="font-semibold">{{ entry.daterecord }}</td>
                <td class="font-semibold">{{ entry.meal }}</td>
                <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{ entry.foodname }}</a></td>
                <td class="text-right">{{ entry.quantity }}</td>
                <td class="text-right">{{ entry.cal }}</td>
                <td class="text-right">{{ entry.protein }}</td>
                <td class="text-right">{{ entry.fat }}</td>
                <td class="text-right">{{ entry.carbs }}</td>
                <td class="text-right">{{ entry.notes }}</td>
            </tr> -->
        </tbody>
    </table>
</template>
