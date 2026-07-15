<script lang="ts">
import { Entry, MealTimes } from "./entry.ts";
import { NutrientStats } from "./entry.ts"
import { defineComponent } from 'vue'
import type { GoalInfo } from "../settings/goals.ts";
import type { PropType } from 'vue'
import { default as TotalRow, type Row } from "./tabledEntries-TotalRow.vue";

type TabledEntries = { [daterecord: string]: { [meal: string]: Array<Entry> } }
type TabledTotals = { [daterecord: string]: { [meal: string]: NutrientStats } }

export default defineComponent({
    components: { TotalRow },
    props: {
        entries: Array<Entry>,
        goalinfo: { type: Object as PropType<GoalInfo>, required: true }
    },
    methods: {
        sortEntries(first: string, second: string): number {
            return new Date(first).valueOf() - new Date(second).valueOf()
        },
        getRow(stats: NutrientStats, daterecord: string, meal: string = "Total"): Row {
            let row: Row = {
                daterecord: daterecord,
                meal: meal,
                stats: stats
            }
            return row
        }
    },
    computed: {
        tabled_entries(): TabledEntries {
            let entries = this.entries!
            let tabled_entries: TabledEntries = {}

            if (entries.length > 0) {
                entries.forEach((el) => {
                    // This is extremely important, as it ignores any wonky timestamp stuff and only gives us the date
                    let datestr = el.daterecord.toDateString()

                    if (!(datestr in tabled_entries)) {
                        tabled_entries[datestr] = {}
                        for (const meal of MealTimes) {
                            tabled_entries[datestr][meal] = [] as Array<Entry>
                        }
                    }

                    tabled_entries[datestr][el.meal].push(el)
                });
            }
            return tabled_entries
        },
        tabled_totals(): TabledTotals {
            let tabled_totals: TabledTotals = {}
            let tabled_entries = this.tabled_entries

            for (const daterecord in tabled_entries) {
                tabled_totals[daterecord] = {}

                let day_totals: NutrientStats = new NutrientStats()
                for (const meal of MealTimes) {
                    if (meal in tabled_entries[daterecord]) {
                        let meal_totals: NutrientStats = new NutrientStats()

                        for (const entry of tabled_entries[daterecord][meal]) {
                            meal_totals.add(entry)
                        }

                        tabled_totals[daterecord][meal] = meal_totals

                        day_totals.add(meal_totals)
                    }
                }
                tabled_totals[daterecord]["Day"] = day_totals;

            }
            return tabled_totals
        },
        mealTimes(): string[] {
            return MealTimes
        }
    }
})
</script>

<template>
    <table class="text border-gray-500 border-2 table-auto w-full mr-2">
        <thead class="module-background table-header border-gray-500 border-2">
            <tr>
                <th class="text-left">Date/Meal</th>
                <th class="text-left">Food</th>
                <th class="text-left">Quantity</th>
                <th class="text-right">Calorie </th>
                <th class="text-right">Protein (g)</th>
                <th class="text-right">Fat (g)</th>
                <th class="text-right">Carb (g)</th>
                <th class="">Notes</th>
            </tr>
        </thead>
        <tbody>
            <template v-for="(meal, daterecord) in tabled_entries" :key="daterecord">
                <TotalRow class="bg-blue-300 not-dark:text-blue-800 dark:bg-gray-700" :total="getRow(tabled_totals[daterecord]['Day'], daterecord, 'Total')" :goalinfo="goalinfo">
                </TotalRow>
                <template v-for="mealname in mealTimes" :key="mealname">
                    <TotalRow class="bg-blue-200 dark:bg-gray-800/50" v-if="meal[mealname].length > 0" :total="getRow(tabled_totals[daterecord][mealname], '', mealname)" :goalinfo="goalinfo">
                    </TotalRow>
                    <tr v-for="(entry, index) in meal[mealname]" :key="index" class="bg-blue-100 dark:text-gray-300 dark:bg-gray-900/50">
                        <td></td>
                        <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{ entry.foodname }}</a></td>
                        <td class="">{{ entry.quantity }}</td>
                        <td class="text-right">{{ entry.cal.toFixed(2) }}</td>
                        <td class="text-right">{{ entry.protein.toFixed(2) }}</td>
                        <td class="text-right">{{ entry.fat.toFixed(2) }}</td>
                        <td class="text-right">{{ entry.carbs.toFixed(2) }}</td>
                        <td class="text-right">{{ entry.notes }}</td>
                    </tr>
                </template>
            </template>
        </tbody>
    </table>
</template>