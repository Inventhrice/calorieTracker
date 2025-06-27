import { createApp, toRaw } from './vue.esm-browser.prod.js'
import foodDialog from '../components/foodDB-dialog.js'
import deleteDialog from "../components/confirmDialog.js"
import sidebaritem from '../components/sidebarItem.js'
import contentHeader from '../components/headerComponent.js'
createApp({
    components: { sidebaritem, contentHeader, foodDialog, deleteDialog },
    data() {
        return {
            allFoods: [],
            title: "Food Information",
            selected: null,
            showDialog: false,
            showConfirmDeleteDialog: false
        }
    },
    methods: {
        async fetchData(id = "all") {
            const response = await fetch("/api/foodDB/" + id)
            this.allFoods = await response.json()
        },
        showFoodDialog(foodinfo = null) {
            if (foodinfo != null) {
                this.selected = foodinfo
            } else {
                this.selected = { name: "", calperg: 0.0, proteinperg: 0.0, fatperg: 0.0, carbperg: 0.0, notes: "", source: "" }
            }
            this.showDialog = true
        },
        showDeleteDialog(foodinfo) {
            if (foodinfo != null) {
                this.selected = foodinfo
                this.showConfirmDeleteDialog = true
            }
        },
        async editFood() {
            if (this.selected) {
                if (!this.selected.hasOwnProperty('id')) {
                    let response = (await fetch("/api/foodDB/", { method: "POST", body: JSON.stringify(this.selected) }))
                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.id = data.addedID
                        this.allFoods.push(this.selected)
                    } else {
                        console.log(response.Error)
                    }
                } else {
                    let index = this.allFoods.findIndex((el) => this.selected.id == el.id)
                    let response = await fetch("/api/foodDB/" + this.selected.id, { method: "PATCH", body: JSON.stringify(this.selected) })
                    if (response.ok) {
                        this.allFoods[index] = this.selected
                    } else {
                        let errmsg = await (response.json()).Error
                        console.log(errmsg)
                    }
                }
                this.showDialog = false
            }
        },
        async deleteFood() {
            if (this.selected) {
                let response = await fetch("/api/foodDB/" + this.selected.id, { method: "DELETE" })
                if (response.ok) {
                    this.allFoods.splice(this.allFoods.indexOf(this.selected), 1)
                } else {
                    console.log(response.Error)
                }
                this.showConfirmDeleteDialog = false
            }

        }
    },
    mounted() {
        this.fetchData()
    }
}).mount('#food-db')