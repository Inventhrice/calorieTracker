export default {
    data() {
        return {
            editFields: false,
            allFoods: [],
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"]
        }
    },
    computed: {
        isFoodID: {
            get() {
                return this.selected.foodID !== undefined
            }
        },
        foodinfo: {
            get() {
                return this.allFoods.find((el) => el.id == this.selected.foodID)
            }
        },
        dateRecord: {
            get(){
                return this.selected.daterecord.toISOString().split('T')[0]
            },
            set(dateStr){
                this.selected.daterecord = new Date(dateStr+"T00:00:00")
            }
        }
    },
    props: {
        selected: Object
    },
    methods: {
        async fetchAllFoods() {
            const response = await fetch("/api/foodDB/all")
            this.allFoods = await response.json()
        },
        updateValues(quantity) {
            if (this.isFoodID) {
                let foodinfo = this.foodinfo
                this.selected.foodname = foodinfo.name
                this.selected.cal = foodinfo.calperg * quantity
                this.selected.carbs = foodinfo.carbperg * quantity
                this.selected.protein = foodinfo.proteinperg * quantity
                this.selected.fat = foodinfo.fatperg * quantity
            }

        }
    },
    created() {
        this.$watch('isFoodID', (isFoodID) => {
            if (isFoodID) {
                this.updateValues(this.selected.quantity)
            }
        })

        this.$watch('selected.quantity', this.updateValues)

        this.$watch('editFields', (editFields) => {
            if (editFields) {
                this.selected.foodID = undefined
            }
        })

        this.fetchAllFoods()

    },
    template: `
        <dialog class="dialog text flex flex-col" open>
            <label for="date">Date</label>
            <input class="dialog-input" type="date" v-model="dateRecord" />

            <label for="foodName">Food</label>
            <span class="grid grid-cols-2 gap-3">
            <select class="dialog-input" v-model="selected.foodID">
                    <option :value=undefined>Select</option>
                    <option v-for="food in allFoods" :value=food.id> {{food.name}} </option>
            </select>
            <input class="dialog-input" v-model="selected.foodname" v-if="editFields || !isFoodID"/>
            </span>
            
            <div><label for="meal">Meal</label>
            <select class="dialog-input" name="meal" v-model="selected.meal">
                <option v-for="meal in this.mealTimes"> {{meal}} </option>
            </select></div>

            <label for="grams">Quantity</label>
            <input class="dialog-input" type="number" v-model="selected.quantity" />

            <span v-if="isFoodID">
                <input type="checkbox" name="editFields" v-model="editFields" />
                <label for="editFields">Override the values?</label>
            </span>

            <div class="grid grid-cols-2 gap-3">
                <label for="cal">Calorie (g)</label>
                <label for="protein">Protein (g)</label>
                <input type="number" class="dialog-input" step="0.01" v-model="selected.cal" :disabled="isFoodID"/>
                <input type="number" class="dialog-input" step="0.01" id="protein" v-model="selected.protein" :disabled="isFoodID"/>
                <label for="fat">Fat (g)</label>
                <label for="carb">Carbs (g)</label>
                <input type="number" class="dialog-input" step="0.01" id="fat" v-model="selected.fat" :disabled="isFoodID"/>
                <input type="number" class="dialog-input" step="0.01" id="carb" v-model="selected.carbs" :disabled="isFoodID"/>
            </div>

            <label for="notes">Notes</label>
            <textarea id="notes" class="dialog-input" v-model="selected.notes"></textarea>

            <div class="flex justify-end">
                <button class="btn" @click="$emit('close-dialog')">Cancel</button>
                <button class="btn btn-uhoh px-1" @click="$emit('delete-dialog')">Delete</button>
                <button class="btn btn-confirm" @click="$emit('confirm-dialog')">Save</button>
            </div>
        </dialog>
    `
}