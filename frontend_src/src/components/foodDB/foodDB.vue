<script>
import foodDialog from './foodDB-dialog.vue'
import deleteDialog from "../confirmDialog.vue"
import { api_call, api_get } from '../../js/auth.js'
export default {
    components: { foodDialog, deleteDialog },
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
            if (response.ok) {
                this.allFoods = await response.json()
            } else {
                this.allFoods = []
                if (response.status === 401) {
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
}
</script>

<template>
    <div class="content-list">
        <div class="flex justify-between content-center w-98/100 text">
            <span class="content-center font-semibold">{{ allFoods.length }} Items found</span>
            <button class="flex btn btn-confirm" @click="showFoodDialog()">
                <span class="iconify btn-icon text-2xl pr-1" data-icon="mdi-hamburger-plus"></span>
                <span>Add new food</span>
            </button>
        </div>
        <table class="w-98/100 table-border table-auto text">
            <thead class="module-background">
                <tr>
                    <th class="text-left pl-2">Name</th>
                    <th>Calorie (g)</th>
                    <th>Protein (g)</th>
                    <th>Fat (g)</th>
                    <th>Carb (g)</th>
                    <th class="text-center">Actions</th>
                </tr>
            </thead>
            <tbody>
                <tr class="table-border" v-for="foodinfo in allFoods" :key="foodinfo.id">
                    <td class="text-left pl-2">
                        <a v-if=foodinfo.source.Valid :href=foodinfo.source.String
                            class="font-semibold underline">{{ foodinfo.name }}</a>
                        <span v-else>{{ foodinfo.name }}</span>
                    </td>
                    <td>{{ foodinfo.calperg }}</td>
                    <td>{{ foodinfo.proteinperg }}</td>
                    <td>{{ foodinfo.fatperg }}</td>
                    <td>{{ foodinfo.carbperg }}</td>
                    <td class="text-center py-1">
                        <span class="p-1">
                            <button class="btn" @click="showFoodDialog(foodinfo)"> <span class="iconify btn-icon"
                                    data-icon="mdi-pencil"></span> </button>
                        </span>
                        <span class="p-1">
                            <button class="btn btn-uhoh px-1" @click="showDeleteDialog(foodinfo)"> <span
                                    class="iconify btn-icon" data-icon="mdi-trash-can"></span> </button>
                        </span>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
    <food-dialog v-if=showDialog @close-dialog="this.showDialog = false" :selected
        @confirm-dialog="editFood"></food-dialog>
    <delete-dialog v-if=showConfirmDeleteDialog @confirm-dialog="deleteFood"
        @close-dialog="this.showConfirmDeleteDialog = false"></delete-dialog>
</template>
