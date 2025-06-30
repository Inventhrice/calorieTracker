export default {
    data() {
        return {
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"]
        }
    },
    computed: {
        graftTable: {
            get() {
                let graftTable = JSON.parse(JSON.stringify(this.entries))
                graftTable.sort(this.sortEntries)
                let add = []
                let totalCal = 0

                for (let index = graftTable.length - 1; index > 0; index--) {
                    let entry = graftTable[index]
                    if (entry.daterecord.valueOf() == graftTable[index-1].daterecord.valueOf()) {
                        totalCal += entry.cal
                        if (entry.meal !== graftTable[index - 1].meal) {
                            add.push({ addIndex: index, data: this.computeStats("", entry.meal, totalCal) })
                            totalCal = 0
                        }
                    } else {
                        add.push({ addIndex: index, data: this.computeStats(new Date(entry.daterecord).toDateString(), entry.meal, totalCal) })
                        totalCal = 0
                    }
                    entry.meal = ""
                    entry.daterecord = ""
                }

                let entry = graftTable[0]
                if (entry) {
                    add.push({ addIndex: 0, data: this.computeStats(new Date(entry.daterecord).toDateString(), entry.meal, totalCal) })
                    entry.meal = ""
                    entry.daterecord = ""
                }

                for(let index = 0; index < add.length; index++){
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
                msg = (diffCal > tolerance) ? "Great job!" :
                    ((diffCal <= tolerance && diffCal >= tolerance * -1) ? "Spot on!" : "Next time!")
            }
            return { daterecord: dateRecord, meal: meal, foodname: "", notes: msg, quantity: totalCal, cal: diffCal, protein: 0, carbs: 0, fat: 0 }

        },
        sortEntries(first, second) {
            let diffDate = first.daterecord - second.daterecord
            if (diffDate == 0) {
                return this.mealTimes.indexOf(first.meal) - this.mealTimes.indexOf(second.meal)
            } else {
                return diffDate
            }
        }
    },
    template: `
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
            <tr v-for="(entry,index) in graftTable" :key="index" class="table-border">
                <td class="font-semibold">{{entry.daterecord}}</td>
                <td class="font-semibold">{{entry.meal}}</td>
                <td class="font-semibold"><a @click="$emit('showDialog', entry.id)">{{entry.foodname}}</a></td>
                <td class="text-right">{{entry.quantity}}</td>
                <td class="text-right">{{entry.cal.toFixed(2)}}</td>
                <td class="text-right">{{entry.protein.toFixed(2)}}</td>
                <td class="text-right">{{entry.fat.toFixed(2)}}</td>
                <td class="text-right">{{entry.carbs.toFixed(2)}}</td>
                <td class="text-right">{{entry.notes}}</td>
            </tr>
        </tbody>
    </table>
    `
}