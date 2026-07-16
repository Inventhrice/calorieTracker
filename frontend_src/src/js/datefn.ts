export function getLocalDate(date: Date | undefined): string{
    if (date === undefined) {
        date = new Date()
        date.setHours(8)
    }
    return date.toISOString().split('T')[0]
}

export function getToday(): Date{
    return new Date(getLocalDate(undefined) + "T00:00:00")
}

export function getLastMon(date: Date): Date {
    // Find the difference of dates till the monday, and get the last Monday that has occured in the week. If the day is the same (1==1), then -1*0 is 9, nothing is added.
    const MONDAY = 1
    date.setHours(8)
    let diffStartToMonday = date.getDay() - MONDAY
	if(date.getDay() == 0){ // this is if the date is sunday
		diffStartToMonday = 7 - MONDAY
	}
    return addDate(date, -1 * diffStartToMonday)
}

export function addDate(date: Date, numDays: number): Date {
    // ret is needed for SetDate
    let ret = new Date(date)
    ret.setDate(ret.getDate() + numDays)
    return ret
}
