<script>
import { api_call, api_get } from "../js/auth.js"
export default {
    data() {
        return {
            title: "Settings",
            loggedInUser: { firstname: "", lastname: "", pronouns: "", username: "", password: "", repeatpwd: "" },
            errorMessage: ""
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
            let user = this.loggedInUser
            if (user.password === user.repeatpwd) {
                let response = await api_call("/api/profile/password", "PATCH", JSON.stringify({ "password": user.password }))
                if (response.ok) {
                    user.password = ""
                    user.repeatpwd = ""
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
    <div class="content-list">
        <div class="module-background rounded-2xl flex flex-col">
            <span class="mx-2"><label for="firstname">Firstname: </label><input v-model="loggedInUser.firstname"
                    class="dialog-input m-3" type="text"></span>
            <span class="mx-2"><label for="lastname">Lastname: </label><input v-model="loggedInUser.lastname"
                    class="dialog-input m-3" type="text"></span>
            <span class="mx-2"><label for="pronouns">Pronouns: </label><input v-model="loggedInUser.pronouns"
                    class="dialog-input m-3" type="text"></span>
        </div>
        <div class="module-background rounded-2xl flex flex-col">
            <span class="mx-2">
                <label for="username">Username: </label><input v-model="loggedInUser.username" class="dialog-input my-2"
                    type="text">
            </span>
            <span class="mx-2">
                <label for="password">New Password: </label>
                <input v-model="loggedInUser.password" class="dialog-input my-2" type="password">
            </span>
            <span class="mx-2">
                <label for="password">Reenter Password: </label>
                <input v-model="loggedInUser.repeatpwd" class="dialog-input my-2" type="password">
            </span>
            <span class="m-2 flex justify-between">
                <span>{{ errorMessage }}</span>
                <button class="btn btn-confirm" @click="changePassword">Save</button>
            </span>
        </div>
        <div class="module-background rounded-2xl flex flex-col">
            <span class="m-2"><label for="goalsData">JSON Data for Goals:</label><textarea
                    class="w-full dialog-input my-3"></textarea></span>
        </div>
    </div>
</template>