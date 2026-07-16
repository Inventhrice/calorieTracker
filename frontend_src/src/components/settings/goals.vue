<script lang="ts">
import { api_call, api_get } from "../../js/api.ts"
import { defineComponent } from 'vue';
import { calc_day_goals } from "./goals.ts";

export default defineComponent({
    data() {
        return {
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            userGoals: { "Weight": 84.5, "% Error": 15, "Multiplier": 10, "proteinGPerLBS": 0.8, "fatGPerLBS": 0.4 },
            mealGoals: { "Breakfast": 0, "Lunch": 0, "Dinner": 0, "Snacks": 0 } as any
        }
    },
    computed: {
        weightKgs: {
            get(): string {
                return this.userGoals['Weight'].toFixed(2)
            },
            set(newVal: string) {
                this.userGoals['Weight'] = parseFloat(newVal)
            }
        },
        weightLbs(): number {
            return Math.round(this.userGoals['Weight'] * 2.2)
        },
        calErrorMargin(): number {
            return (this.userGoals['% Error'] * this.totalMacros["Calories"] / 100)
        },
        totalMacros(): any {
            let goalinfo = {
                multiplier: this.userGoals['Multiplier'],
                goalLbs: this.weightLbs,
                proteinGPerLBS: this.userGoals["proteinGPerLBS"],
                fatGPerLBS: this.userGoals["fatGPerLBS"]
            }

            let calc_goals = calc_day_goals(goalinfo)

            let tMacros = { "Calories": 0, "Protein": 0, "Fat": 0, "Carbs": 0 }
            tMacros["Calories"] = calc_goals.cal
            tMacros["Protein"] = calc_goals.protein
            tMacros["Fat"] = calc_goals.fat
            tMacros["Carbs"] = calc_goals.carbs
            return tMacros
        },
        mealsAllottedAddsUp(): boolean {
            let add = 0
            for (let meal in this.mealGoals) {
                add += this.mealGoals[meal]
            }
            return add == 100
        }
    },
    methods: {
        async fetchPreferences() {
            let response = await api_get("/api/goals")
            if (response.ok) {
                let userData = await response.json()
                this.userGoals["Weight"] = (userData.goalLbs / 2.2)
                this.userGoals["Multiplier"] = userData.multiplier
                this.userGoals["% Error"] = userData.acceptablePercent * 100
                this.userGoals["proteinGPerLBS"] = userData["proteinGPerLBS"]
                this.userGoals["fatGPerLBS"] = userData["fatGPerLBS"]
                for (let index in this.mealTimes) {
                    this.mealGoals[this.mealTimes[index]] = JSON.parse(userData.goalsPerMeal)[index] * 100
                }
            }
        },
        async fetchData() {
            let response = await api_get("/api/profile")
            if (response.ok) {
                await this.fetchPreferences()
            } else {
                console.error("Not signed in.")
            }
        },
        async updateGoals() {
            if (this.mealsAllottedAddsUp) {
                let temp = []
                for (let meal in this.mealGoals) {
                    temp.push(this.mealGoals[meal] / 100)
                }

                let obj = {
                    "goalLbs": this.weightLbs,
                    "multiplier": this.userGoals["Multiplier"],
                    "acceptablePercent": this.userGoals["% Error"] / 100,
                    "goalsPerMeal": JSON.stringify(temp),
                    "proteinGPerLBS": this.userGoals["proteinGPerLBS"],
                    "fatGPerLBS": this.userGoals["fatGPerLBS"]
                }

                let response = await api_call("/api/goals", "POST", JSON.stringify(obj))
                if (!response.ok) {
                    console.error("POST request for updating goals failed.")
                }
            } else {
                console.error("Cannot save goals, meals do not add up.")
            }

        }
    },
    created() {
        this.fetchData()
    }
})

</script>

<template>
    <div class="grid grid-cols-2 gap-x-5">
        <span class="col-span-full text-2xl font-bold pb-3">GOALS</span>
        <span>
            <label>Target Weight (kg): </label>
            <input type="number" class="dialog-input w-[5em]" step=0.01 v-model="weightKgs">
        </span>
        <span>
            in Lbs:
            <span class="opacity-50">{{ weightLbs.toFixed(2) }} lbs</span>
        </span>
        <span>
            <label>Multiplier:</label>
            <input type="number" class="dialog-input w-[4em]" v-model="userGoals['Multiplier']">
        </span>

        <span>
            Total Calories:
            <span class="opacity-50"> {{ totalMacros["Calories"] }} calories</span>
        </span>

        <span>
            <label>% Error</label>
            <input type="number" class="dialog-input w-[4em]" v-model="userGoals['% Error']">
        </span>
        <span>
            Margin of Error:
            <span class="opacity-50"> {{ calErrorMargin }} cals</span>
        </span>
        <span>
            Protein (g/lb body weight):
            <input type="number" class="dialog-input w-[4em]" v-model="userGoals['proteinGPerLBS']" />
        </span>
        <span>
            Fat (g/lb body weight):
            <input type="number" class="dialog-input w-[4em]" v-model="userGoals['fatGPerLBS']" />
        </span>
    </div>

    <div class="flex flex-col md:grid md:grid-cols-2 pt-10 gap-5">
        <span class="col-span-full flex flex-row items-center" v-if="!mealsAllottedAddsUp">
            <span class="icon mdi--error text-red-400/50">The allotted percent towards meals do not add up to 100%</span>
        </span>
        <div class="flex flex-col" v-for="(meal, mealName) in mealGoals" :key="mealName">
            <label class="text-xl">{{ mealName }}</label>
            <span>
                <label>% Allotted </label>
                <input type="number" class="dialog-input w-[4em]" v-model="mealGoals[mealName]"> |
                <label>Margin: {{ (meal * calErrorMargin / 100) }} cals </label>
            </span>
            <span class="grid grid-cols-2 gap-1 m-1">
                <span v-for="macro, macroName in totalMacros">
                    <span>{{ macroName }}: </span>
                    <span class="opacity-50">{{ meal * macro / 100 }} {{ macroName == "Calories" ? 'cals' : 'g'
                        }} </span>
                </span>
            </span>
        </div>
    </div>
    <div class="flex justify-end mt-5"><button class="btn btn-confirm" @click="updateGoals">Save Goals</button>
    </div>

</template>