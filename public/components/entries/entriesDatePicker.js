import {getLocalDate, getLastMon, addDate} from "./dateFunctions.js"
export default {
    setup(){
        return {addDate, getLocalDate}
    },
    data() {
        return {
            start:  getLastMon(new Date(getLocalDate(undefined)+"T00:00:00")), //The current displaying week, beginning on Monday and ending on Sunday
        }
    },
    computed: {
        currentWeek: {
            get() {
                // Make a local copy of the start date
                let startLocal = getLastMon(this.start)
                let endLocal = addDate(startLocal, 6)
                return { start: startLocal, end: endLocal }
            },
            set(startDate) {
                this.start = getLastMon(startDate)
            }
        },
        datePickerWrapperCurrentWeek: {
            get() {
                return getLocalDate(this.currentWeek.start)
            },
            set(dateStr) {
                this.currentWeek = new Date(dateStr + "T00:00:00")
            }
        }
    },
	methods: {
		refreshData(daysAdd=undefined){
			if(daysAdd === undefined){
				this.currentWeek = new Date()
			} else{
				this.currentWeek = addDate(this.currentWeek.start, daysAdd)
			}
			this.$emit('fetchEntries', {start: getLocalDate(this.currentWeek.start), end: getLocalDate(this.currentWeek.end)})
		}
	},
    created() {
        this.$emit('fetchEntries', { start: getLocalDate(this.currentWeek.start), end: getLocalDate(this.currentWeek.end) })
    },
    template: `
    <span>
        <button @click="refreshData(-7)" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-navigate-before"></span>
        </button>
        <span class="my-1 p-1">
            From: <input class="dialog-input" type="date" v-model="datePickerWrapperCurrentWeek" />
        </span>
        <button @click="refreshData(7)" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-navigate-next"></span>
        </button>
        <span class="my-1 p-1">To: {{currentWeek.end.toDateString()}}</span>
        <button class="btn" @click="refreshData(undefined)">
            <span class="iconify btn-icon" data-icon="mdi-calendar-today"></span>
        </button>
    </span>
    `
}
