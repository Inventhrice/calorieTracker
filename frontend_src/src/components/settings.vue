<script>
import { api_call, api_get } from "../js/auth.js"
export default {
    data() {
        return {
            title: "Settings",
            loggedInUser: { "First Name": "", "Last Name": "", "Pronouns": "", "Username": ""},
            loginpwd:  {"Password": "", "Repeat Password": "" },
            mealTimes: ["Breakfast", "Lunch", "Dinner", "Snacks"],
            userGoals: { "Weight": 0, "% Error": 0, "Multiplier": 0 },
            macroGoals: { calories: 0, protein: 0, fat: 0, carbs: 0 },
            mealGoals: {},
            errorMessage: "",
            changepwd: false
        }
    },
    methods: {
        async fetchPreferences() {
            let response = await api_get("/api/goals")
            if (response.ok) {
                let userData = await response.json()
                this.userGoals["Weight"] = userData.goalLbs
                this.userGoals["Multiplier"] = userData.multiplier
                this.userGoals["% Error"] = userData.acceptablePercent
                for (let index in mealTimes){
                    this.mealGoals[this.mealTimes[index]] = JSON.parse(userData.goalsPerMeal)[index]
                }
                
            }
        },
        async fetchUserInfo() {
            let response = await api_get("/api/profile")
            if (response.ok){
                let userinfo = await response.json()
                this.loggedInUser["First Name"] = userinfo.firstname
                this.loggedInUser["Last Name"] = userinfo.lastname
                this.loggedInUser["Pronouns"] = userinfo.pronouns
                this.loggedInUser["Username"] = userinfo.username
            } else{
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
                this.errorMessage = "Passwords do not match."
            }
        }
    },
    created() {
        this.fetchData()
    }
}
</script>

<template>
    <div class="flex flex-row md:grid md:grid-cols-2 mt-2 ml-2 gap-y-3">
        <div class="module-background rounded-xl w-fit flex flex-col">
            <span class="informationSpan" v-for="info, key in loggedInUser">
                <label :for="key">{{key}}:</label>
                <input v-model="loggedInUser[key]" class="dialog-input grow my-3" type="text">
            </span>
            <span class="informationSpan flex-row my-2 justify-between gap-x-4">
                <button class="btn btn-uhoh" @click="updateInformation">Save information</button>
                <button class="btn" @click="changepwd = !changepwd">Change Password</button>
            </span>

            <div v-if="changepwd" class="">
                <span class="informationSpan" v-for="data, key in loginpwd">
                    <label :for="key">{{key}}: </label>
                    <input v-model="loginpwd[key]" class="dialog-input grow my-2" type="password">
                </span>
                <span class="m-2 flex justify-between">
                    <span>{{ errorMessage }}</span>
                    <button class="btn btn-confirm" @click="changePassword">Save Password</button>
                </span>
            </div>
        </div>
        <div class="module-background rounded-xl justify-self-center grid grid-cols-2">
            <span class="text-xl py-2 place-content-center text-center col-span-2">Goals Information</span>

            <div class="flex flex-col">
                <span v-for="goals, key in userGoals" :key class="informationSpan">
                    <label for="key"> {{ key }}</label>
                    <input type="text" class="dialog-input grow my-2" v-model="goals[key]">
                </span>
            </div>

            <span class="col-span-1"></span>

            <!-- TEMP -->
            <div class="flex flex-col">
                <span v-for="goals, key in userGoals" :key class="informationSpan">
                    <label for="key"> {{ key }}</label>
                    <input type="text" class="dialog-input grow my-2" v-model="goals[key]">
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
</style>