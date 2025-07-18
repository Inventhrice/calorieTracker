<script>
    import { api_logout } from "../js/auth.js"
    export default {
        props: {
            routes: Object
        },
        methods: {
            async logout(){
                await api_logout()
                this.$emit('logout')
            }
        }
    }
</script>

<template>
    <div class="sidebarcontainer">
    <span v-for="icon in routes" >
		<a :href=icon.link class="sidebarIconContainer group">
			<span class="iconify" :data-icon="'mdi-' + icon.icon"></span>
			<span class="sidebarFloatingText">{{icon.text}}</span>
		</a>
    </span>
	<span @click="logout" class="sidebarIconContainer group">
		<span class="iconify" :data-icon="'mdi-logout'"></span>
		<span class="sidebarFloatingText">Logout</span>
	</span>
</div>
</template>

<style scoped>
@reference "../index.css";

.sidebarcontainer{
    @apply flex flex-col w-auto h-screen p-2 bg-blue-900/50 dark:bg-gray-900 shadow-lg;
}

.sidebarIconContainer {
    @apply relative flex items-center justify-center w-auto m-1 p-3 bg-white hover:bg-blue-600 dark:bg-gray-800 text-blue-500 hover:text-white hover:rounded-xl rounded-3xl transition-all duration-300 ease-linear cursor-pointer shadow-lg;
}

.sidebarFloatingText {
    @apply absolute p-2 m-3 min-w-max left-14 rounded-md shadow-md z-50 text-white bg-gray-900 text-xs font-bold transition-all duration-100 scale-0 origin-left group-hover:scale-100;
}

</style>

