/*
<span class="self-end" @click="logout">
		<span class="iconify" :data-icon="'mdi-logout'"></span>
		<span class="absolute p-2 m-3 min-w-max left-14 rounded-md shadow-md z-50 text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100">Logout</span>
	</span>
*/

import { api_logout } from "../js/auth.js"
export default {
    data() {
        return {
            sidebarIcons: [
                { icon: "view-dashboard", text: "Dashboard", link: "/app/index.html" },
                { icon: "create", text: "Entries", link: "/app/entries.html" },
                { icon: "food", text: "Food Information", link: "/app/foodDB.html" },
                { icon: "settings", text: "Settings", link: "/app/settings.html" }
            ]
        }
    },
    methods: {
        async logout(){
            await api_logout()
        }
    },
    template: `
<div class="flex flex-col w-auto h-screen p-2 bg-white dark:bg-gray-900 shadow-lg">
    <span v-for="icon in sidebarIcons" >
		<a :href=icon.link class="relative flex items-center justify-center w-auto m-1 p-3 bg-gray-400 hover:bg-blue-600 dark:bg-gray-800 text-blue-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg group">
			<span class="iconify" :data-icon="'mdi-' + icon.icon"></span>
			<span class="absolute p-2 m-3 min-w-max left-14 rounded-md shadow-md z-50 text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100">{{icon.text}}</span>
		</a>
    </span>
</div>
`
}


