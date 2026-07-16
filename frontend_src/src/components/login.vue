<script>

import { api_login } from '../js/api.js'
import ErrPopup from './errPopup.vue';

export default {
    components: { ErrPopup },
    data() {
        return {
            loginCreds: { username: "", password: "" },
            error: { message: "", show: false }
        }
    },
    methods: {
        async login() {
            let response = await api_login(JSON.stringify(this.loginCreds))
            if (!response.ok) {
                let resText = await response.json()
                this.raiseError(resText.Error, "Unable to login.")
            } else {
                this.$emit("login")
            }
        },
        raiseError(errlog, message){
            console.error(errlog)
            this.error.message = message
            this.error.show = true
        }
    }
}
</script>

<template>
    <div id="main-content" class="content-list items-center justify-center">
        <span class="text-3xl lg:text-5xl text my-4">🍪Calorie Tracker🍪</span>
        <div class="module-background lg:w-1/2 flex flex-col gap-2 text p-3" @keyup.enter="login">
            <label class="mx-2" for="username">Username:</label>
            <input class="dialog-input mx-2" type="text" v-model="loginCreds.username">
            <label class="mx-2" for="password">Password:</label>
            <input class="dialog-input mx-2" type="password" v-model="loginCreds.password">
            <span class="justify-end mx-2"><button @click="login" class="btn btn-confirm">Login</button></span>
        </div>
    </div>
    <ErrPopup :message="error.message" :show-err="error.show" @added="error.show = false"></ErrPopup>
</template>
