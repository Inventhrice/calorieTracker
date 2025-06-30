export default {
    computed: {
        graftTable: {
            get() {
                let graftTable = JSON.parse(JSON.stringify(this.entries))
                graftTable.sort(this.sortEntries)
                
                let stats = {}
                let totalCal = 0
                for (let index = graftTable.length - 1; index > 0; index--) {
                    let entry = graftTable[index]
                    if (entry.daterecord === graftTable[index - 1].daterecord) {
                        totalcal += entry.cal
                        if (entry.meal !== graftTable[index - 1].meal) {
                            stats = this.computeStats(totalCal, entry.meal)
                            graftTable.splice(index,0, {daterecord: "", meal: entry.meal, foodname: msg, quantity: totalCal, cal: diffCal, protein: 0, carbs: 0, fat: 0})
                            totalCal = 0
                        }
                        entry.meal = ""
                    } else {
                        graftTable.splice(index,0, {daterecord: new Date(entry.daterecord).toDateString(), meal: entry.meal, foodname: msg, quantity: totalCal, cal: diffCal, protein: 0, carbs: 0, fat: 0})
                    }
                    entry.daterecord = ""
                }
                let entry = graftTable[0]
                if (entry) {                    
                    graftTable.splice(0,0, {daterecord: new Date(entry.daterecord).toDateString(), meal: entry.meal, foodname: msg, quantity: totalCal, cal: diffCal, protein: 0, carbs: 0, fat: 0})
                    entry.daterecord = ""
                    entry.meal = ""
                }
                return graftTable
            }
        }
    },
    props: {
        entries: Array
    },
    methods: {
        computeStats(totalCalForMeal, meal) {
            let diffCal = (goals[meal]-totalCalForMeal)
            let tolerance = goals[meal]*(goals.percentAllowed)
            let msg = (diffCal > tolerance) ? "Great job!" : 
                    ((diffCal <= tolerance && diffCal >= tolerance*-1) ? "Spot on!" : "Next time!")
            return {msg: msg, totalCal: totalCalForMeal, diffCal: diffCal}
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
            </tr>
        </thead>
        <tbody>
            <tr v-for="(entry,index) in graftTable" :key="index" class="table-border">
                <td class="font-semibold">{{entry.daterecord}}</td>
                <td class="font-semibold">{{entry.meal}}</td>
                <td class="font-semibold"><a @click="$emit('showDialog', index)">{{entry.foodname}}</a></td>
                <td class="text-right">{{entry.quantity}}</td>
                <td class="text-right">{{entry.cal.toFixed(2)}}</td>
                <td class="text-right">{{entry.protein.toFixed(2)}}</td>
                <td class="text-right">{{entry.fat.toFixed(2)}}</td>
                <td class="text-right">{{entry.carbs.toFixed(2)}}</td>
            </tr>
        </tbody>
    </table>
    `
}