<script lang="ts">
import { NutrientStats } from './entry';
import { defineComponent } from 'vue'
import type { PropType } from 'vue'
import type { GoalInfo } from "../settings/goals.ts";

export type Row = { daterecord: string, meal: string, stats: NutrientStats }

export default defineComponent({
    props: {
        goalinfo: { type: Object as PropType<GoalInfo>, required: true },
        total: { type: Object as PropType<Row>, required: true }
    },
    computed: {
        goals(): NutrientStats {
            return this.goalinfo.totals[this.total.meal]
        },
        stats(): NutrientStats {
            return this.total.stats
        },
        error(): NutrientStats {
            return this.goalinfo.marginOfError[this.total.meal]
        },
        diffs(): NutrientStats {
            let diff: NutrientStats = NutrientStats.clone(this.goals)
            diff.subtract(this.stats)
            return diff
        }
    },
    methods: {
        getResult(fieldname: "cal" | "protein" | "fat" | "carbs"): string {
            let fieldval = this.diffs.getField(fieldname)
            if(fieldval > 0) return "bg-green-600/75"
            else if(fieldval >= this.error.getField(fieldname)*-1) return "bg-yellow-600/40"
            else return "bg-red-600/90"
        }
    }
})

</script>

<template>
    <tr class="border-gray-500">
        <td class="font-semibold text-left" v-if="total.daterecord != ''">{{ total.daterecord }}</td>
        <td class="font-semibold text-right" v-else>{{ total.meal }}</td>
        <td></td>
        <td></td>
        <td class="text-right font-semibold">
            <span :class="['totalbox', getResult('cal')] ">{{ Math.round(stats.cal) }} <span class="text-xs hidden md:inline">/ {{ Math.round(goals.cal) }}</span> </span>
        </td>
        <td class="text-right font-semibold">
            <span :class="['totalbox', getResult('protein')] ">{{ Math.round(stats.protein) }} <span class="text-xs hidden md:inline">/ {{ Math.round(goals.protein) }}</span></span>
        </td>
        <td class="text-right font-semibold">
            <span :class="['totalbox', getResult('fat')] ">{{ Math.round(stats.fat) }} <span class="text-xs hidden md:inline">/ {{ Math.round(goals.fat) }}</span></span>
        </td>
        <td class="text-right font-semibold">
            <span :class="['totalbox', getResult('carbs')] ">{{ Math.round(stats.carbs) }} <span class="text-xs hidden md:inline">/ {{ Math.round(goals.carbs) }}</span></span>
        </td>
        <td></td>
    </tr>
</template>

<style scoped>
@import "tailwindcss";

.totalbox {
    @apply text-gray-50 py-0.5 px-2 rounded-4xl shadow-xl;
}

</style>