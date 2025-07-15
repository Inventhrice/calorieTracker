import { createApp } from '../../js/vue.esm-browser.js'
import foodDialog from './foodDB-dialog.js'
import deleteDialog from "../confirmDialog.js"
import sidebar from '../sidebar.js'
import titleHeader from '../titleHeader.js'
import { api_call, api_get } from '../../js/auth.js'
createApp({
    components: { sidebar, titleHeader, foodDialog, deleteDialog },
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
            const response = await api_get("/api/foodDB/" + id)
            if(response.ok){
                this.allFoods = await response.json()
            } else{
                this.allFoods = []
				if(response.status === 401){
					window.location.href = "/app/login.html"
				}
            }
            
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
                    let response = await api_call("/api/foodDB/", "POST", JSON.stringify(this.selected))
                    if (response.ok) {
                        let data = await (response).json()
                        this.selected.id = data.addedID
                        this.allFoods.push(this.selected)
                    } else {
                        console.log(response.Error)
                    }
                } else {
                    let index = this.allFoods.findIndex((el) => this.selected.id == el.id)
                    let response = await api_call("/api/foodDB/" + this.selected.id, "PATCH", JSON.stringify(this.selected))
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
                let response = await api_call("/api/foodDB/" + this.selected.id, "DELETE")
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
