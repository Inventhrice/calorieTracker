package models

import (
	"database/sql"
	"errors"
)

func Helper_ExecError(r sql.Result, initerr error, noRowsFound_errMsg string) (int64, error) {
	// CHecks if the initial call to Query has an error
	if initerr != nil {
		return 0, initerr
	}

	// Checks how many rows are affected, and returns nil if there's an error
	rowsAffected, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 0 {
		return 0, errors.New(noRowsFound_errMsg)
	}

	return rowsAffected, nil
}
