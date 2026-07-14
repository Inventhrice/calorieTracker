import { getToday } from "../../js/datefn.ts"

export const MealTimes = ["Breakfast", "Lunch", "Dinner", "Snacks"]

export class NutrientStats {
    cal: number
    protein: number
    fat: number
    carbs: number 

    constructor(cal = 0, protein = 0, fat = 0, carbs = 0){
        this.cal = cal
        this.protein = protein
        this.fat = fat
        this.carbs = carbs
    }

    static clone(stats: NutrientStats){
        let e = new NutrientStats()
        e.add(stats)
        return e
    }

    add(entry: NutrientStats | Entry) {
        this.cal += entry.cal
        this.protein += entry.protein
        this.fat += entry.fat
        this.carbs += entry.carbs
    }

    subtract(entry: NutrientStats | Entry) {
        this.cal -= entry.cal
        this.protein -= entry.protein
        this.fat -= entry.fat
        this.carbs -= entry.carbs
    }

    getField(name: "cal" | "protein" | "fat" | "carbs" ): number{
        switch(name){
            case "cal":
                return this.cal
                break
            case "protein":
                return this.protein
                break
            case "fat":
                return this.fat
                break
            case "carbs":
                return this.carbs
                break
            default:
                return -1
        }
    }

}

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
    id?: number

    constructor() {
        this.daterecord = getToday()
        this.foodname = ""
        this.foodID = undefined
        this.id = undefined
        this.quantity = 0
        this.cal = 0
        this.protein = 0
        this.fat = 0
        this.carbs = 0
        this.meal = ""
        this.notes = ""
    }

    static from(json: any): Entry {
        let foodID = (json.foodID.Valid) ? json.foodID.Int32 : undefined;
        let e = new Entry();
        if("id" in json) e.id = json.id
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