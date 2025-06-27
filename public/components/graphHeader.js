export default {
    props: {
        title: {
            type: String,
            default: "Title"
        }
    },
    data() {
        return {
            timerange: ["1w", "1m", "3m", "All", "Custom"]
        }
    },
    template: `
            <div class="graph-header">
                <span class="text-lg font-semibold">{{title}}</span>
                <div class="graph-timespan-selector">
                    <button class="btn text-base m-1" v-for="el in timerange">{{el}}</button>
                </div>
            </div>`
}