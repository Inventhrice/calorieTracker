import { createApp, toRaw } from './vue.esm-browser.prod.js'
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
