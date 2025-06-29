export default{
    computed: {
        graftTable: {
            get() {
                let graftTable = JSON.parse(JSON.stringify(this.entries))
                graftTable.sort(this.sortEntries)

                for (let index = graftTable.length - 1; index > 0; index--) {
                    let entry = graftTable[index]

                    if (entry.daterecord === graftTable[index - 1].daterecord) {
                        entry.daterecord = ""
                        if (entry.meal === graftTable[index - 1].meal) {
                            entry.meal = ""
                        }
                    } else {
                        entry.daterecord = new Date(entry.daterecord).toDateString()
                        
                    }
                }
                let entry = graftTable[0]
                if (entry) {
                    entry.daterecord = new Date(entry.daterecord).toDateString()
                }
                return graftTable
            }
        }
    },
    props:{
        entries: Array
    },
    methods: {
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
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(entry,index) in graftTable" :key="index" class="table-border">
                <td class="font-semibold">{{entry.daterecord}}</td>
                <td class="font-semibold">{{entry.meal}}</td>
                <td class="font-semibold">{{entry.foodname}}</td>
                <td class="text-right">{{entry.quantity}}</td>
                <td class="text-right">{{entry.cal.toFixed(2)}}</td>
                <td class="text-right">{{entry.protein.toFixed(2)}}</td>
                <td class="text-right">{{entry.fat.toFixed(2)}}</td>
                <td class="text-right">{{entry.carbs.toFixed(2)}}</td>
                <td class="text-center py-1">
                    <span class="p-1">
                        <button class="btn" @click="showEntriesDialogFn(index)"> <span
                                class="iconify btn-icon" data-icon="mdi-pencil"></span> </button>
                    </span>
                    <span class="p-1">
                        <button class="btn btn-uhoh px-1" @click="showDeleteDialog(index)"> <span
                                class="iconify btn-icon" data-icon="mdi-trash-can"></span> </button>
                    </span>
                </td>
            </tr>
        </tbody>
    </table>
    `
}