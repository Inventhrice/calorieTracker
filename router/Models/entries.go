package models

import (
	"database/sql"
	"time"

	"example.com/m/v2/middlewares"
)

type Entry struct {
	UserID     string        `db:"userid"`
	ID         int           `json:"id" db:"ID"`
	Meal       string        `json:"meal" db:"meal"`
	DateRecord string        `json:"daterecord" db:"dateRecord"`
	FoodID     sql.NullInt32 `json:"foodID" db:"foodID"`
	Foodname   string        `json:"foodname" db:"foodname"`
	Quantity   float32       `json:"quantity" db:"grams"`
	Cal        float32       `json:"cal" db:"cal"`
	Protein    float32       `json:"protein" db:"protein"`
	Fat        float32       `json:"fat" db:"fat"`
	Carbs      float32       `json:"carbs" db:"carbs"`
	Notes      string        `json:"notes" db:"notes"`
}

func UpdateEntry(entry Entry) error {
	result, err := middlewares.Database.NamedExec("UPDATE entries SET dateRecord=:dateRecord,meal=:meal,foodID=:foodID,grams=:grams,foodname=:foodname,cal=:cal,protein=:protein,fat=:fat,carbs=:carbs WHERE ID=:ID", entry)
	if _, err := Helper_ExecError(result, err, "No entry with the provided ID found"); err != nil {
		return err
	} else {
		return nil
	}
}

func AddEntry(entry Entry) (int64, error) {

	result, err := middlewares.Database.NamedExec("INSERT INTO entries (userid, dateRecord, meal, foodID, foodname, grams, cal, protein, fat, carbs) VALUES (:userid, :dateRecord, :meal, :foodID, :foodname, :grams, :cal, :protein, :fat, :carbs)", entry)

	if _, err := Helper_ExecError(result, err, "Entry was unable to be added"); err != nil {
		return -1, err
	} else {
		newId, _ := result.LastInsertId()
		return newId, nil
	}
}

func DeleteEntry(id int) error {
	result, err := middlewares.Database.Exec("DELETE FROM entries WHERE id = ?", id)
	if _, err := Helper_ExecError(result, err, "No entry with the provided ID found"); err != nil {
		return err
	}
	return nil
}

func GetEntriesByWeek(start string, end string, userID string) ([]Entry, error) {
	entries := []Entry{}
	startDate, _ := time.Parse(time.DateOnly, start)
	endDate, _ := time.Parse(time.DateOnly, end)

	if err := middlewares.Database.Select(&entries, "SELECT ID, dateRecord, meal, foodname, foodID, grams, cal, carbs, protein, fat, notes FROM processed_entries WHERE dateRecord BETWEEN ? AND ? AND userid=?", startDate.Format(time.DateOnly), endDate.Format(time.DateOnly), userID); err != nil {
		return nil, err
	} else {
		return entries, nil
	}
}
