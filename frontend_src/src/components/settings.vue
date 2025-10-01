<script>
import { api_call, api_get } from "../js/auth.js"
export default {
    data() {
        return {
            title: "Settings",
            loggedInUser: { "First Name": "", "Last Name": "", "Pronouns": "", "Username": "" },
            loginpwd: { "Password": "", "Repeat Password": "" },
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            userGoals: { "Weight": 84.5, "% Error": 15, "Multiplier": 10 },
            macroGoals: { "Total Calories": 1900, "Protein": 0, "Fat": 0, "Carbs": 0 },
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
                this.userGoals["WeightLBS"] = userData.goalLbs
                this.userGoals["Multiplier"] = userData.multiplier
                this.userGoals["% Error"] = userData.acceptablePercent * 100
                for (let index in this.mealTimes) {
                    this.mealGoals[this.mealTimes[index]] = JSON.parse(userData.goalsPerMeal)[index]
                }
                this.macroGoals["Total Calories"] = this.userGoals["Weight"] * this.userGoals["Multiplier"]
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
        }
    },
    created() {
        this.fetchData()
    },
    computed: {
        calErrorMargin:{
            get(){
                return this.userGoals['% Error']*this.totalCals/100
            }
        },
        totalCals: {
            get(){
                return this.userGoals["WeightLBS"]*this.userGoals['Multiplier']
            }
        },
        targetWeight: {
            get(){
                return (this.userGoals["WeightLBS"]/2.2).toFixed(2)
            },
            set(newVal){
                this.userGoals["WeightLBS"] = (newVal * 2.2).toFixed(2)
            }
        },
        mealTimesArray: {
            get(){
                return this.mealTimes
            },
            set(newVal){
                this.mealTimes = newVal.trim().split(',')
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
        <div class="module-background containerDefaults flex flex-col">
            <span class="text-2xl font-bold pb-3">GOALS</span>
            <span>
                <label>Target Weight (kg): </label>
                <input type="number" class="w-[4em]" v-model="targetWeight"> | 
                <span class="opacity-50"><input type="number" class="remove-spinner w-[2.5em]" readonly :value="this.userGoals['WeightLBS']"> lbs</span>
            </span>
            <span>
                <label>Multiplier:</label>
                <input type="number" class="w-[4em]" v-model="this.userGoals['Multiplier']"> |
                <span class="pr-2 opacity-50"><input type="number" class="remove-spinner w-[2.5em]" readonly :value="this.totalCals"> calories</span> 
            </span>
            <span>
                <label>% Error</label>
                <input type="number" class="w-[4em]" v-model="this.userGoals['% Error']">
                <span class="px-2 opacity-50"><input type="number" class="remove-spinner w-[2.5em]" readonly :value="this.calErrorMargin">cals</span>
            </span>
        </div>
        <div class="module-background containerDefaults grid grid-cols-2 gap-3">
            <div class="flex flex-col" v-for="(meal, mealName) in this.mealGoals" :key="mealName">
                <label>{{ mealName }}</label>
                <span>
                    <label>% Allotted</label>
                    <input type="number" class="w-[4em]" v-model="this.mealGoals[mealName]"> |
                    <label>Error: </label>
                </span>
                <span class="grid grid-cols-2 gap-1 m-1">
                    <span>
                        Calories:
                        <span class="opacity-50">{{meal*this.macroGoals['Total Calories']}} cal</span>
                    </span>
                    <span>
                        Fat:
                        <span class="opacity-50">{{meal*this.macroGoals['Fat']}} g</span>
                    </span>
                    
                    <span>
                        Carbs:
                        <span class="opacity-50">{{meal*this.macroGoals['Carbs']}} g</span>
                    </span>
                    <span>
                        Protein:
                        <span class="opacity-50">{{meal*this.macroGoals['Protein']}} g</span>
                    </span>
                </span>
            </div>
        </div>
    </div>
</template>

<style>
@import "tailwindcss";

.informationSpan {
    @apply flex mx-2 place-items-center;
}

.containerDefaults{
    @apply rounded-xl w-fit h-fit p-2;
}
</style>