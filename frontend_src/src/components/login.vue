<script>

import { api_login } from '../js/auth'

export default {
    data() {
        return {
            loginCreds: { username: "", password: "" },
            errorMessage: { message: "", empty: true }
        }
    },
    methods: {
        async login() {
            let response = await api_login(JSON.stringify(this.loginCreds))
            if (!response.ok) {
                let resText = await response.json()
                this.errorMessage.message = "Error: " + resText.Error;
                this.errorMessage.empty = false
            } else {
                this.$emit("login")
            }
        }
    }
}
</script>

<template>
    <div id="main-content" class="content-list items-center justify-center">
        <span class="text-3xl lg:text-5xl text my-4">🍪Calorie Tracker🍪</span>
        <div class="error-popup text p-3 my-3" v-if="!errorMessage.empty" id="error">{{ errorMessage.message }}</div>
        <div class="module-background lg:w-1/2 flex flex-col gap-2 text p-3" @keyup.enter="login">
            <label class="mx-2" for="username">Username:</label>
            <input class="dialog-input mx-2" type="text" v-model="loginCreds.username">
            <label class="mx-2" for="password">Password:</label>
            <input class="dialog-input mx-2" type="password" v-model="loginCreds.password">
            <span class="justify-end mx-2"><button @click="login" class="btn btn-confirm">Login</button></span>
        </div>
    </div>
</template>