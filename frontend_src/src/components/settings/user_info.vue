<script lang="ts">
import { api_call, api_get } from "../../js/api.js"
import { defineComponent } from 'vue';

export default defineComponent({
    data() {
        return {
            loggedInUser: { "First Name": "", "Last Name": "", "Pronouns": "", "Username": "" },
            loginpwd: { "Password": "", "Repeat Password": "" },
            changepwd: false,
            errMsg: { "Password": "" }
        }
    },
    methods: {
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
    }
})

</script>

<template>
    <div>
        <span class="informationSpan" v-for="_, key in loggedInUser">
            <label :for="key">{{ key }}:</label>
            <input v-model="loggedInUser[key]" class="dialog-input grow my-3" type="text">
        </span>
        <span class="informationSpan flex-row my-2 justify-between gap-x-4">
            <button class="btn btn-uhoh" @click="updateInformation">Save information</button>
            <button class="btn" @click="changepwd = !changepwd">Change Password</button>
        </span>

        <div v-if="changepwd" class="pt-2">
            <span class="informationSpan" v-for="_, key in loginpwd">
                <label :for="key">{{ key }}: </label>
                <input v-model="loginpwd[key]" class="dialog-input grow my-2" type="password">
            </span>
            <span class="m-2 flex justify-between">
                <span>{{ errMsg['Password'] }}</span>
                <button class="btn btn-confirm" @click="changePassword">Save Password</button>
            </span>
        </div>
    </div>
</template>