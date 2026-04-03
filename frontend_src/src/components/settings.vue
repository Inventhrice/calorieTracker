<script>
import { api_call, api_get } from "../js/api.js"
import errPopup from "./errPopup.vue"
export default {
    components: { errPopup },
    data() {
        return {
            title: "Settings",
            loggedInUser: { "First Name": "", "Last Name": "", "Pronouns": "", "Username": "" },
            loginpwd: { "Password": "", "Repeat Password": "" },
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            userGoals: { "Weight": 84.5, "% Error": 15, "Multiplier": 10, "proteinGPerLBS": 0.8, "fatGPerLBS": 0.4 },
            mealGoals: { "Breakfast": 0, "Lunch": 0, "Dinner": 0, "Snacks": 0 },
            errMsg: {},
            changepwd: false
        }
    },
    methods: {
        async fetchPreferences() {
            let response = await api_get("/api/goals")
            if (response.ok) {
                let userData = await response.json()
                this.userGoals["Weight"] = (userData.goalLbs / 2.2).toFixed(2)
                this.userGoals["Multiplier"] = userData.multiplier
                this.userGoals["% Error"] = userData.acceptablePercent * 100
                this.userGoals["proteinGPerLBS"] = userData["proteinGPerLBS"]
                this.userGoals["fatGPerLBS"] = userData["fatGPerLBS"]
                for (let index in this.mealTimes) {
                    this.mealGoals[this.mealTimes[index]] = JSON.parse(userData.goalsPerMeal)[index] * 100
                }
            }
        },
        async fetchUserInfo() {
            let response = await api_get("/api/profile")
            if (response.ok) {
                let userinfo = await response.json()
                this.loggedInUser["First Name"] = userinfo.firstname
                this.loggedInUser["Last Name"] = userinfo.lastname
                this.loggedInUser["Pronouns"] = userinfo.pronouns
                this.loggedInUser["Username"] = userinfo.username
            } else {
                console.log("User info has an error in being fetched.")
            }
        },
        async updateInformation() {
            let modifiedUser = {
                "firstname": this.loggedInUser["First Name"],
                "lastname": this.loggedInUser["Last Name"],
                "pronouns": this.loggedInUser["Pronouns"],
                "username": this.loggedInUser["Username"]
            }
            let response = await api_call("/api/profile", "PATCH", JSON.stringify(modifiedUser))
            if (!response.ok) {
                console.error("Information was unable to be updated.")
            }
        },
        async fetchData() {
            let response = await api_get("/api/profile")
            if (response.ok) {
                await this.fetchPreferences()
                await this.fetchUserInfo()
            } else {
                console.error("Not signed in.")
            }
        },
        async changePassword() {
            let user = this.loginpwd
            if (user["Password"] === user["Repeat Password"]) {
                let response = await api_call("/api/profile/password", "PATCH", JSON.stringify({ "password": user["Password"] }))
                if (response.ok) {
                    user["Password"] = ""
                    user["Repeat Password"] = ""
                } else {
                    let msg = await response.text()
                    console.log(msg)
                }
            } else {
                this.errMsg['Password'] = "Passwords do not match."
            }
        },
        async updateGoals() {
            if (this.mealsAllottedAddsUp) {
                let temp = []
                for (let meal in this.mealGoals) {
                    temp.push(this.mealGoals[meal] / 100)
                }

                let obj = {
                    "goalLbs": parseFloat(this.weightLbs),
                    "multiplier": this.userGoals["Multiplier"],
                    "acceptablePercent": this.userGoals["% Error"] / 100,
                    "goalsPerMeal": JSON.stringify(temp),
                    "proteinGPerLBS": parseFloat(this.userGoals["proteinGPerLBS"]),
                    "fatGPerLBS": parseFloat(this.userGoals["fatGPerLBS"])
                }

                let response = await api_call("/api/goals", "POST", JSON.stringify(obj))
                if (!response.ok) {
                    console.error("POST request for updating goals failed.")
                }
            } else{
                console.error("Cannot save goals, meals do not add up.")
            }

        }
    },
    created() {
        this.fetchData()
    },
    computed: {
        weightLbs: {
            get() {
                return Math.round(this.userGoals['Weight'] * 2.2).toFixed(2)
            }
        },
        calErrorMargin: {
            get() {
                return (this.userGoals['% Error'] * this.totalMacros["Calories"] / 100)
            }
        },
        totalMacros: {
            get() {
                const FAT_CALPERGRAM = 9
                const PROTEIN_CALPERGRAM = 4
                const CARB_CALPERGRAM = 4
                let tMacros = { "Calories": 0, "Protein": 0, "Fat": 0, "Carbs": 0 }
                tMacros["Calories"] = this.weightLbs * this.userGoals['Multiplier']
                tMacros["Protein"] = Math.round(this.weightLbs * this.userGoals["proteinGPerLBS"])
                tMacros["Fat"] = Math.round(this.weightLbs * this.userGoals["fatGPerLBS"])
                tMacros["Carbs"] = (tMacros["Calories"] - ((tMacros["Fat"] * FAT_CALPERGRAM) + (tMacros["Protein"] * PROTEIN_CALPERGRAM))) / CARB_CALPERGRAM
                return tMacros
            }
        },
        mealsAllottedAddsUp: {
            get() {
                let add = 0
                for (let meal in this.mealGoals) {
                    add += this.mealGoals[meal]
                }
                return add == 100
            }
        }
    }
}
</script>

<template>
    <div class="content-list items-start space-y-3">
        <div class="module-background containerDefaultsq flex flex-col">
            <span class="informationSpan" v-for="info, key in loggedInUser">
                <label :for="key">{{ key }}:</label>
                <input v-model="loggedInUser[key]" class="dialog-input grow my-3" type="text">
            </span>
            <span class="informationSpan flex-row my-2 justify-between gap-x-4">
                <button class="btn btn-uhoh" @click="updateInformation">Save information</button>
                <button class="btn" @click="changepwd = !changepwd">Change Password</button>
            </span>
        </div>

        <div v-if="changepwd" class="module-background containerDefaults flex flex-col">
            <span class="informationSpan" v-for="data, key in loginpwd">
                <label :for="key">{{ key }}: </label>
                <input v-model="loginpwd[key]" class="dialog-input grow my-2" type="password">
            </span>
            <span class="m-2 flex justify-between">
                <span>{{ errMsg['Password'] }}</span>
                <button class="btn btn-confirm" @click="changePassword">Save Password</button>
            </span>
        </div>

        <div class="module-background containerDefaults">
            <div class="grid grid-cols-2 gap-x-5">
                <span class="col-span-full text-2xl font-bold pb-3">GOALS</span>
                <span>
                    <label>Target Weight (kg): </label>
                    <input type="number" class="dialog-input w-[5em]" v-model="this.userGoals['Weight']">
                </span>
                <span>
                    in Lbs:
                    <span class="opacity-50">{{ this.weightLbs }} lbs</span>
                </span>
                <span>
                    <label>Multiplier:</label>
                    <input type="number" class="dialog-input w-[4em]" v-model="this.userGoals['Multiplier']">
                </span>

                <span>
                    Total Calories:
                    <span class="opacity-50"> {{ this.totalMacros["Calories"] }} calories</span>
                </span>

                <span>
                    <label>% Error</label>
                    <input type="number" class="dialog-input w-[4em]" v-model="this.userGoals['% Error']">
                </span>
                <span>
                    Margin of Error:
                    <span class="opacity-50"> {{ this.calErrorMargin }} cals</span>
                </span>
                <span>
                    Protein (g/lb body weight):
                    <input type="number" class="dialog-input w-[4em]" v-model="this.userGoals['proteinGPerLBS']" />
                </span>
                <span>
                    Fat (g/lb body weight):
                    <input type="number" class="dialog-input w-[4em]" v-model="this.userGoals['fatGPerLBS']" />
                </span>
            </div>

            <div class="flex flex-col md:grid md:grid-cols-2 pt-10 gap-5">
                <span class="col-span-full flex flex-row items-center" v-if="!this.mealsAllottedAddsUp">
                    <err-popup></err-popup> The allotted percent towards meals do not add up to 100%
                </span>
                <div class="flex flex-col" v-for="(meal, mealName) in this.mealGoals" :key="mealName">
                    <label class="text-xl">{{ mealName }}</label>
                    <span>
                        <label>% Allotted </label>
                        <input type="number" class="dialog-input w-[4em]" v-model="this.mealGoals[mealName]"> |
                        <label>Margin: {{ (meal * this.calErrorMargin / 100) }} cals </label>
                    </span>
                    <span class="grid grid-cols-2 gap-1 m-1">
                        <span v-for="macro, macroName in this.totalMacros">
                            <span>{{ macroName }}: </span>
                            <span class="opacity-50">{{ meal * macro / 100 }} {{ macroName == "Calories" ? 'cals' : 'g'
                            }} </span>
                        </span>
                    </span>
                </div>
            </div>
            <div class="flex justify-end mt-5"><button class="btn btn-confirm" @click="updateGoals">Save Goals</button>
            </div>
        </div>
    </div>
</template>

<style>
@import "tailwindcss";

.informationSpan {
    @apply flex mx-2 place-items-center;
}

.containerDefaults {
    @apply rounded-xl w-fit h-fit p-2;
}
</style>
