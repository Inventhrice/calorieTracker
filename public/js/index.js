import { createApp, toRaw } from './vue.esm-browser.js'
import sidebaritem from '../components/sidebarItem.js'
import contentHeader from '../components/headerComponent.js'
import graphHeader from '../components/graphHeader.js'



createApp({
    components: { sidebaritem, contentHeader, graphHeader },
    data() {
        return {
            title: "Dashboard"
        }
    }
}).mount('#dashboard')


createApp({
    components: { sidebaritem, contentHeader },
    data() {
        return {
            title: "Settings"
        }
    }
}).mount('#settings')

createApp({
    data() {
        return {
            loginCreds: {email: "", password: ""},
            errorMessage: {message: "", empty: true}
        }
    },
    methods: {
        async login(){
            let response = await fetch("/api/profile/login", {method: "POST", body: JSON.stringify(this.loginCreds)})
            if(!response.ok){
                let resText = await response.text()
                this.errorMessage.message = "Error: " + resText;
                this.errorMessage.empty = false
            } else{
                window.location.href = "/app/index.html"
            }
        }
    }
}).mount('#login')
