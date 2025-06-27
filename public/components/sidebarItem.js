//https://docs.emmet.io/cheat-sheet/
//https://icon-sets.iconify.design/mdi/page-2.html
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
    template: `<div class="sidebar">
    <a v-for="icon in sidebarIcons" :href=icon.link>
        <div class="sidebar-icon group">
            <span class="iconify" :data-icon="'mdi-' + icon.icon"></span>
            <span class="sidebar-tooltip group-hover:scale-100">{{icon.text}}</span>
        </div>
    </a></div>`
}