<script>
import { api_get } from '../../js/auth'
export default {
    data() {
        return {
            weightRecorded: 0
        }
    },
    methods: {
        async fetchData() {
            let response = await api_get("/api/weight/" + this.start)
            if (response.ok) {
                let data = await response.json()
                this.weightRecorded = data.kg
            }

        },
        async editWeight() {
            let response = await api_get("/api/weight/", {
                method: "POST", body:
                    JSON.stringify({ daterecord: this.start, kg: this.weightRecorded })
            })

            if (!response.ok) {
                console.log(response.Error)
            }
        }
    },
    props: {
        start: String
    },
    created() {
        this.$watch('start', () => {
            this.fetchData()
        })
    }
}
</script>

<template>
    <span>
        <label for="weight">Weight: </label>
        <input class="remove-spinner w-[4em]" name="weight" type="number" step="0.01" v-model="weightRecorded" />
        <button @click="editWeight()" class="btn btn-confirm pl-1">
            <span class="iconify btn-icon" data-icon="mdi-refresh"></span>
        </button>
    </span>
</template>