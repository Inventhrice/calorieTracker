package models

import "example.com/m/v2/middlewares"

type EntryTemplate struct {
	ID       int    `json:"id" db:"id"`
	Meal     string `json:"meal" db:"meal"`
	FoodID   int    `json:"food_id" db:"food_id"`
	Quantity int    `json:"quantity" db:"quantity"`
	UserID   string `db:"user_id"`
}

func DeleteTemplate(id int) error {
	result, err := middlewares.Database.NamedExec("DELETE FROM entryTemplates WHERE id = ?", id)

	if _, err := Helper_ExecError(result, err, "Template was unable to be added"); err != nil {
		return err
	}
	return nil
}

func AddTemplate(template EntryTemplate) (int64, error) {
	result, err := middlewares.Database.NamedExec("INSERT INTO entryTemplates (user_id, meal,food_id,quantity) VALUES (:user_id, :meal,:food_id,:quantity)", template)
	if _, err := Helper_ExecError(result, err, "Template was unable to be added"); err != nil {
		return -1, err
	} else {
		newId, _ := result.LastInsertId()
		return newId, nil
	}
}

func GetAllTemplates(userID string) ([]EntryTemplate, error) {
	var templates []EntryTemplate
	if err := middlewares.Database.Select(&templates, "SELECT id, meal, food_id, quantity FROM entryTemplates WHERE user_id = ?"); err != nil {
		return nil, err
	}
	return templates, nil
}

func GetTemplate(userID string, id int) (EntryTemplate, error) {
	var template EntryTemplate
	if err := middlewares.Database.Get(&template, "SELECT id, meal, food_id, quantity FROM entryTemplates WHERE user_id=? AND id=?", userID, id); err != nil {
		return EntryTemplate{}, err
	}
	return template, nil
}

func UpdateTemplate(template EntryTemplate) error {
	result, err := middlewares.Database.NamedExec("UPDATE entryTemplates SET meal=:meal, food_id=:food_id, quantity=:quantity WHERE id=:id", template)
	if _, err := Helper_ExecError(result, err, "No template with the provided ID found"); err != nil {
		return err
	} else {
		return nil
	}
}
