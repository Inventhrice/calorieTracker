import { createApp } from './vue.esm-browser.js'
import sidebar from '../components/sidebar.js'
import titleHeader from '../components/titleHeader.js'
import graphHeader from '../components/graphHeader.js'
import {api_get, api_login, api_call} from "./auth.js"


createApp({
    components: { sidebar, titleHeader, graphHeader },
    data() {
        return {
            title: "Dashboard"
        }
    }
}).mount('#dashboard')


createApp({
    components: { sidebar, titleHeader },
    data() {
        return {
            title: "Settings",
            loggedInUser: {firstname: "", lastname: "", pronouns: "", username: "", password: "", repeatpwd: ""},
			errorMessage: ""
        }
    },
    methods: {
        async fetchData(){
            let response = await api_get("/api/profile")
            if(response.ok){
                this.loggedInUser = await response.json()
            } else{
				window.location.href = "/app/login.html"
            }
        },
		async changePassword(){
			let user = this.loggedInUser
			if(user.password === user.repeatpwd){
				let response = await api_call("/api/profile/password", "PATCH", JSON.stringify({"password": user.password}))
				if(response.ok){
					user.password = ""
					user.repeatpwd = ""
				} else{
					let msg = await response.text()
					console.log(msg)
				}
			} else{
				this.errorMessage = "Passwords do not match."
			}
		}
    },
    created() {
        this.fetchData()
    }
}).mount('#settings')

createApp({
    data() {
        return {
            loginCreds: {username: "", password: ""},
            errorMessage: {message: "", empty: true}
        }
    },
    methods: {
        async login(){
            let response = await api_login(JSON.stringify(this.loginCreds))
            if(!response.ok){
                let resText = await response.json()
                this.errorMessage.message = "Error: " + resText.Error;
                this.errorMessage.empty = false
            } else{
                window.location.href = "/app/index.html"
            }
        }
    }
}).mount('#login')
