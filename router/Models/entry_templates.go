package models

import (
	"fmt"

	"example.com/m/v2/middlewares"
)

type EntryTemplate struct {
	ID       int    `json:"id" db:"ID"`
	Meal     string `json:"meal" db:"meal"`
	FoodID   int    `json:"foodID" db:"foodID"`
	Quantity int    `json:"quantity" db:"quantity"`
	UserID   string `db:"userID"`
}

func DeleteTemplate(id int) error {
	fmt.Println(id)
	result, err := middlewares.Database.Exec("DELETE FROM entryTemplates WHERE ID=?", id)

	if _, err := Helper_ExecError(result, err, "Template was unable to be deleted"); err != nil {
		return err
	}
	return nil
}

func AddTemplate(template EntryTemplate) (int64, error) {
	result, err := middlewares.Database.NamedExec("INSERT INTO entryTemplates (userID, meal,foodID,quantity) VALUES (:userID, :meal,:foodID,:quantity)", template)
	if _, err := Helper_ExecError(result, err, "Template was unable to be added"); err != nil {
		return -1, err
	} else {
		newId, _ := result.LastInsertId()
		return newId, nil
	}
}

func GetAllTemplates(userID string) ([]EntryTemplate, error) {
	var templates []EntryTemplate
	if err := middlewares.Database.Select(&templates, "SELECT ID, meal, foodID, quantity FROM entryTemplates WHERE userID=?", userID); err != nil {
		return nil, err
	}
	return templates, nil
}

func GetTemplate(userID string, id int) (EntryTemplate, error) {
	var template EntryTemplate
	if err := middlewares.Database.Get(&template, "SELECT ID, meal, foodID, quantity FROM entryTemplates WHERE userID=? AND ID=?", userID, id); err != nil {
		return EntryTemplate{}, err
	}
	return template, nil
}

func UpdateTemplate(template EntryTemplate) error {
	result, err := middlewares.Database.NamedExec("UPDATE entryTemplates SET meal=:meal, foodID=:foodID, quantity=:quantity WHERE ID=:ID", template)
	if _, err := Helper_ExecError(result, err, "No template with the provided ID found"); err != nil {
		return err
	} else {
		return nil
	}
}
