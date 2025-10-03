package models

import (
	"example.com/m/v2/middlewares"
)

type FoodInfo struct {
	UserID      string  `db:"userid"`
	ID          int     `json:"id" db:"ID"`
	Name        string  `json:"name" db:"name"`
	CalPerG     float32 `json:"calperg" db:"calPerG"`
	ProteinPerG float32 `json:"proteinperg" db:"proteinPerG"`
	FatPerG     float32 `json:"fatperg" db:"fatPerG"`
	CarbPerG    float32 `json:"carbperg" db:"carbPerG"`
	Notes       string  `json:"notes"  db:"notes"`
	Source      string  `json:"source"  db:"source"`
}

func UpdateFood(food FoodInfo) error {
	result, err := middlewares.Database.NamedExec("UPDATE food_info SET name=:name, calPerG=:calPerG, proteinPerG=:proteinPerG, fatPerG=:fatPerG, carbPerG=:carbPerG, notes=:notes WHERE id = :ID", food)
	if _, err := Helper_ExecError(result, err, "No food info with the provided ID found"); err != nil {
		return err
	} else {
		return nil
	}
}

func AddFood(food FoodInfo) (int64, error) {
	result, err := middlewares.Database.NamedExec("INSERT INTO food_info (userid, name, calPerG, proteinPerG, fatPerG, carbPerG, notes, source) VALUES (:userid, :name, :calPerG, :proteinPerG, :fatPerG, :carbPerG, :notes, :source)", food)
	if _, err := Helper_ExecError(result, err, "Food info was unable to be added"); err != nil {
		return -1, err
	} else {
		newID, _ := result.LastInsertId()
		return newID, nil

	}
}

func GetFood(id int, userID string) (FoodInfo, error) {
	var food FoodInfo

	if err := middlewares.Database.Get(&food, "SELECT * FROM food_info WHERE id=? AND userid=?", id, userID); err != nil {
		return FoodInfo{}, err
	} else {
		return food, nil
	}
}

func DeleteFood(id int) error {
	result, err := middlewares.Database.Exec("DELETE FROM food_info WHERE id = ?", id)
	if _, err := Helper_ExecError(result, err, "No food info with the provided ID found"); err != nil {
		return err
	} else {
		return nil
	}
}

func GetAllFoods(userid string) ([]FoodInfo, error) {
	var listFoods []FoodInfo
	if err := middlewares.Database.Select(&listFoods, "SELECT ID, name, calPerG, proteinPerG, fatPerG, carbPerG, notes, source FROM food_info WHERE userid=?", userid); err != nil {
		return nil, err
	} else {
		return listFoods, nil
	}

}
