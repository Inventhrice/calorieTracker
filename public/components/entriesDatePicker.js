export default {
    data() {
        return {
            start: this.getLastMon(new Date()), //The current displaying week, beginning on Monday and ending on Sunday
        }
    },
    computed: {
        currentWeek: {
            get() {
                // Make a local copy of the start date
                let startLocal = this.getLastMon(this.start)
                let endLocal = this.addDate(startLocal, 6)
                return { start: startLocal, end: endLocal }
            },
            set(startDate) {
                this.start = this.getLastMon(startDate)
            }
        },
        datePickerWrapperCurrentWeek: {
            get() {
                return this.getLocalDate(this.currentWeek.start)
            },
            set(dateStr) {
                this.currentWeek = new Date(dateStr + "T00:00:00")
            }
        }
    },
    methods: {
        getLocalDate(date = undefined) {
            if (date === undefined) {
                date = new Date()
                date.setUTCHours(8)
            }
            return date.toISOString().split('T')[0]
        },
        getLastMon(date) {
            // Find the difference of dates till the monday, and get the last Monday that has occured in the week. If the day is the same (1==1), then -1*0 is 9, nothing is added.
            const MONDAY = 1
            let diffStartToMonday = date.getDay() - MONDAY
            return this.addDate(date, -1 * diffStartToMonday)
        },
        addDate(date, numDays) {
            // ret is needed for SetDate
            let ret = new Date(date)
            ret.setDate(ret.getDate() + numDays)
            return ret
        }
    },
    created() {
        this.$emit('fetchEntries', { start: this.getLocalDate(this.currentWeek.start), end: this.getLocalDate(this.currentWeek.end) })
    },
    template: `
    <span>
        <button @click="currentWeek = addDate(currentWeek.start, -7)" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-navigate-before"></span>
        </button>
        <span class="my-1 p-1">
            From: <input class="dialog-input" type="date" v-model="datePickerWrapperCurrentWeek" />
        </span>
        <button @click="currentWeek = addDate(currentWeek.start, 7)" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-navigate-next"></span>
        </button>
        <span class="my-1 p-1">To: {{currentWeek.end.toDateString()}}</span>
        <button class="btn" @click="currentWeek = new Date()">
            <span class="iconify btn-icon" data-icon="mdi-calendar-today"></span>
        </button>
        <button @click="$emit('fetchEntries', {start: getLocalDate(currentWeek.start), end: getLocalDate(currentWeek.end)})" class="btn btn-confirm">
            <span class="iconify btn-icon" data-icon="mdi-refresh"></span>
        </button>
    </span>
    `
}