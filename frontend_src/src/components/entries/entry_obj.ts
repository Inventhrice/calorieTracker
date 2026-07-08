import { getToday } from "../../js/datefn.js"

export const MealTimes = ["Breakfast", "Lunch", "Dinner", "Snacks"]

export class Entry {

    daterecord: Date
    foodname: string
    foodID?: number
    quantity: number
    cal: number
    protein: number
    fat: number
    carbs: number
    meal: string
    notes: string

    constructor() {
        this.daterecord = getToday()
        this.foodname = ""
        this.foodID = undefined
        this.quantity = 0
        this.cal = 0
        this.protein = 0
        this.fat = 0
        this.carbs = 0
        this.meal = ""
        this.notes = ""
    }

    static from(json: any): Entry{
        let foodID = (json.foodID.Valid) ? json.foodID.Int32 : undefined;
        let e = new Entry();
        e.daterecord = json.daterecord
        e.foodname = json.foodname
        e.foodID = foodID
        e.quantity = json.quantity
        e.cal = json.cal
        e.protein = json.protein
        e.fat = json.fat
        e.carbs = json.carbs
        e.notes = json.notes
        e.meal = json.meal
        return e
    }
     

}