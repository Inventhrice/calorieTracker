export function getLocalDate(date = undefined) {
    if (date === undefined) {
        date = new Date()
        date.setHours(8)
    }
    return date.toISOString().split('T')[0]
}

export function getLastMon(date) {
    // Find the difference of dates till the monday, and get the last Monday that has occured in the week. If the day is the same (1==1), then -1*0 is 9, nothing is added.
    const MONDAY = 1
    date.setHours(8)
    let diffStartToMonday = date.getDay() - MONDAY
    return addDate(date, -1 * diffStartToMonday)
}

export function addDate(date, numDays) {
    // ret is needed for SetDate
    let ret = new Date(date)
    ret.setDate(ret.getDate() + numDays)
    return ret
}
