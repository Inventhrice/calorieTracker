export default{
    data() {
        return {
            weightRecorded: 0
        }
    },
    methods: {
        async fetchData(){
            let response = await (await fetch("/api/weight/" + this.start)).json()
            this.weightRecorded = response.kg
        },
        async editWeight() {
            let response = await fetch("/api/weight/", {
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
    template: `
    <span>
        <label for="weight">Weight: </label>
        <input class="remove-spinner" name="weight" type="number" step="0.01" v-model="weightRecorded" />
        <button @click="editWeight()" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-refresh"></span>
        </button>
    </span>
    `
}