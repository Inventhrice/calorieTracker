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
            let response = await api_logout()
			window.location.href = "/app/login.html"
        }
    },
    template: `
<div class="flex flex-col w-auto h-screen p-2 bg-blue-900/50 dark:bg-gray-900 shadow-lg">
    <span v-for="icon in sidebarIcons" >
		<a :href=icon.link class="relative flex items-center justify-center w-auto m-1 p-3 bg-white hover:bg-blue-600 dark:bg-gray-800 text-blue-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg group">
			<span class="iconify" :data-icon="'mdi-' + icon.icon"></span>
			<span class="absolute p-2 m-3 min-w-max left-14 rounded-md shadow-md z-50 text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100">{{icon.text}}</span>
		</a>
    </span>
	<span @click="logout" class="relative flex items-center justify-center w-auto m-1 p-3 bg-white hover:bg-blue-600 dark:bg-gray-800 text-blue-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg group">
		<span class="iconify" :data-icon="'mdi-logout'"></span>
		<span class="absolute p-2 m-3 min-w-max left-14 rounded-md shadow-md z-50 text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100">Logout</span>
	</span>
</div>
`
}


