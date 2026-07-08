import type { NutrientStats } from "./entry.ts"
import { MealTimes } from './entry.ts'

export function calc_day_goals(goalinfo: { multiplier: number, goalLbs: number, proteinGPerLBS: number, fatGPerLBS: number }): NutrientStats {
    const FAT_CALPERGRAM = 9
    const PROTEIN_CALPERGRAM = 4
    const CARB_CALPERGRAM = 4

    let day_nutrients: NutrientStats = {
        cal: goalinfo.goalLbs * goalinfo.multiplier,
        protein: Math.round(goalinfo.proteinGPerLBS * goalinfo.goalLbs),
        fat: Math.round(goalinfo.fatGPerLBS * goalinfo.goalLbs),
        carbs: 0
    }

    day_nutrients.carbs = (day_nutrients.cal - (day_nutrients.fat * FAT_CALPERGRAM) + (day_nutrients.protein * PROTEIN_CALPERGRAM)) / CARB_CALPERGRAM
    return day_nutrients
}

// Applicable for calculating per meal goals AND tolerances
export function calc_pct_goals(totals: NutrientStats, pct: number): NutrientStats {
    let nutrients: NutrientStats = {
        cal: Math.round(totals.cal * pct),
        protein: Math.round(totals.protein * pct),
        fat: Math.round(totals.fat * pct),
        carbs: Math.round(totals.carbs * pct)
    }
    return nutrients
}


export function parse_goals(goalinfo: { "goalLbs": number, "multiplier": number, "acceptablePercent": number, "goalsPerMeal": string, "proteinGPerLBS": number, "fatGPerLBS": number, "UserID": "" }) {
    let all_calculated: any = {}
    all_calculated["goalsPerMeal"] = JSON.parse(goalinfo.goalsPerMeal)
    all_calculated["percentAllowed"] = goalinfo.acceptablePercent

    let day_goals = calc_day_goals(goalinfo)
    let err_daygoals = calc_pct_goals(day_goals, goalinfo.acceptablePercent)
    
    all_calculated["totals"] = { "day": day_goals }
    all_calculated["marginOfError"] = { "day": err_daygoals }

    for (let index in MealTimes) {
        let meal = MealTimes[index]
        let pct_meal = all_calculated["goalsPerMeal"][index]

        all_calculated["totals"][meal] = calc_pct_goals(day_goals, pct_meal)
        all_calculated["marginOfError"][meal] = calc_pct_goals(err_daygoals, pct_meal)
    }
    return all_calculated
}