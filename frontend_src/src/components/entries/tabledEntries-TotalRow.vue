<script lang="ts">
import { NutrientStats } from './entry';
import { defineComponent } from 'vue'
import type { PropType } from 'vue'
import type { GoalInfo } from "../settings/goals.ts";

export type Row = { daterecord: string, meal: string, stats: NutrientStats}

export default defineComponent({
    props: {
        goalinfo: {type: Object as PropType<GoalInfo>, required: true},
        total: {type: Object as PropType<Row>, required: true}
    },
    computed: {
        goals(): NutrientStats {
            return this.goalinfo.totals[this.total.meal]
        },
        stats(): any {
            
        }
    }
})

</script>

<template>
    <tr class="table-border">
        <td class="font-semibold"><span v-if="total!.daterecord != ''">{{ total!.daterecord }}</span></td>
        <td class="font-semibold">{{ total!.meal }}</td>
        <td></td>
        <td></td>
        <td class="text-right">{{ Math.round(stats.cal)}} }}</td>
        <td class="text-right">{{ Math.round(stats.protein)}} }}</td>
        <td class="text-right">{{ Math.round(stats.fat) }} </td>
        <td class="text-right">{{ Math.round(stats.carbs) }} </td>
        <td></td>
    </tr>
</template>
