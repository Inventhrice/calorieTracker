<script>
import { api_call, api_get } from '../../js/api.js'
export default {
    data() {
        return {
            weightRecorded: 0
        }
    },
    methods: {
        async fetchData() {
            try {
                let response = await api_get("/api/weight/" + this.start)
                let data = await response.json()
                if (response.ok) {
                    this.weightRecorded = data.kg
                } else{
                    this.$emit('raise-error', data.Error, "No weight recorded.")
                    this.weightRecorded = 0    
                }
            } catch (error) {
                this.$emit('raise-error', error, "No weight recorded.")
                this.weightRecorded = 0
            }
        },
        async editWeight() {
            let response = await api_call("/api/weight/", "POST", JSON.stringify({ daterecord: this.start, kg: this.weightRecorded }))
            if (!response.ok) {
                this.$emit('raise-error', response.Error, "Failed to edit weight.")
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
    },
    emits: {
        "raise-error": null
    }
}
</script>

<template>
    <span>
        <label for="weight">Weight: </label>
        <input class="remove-spinner w-[4em]" name="weight" type="number" step="0.01" v-model="weightRecorded" />
        <button @click="editWeight()" class="btn btn-confirm">
            <span class="icon btn-icon mdi--upload"></span>
        </button>
    </span>
</template>
