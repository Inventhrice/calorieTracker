package models

import (
	"database/sql"
	"errors"
	"time"

	"example.com/m/v2/middlewares"
)

type WeightData struct {
	UserID string  `db:"userid"`
	Date   string  `json:"daterecord" db:"dateRecord"`
	KG     float32 `json:"kg" db:"kg"`
}

func GetWeight(start string, userID string) (WeightData, error) {
	weight := WeightData{}
	startDate, _ := time.Parse(time.DateOnly, start)
	if err := middlewares.Database.Get(&weight, "SELECT dateRecord, kg FROM weightTrack WHERE dateRecord=? AND userid=?", startDate.Format(time.DateOnly), userID); err != nil {
		return WeightData{}, err
	} else {
		return weight, nil
	}
}

func AddUpdateWeight(userID string) error {
	var weight WeightData
	weight.UserID = userID

	mode := "PATCH"
	var temp float32
	if err := middlewares.Database.QueryRow("SELECT kg FROM weightTrack WHERE dateRecord=? AND userid=?;", weight.Date, userID).Scan(&temp); err != nil {
		if err == sql.ErrNoRows {
			mode = "POST"
		} else {
			return err
		}
	}

	if mode == "POST" {
		result, err := middlewares.Database.NamedExec("INSERT INTO weightTrack (userid, dateRecord, kg) VALUES (:userid, :dateRecord, :kg)", weight)
		if _, err := Helper_ExecError(result, err, "Information was unable to be added"); err != nil {
			return err
		}
	} else if mode == "PATCH" {
		result, err := middlewares.Database.NamedExec("UPDATE weightTrack SET kg=:kg WHERE dateRecord=:dateRecord AND userid=:userid", weight)
		if _, err := Helper_ExecError(result, err, "Information was unable to be added"); err != nil {
			return err
		}
	} else {
		return errors.New("How the hell did you even get this. mode wasn't POST or PATCH")
	}
	return nil
}
