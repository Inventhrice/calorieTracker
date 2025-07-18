<script>
export default {
    data() {
        return {
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"]
        }
    },
    computed: {
        graftTable: {
            get() {
                let graftTable = JSON.parse(JSON.stringify(this.entries)).toSorted(this.sortEntries)
                let add = []
                let caloriesOfDay = 0
                let totalCal = 0

                for (let index = graftTable.length - 1; index > 0; index--) {
                    let entry = graftTable[index]
                    totalCal += entry.cal
                    if (entry.daterecord.valueOf() == graftTable[index - 1].daterecord.valueOf()) {
                        if (entry.meal !== graftTable[index - 1].meal) {
                            add.push({ addIndex: index, data: this.computeStats("", entry.meal, totalCal) })
                            caloriesOfDay += totalCal
                            totalCal = 0
                        }
                    } else {
                        add.push({ addIndex: index, data: this.computeStats("", entry.meal, totalCal) })

                        caloriesOfDay += totalCal
                        add.push({ addIndex: index, data: this.computeStats(new Date(entry.daterecord).toDateString(), "Total", caloriesOfDay) })

                        totalCal = 0
                        caloriesOfDay = 0
                    }
                    entry.meal = ""
                    entry.daterecord = ""
                }

                let entry = graftTable[0]
                if (entry) {
                    totalCal += entry.cal
                    add.push({ addIndex: 0, data: this.computeStats("", entry.meal, totalCal) })
                    caloriesOfDay += totalCal
                    add.push({ addIndex: 0, data: this.computeStats(new Date(entry.daterecord).toDateString(), "Total", caloriesOfDay) })

                    entry.meal = ""
                    entry.daterecord = ""
                }

                for (let index = 0; index < add.length; index++) {
                    let data = add[index]
                    graftTable.splice(data.addIndex, 0, data.data)
                }

                return graftTable
            }
        }
    },
    props: {
        entries: Array,
        goalsinfo: Object
    },
    methods: {
        computeStats(dateRecord = "", meal, totalCal) {
            let msg = ""
            let diffCal = 0
            if (this.goalsinfo) {
                let goals = this.goalsinfo
                diffCal = (goals[meal] - totalCal)
                let tolerance = goals[meal] * (goals.percentAllowed)
                msg = (diffCal >= 0) ? "Great job!" :
                    ((diffCal >= tolerance * -1) ? "Spot on!" : "Next time!")
            }
            return { daterecord: dateRecord, meal: meal, foodname: "", notes: msg, quantity: totalCal, cal: diffCal, protein: 0, carbs: 0, fat: 0 }

        },
        sortEntries(first, second) {
            let diffDate = new Date(first.daterecord).valueOf() - new Date(second.daterecord).valueOf()
            if (diffDate == 0) {
                if (this.mealTimes) {
                    return this.mealTimes.indexOf(first.meal) - this.mealTimes.indexOf(second.meal)
                } else {
                    console.log("Unable to sort, this.mealtimes not defined")
                    return -1
                }
            } else {
                return diffDate
            }
        }
    }
}
</script>

<template>
    <table class="text table-border table-auto w-full">
        <thead class="module-background table-header">
            <tr>
                <th>Date</th>
                <th>Meal</th>
                <th>Food</th>
                <th>Quantity</th>
                <th>Calorie (g)</th>
                <th>Protein (g)</th>
                <th>Fat (g)</th>
                <th>Carb (g)</th>
                <th>Notes</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(entry, index) in graftTable" :key="index" class="table-border">
                <td class="font-semibold">{{ entry.daterecord }}</td>
                <td class="font-semibold">{{ entry.meal }}</td>
                <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{ entry.foodname }}</a></td>
                <td class="text-right">{{ entry.quantity }}</td>
                <td class="text-right">{{ entry.cal.toFixed(2) }}</td>
                <td class="text-right">{{ entry.protein.toFixed(2) }}</td>
                <td class="text-right">{{ entry.fat.toFixed(2) }}</td>
                <td class="text-right">{{ entry.carbs.toFixed(2) }}</td>
                <td class="text-right">{{ entry.notes }}</td>
            </tr>
        </tbody>
    </table>
</template>
