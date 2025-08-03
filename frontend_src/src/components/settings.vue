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
            mealGoals: { breakfast: 0, lunch: 0, dinner: 0, snacks: 0 },
            macroGoals: { calories: 0, protein: 0, fat: 0, carbs: 0 },
            errorMessage: "",
            changepwd: false
        }
    },
    methods: {
        async fetchData() {
            let response = await api_get("/api/profile")
            if (response.ok) {
                this.loggedInUser = await response.json()
            } else {
                window.location.href = "/app/login.html"
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