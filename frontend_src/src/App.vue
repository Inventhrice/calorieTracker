<script>
import sidebar from './components/sidebar.vue';
import dashboard from './components/dashboard.vue';
import entries from './components/entries/entries.vue';
import foodDB from './components/foodDB/foodDB.vue';
import settings from './components/settings.vue'
import Sidebar from './components/sidebar.vue';
import login from './components/login.vue';


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
    }
  },
  mounted() {
    window.addEventListener('hashchange', () => {
      this.currentPath = window.location.hash
    })
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
