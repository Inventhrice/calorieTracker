<script>
import sidebar from './components/sidebar.vue';
import dashboard from './components/dashboard.vue';
import entries from './components/entries/entries.vue';
import foodDB from './components/foodDB/foodDB.vue';
import settings from './components/settings.vue'
import Sidebar from './components/sidebar.vue';
import login from './components/login.vue';
import { api_get } from './js/auth';


const routes = {
  "/": { link: "#/", icon: "view-dashboard", text: "Dashboard", component: dashboard },
  "/entries": { link: "#/entries", icon: "create", text: "Entries", component: entries },
  "/foodDB": { link: "#/foodDB", icon: "food", text: "Food Information", component: foodDB },
  "/settings": { link: "#/settings", icon: "settings", text: "Settings", component: settings }
}

export default {
  components: { sidebar, entries, foodDB, settings, login },
  data() {
    return {
      currentPath: window.location.hash,
      routes: routes,
      loggedin: false
    }
  },
  computed: {
    currentView() {
      if (this.routes) {
        return this.routes[this.currentPath.slice(1) || '/'].component
      } else {
        return entries
      }
    },
    title() {
      if (this.routes) {
        return this.routes[this.currentPath.slice(1) || '/'].text
      }
    }
  },
  methods: {
    logout(){
      document.cookie = "token=;Path=/"
      this.loggedin = false
    },
    async checkLoggedin(){
      if(document.cookie["token"] != ""){
        let response = await api_get("/api/profile")
        if(response.ok){
          this.loggedin = response.status == 200
        } else{
          this.loggedin = false
        }
        
      } else {
        this.loggedin = false
      }
      
    }
  },
  mounted() {
    window.addEventListener('hashchange', () => {
      this.currentPath = window.location.hash
      this.checkLoggedin()
    })
    this.checkLoggedin()
  }
}

</script>

<template>
  <sidebar v-if="loggedin" @logout="logout" :routes></sidebar>
  <div id="content" class="content-container">
    <div v-if="loggedin" id="header" class="top-navigation">
      <h5 class="text title-text">{{ title }}</h5>
    </div>
    <login v-if="!loggedin" @login="loggedin = true"></login>
    <Component v-if="loggedin" :is="currentView"></Component>
  </div>
</template>
