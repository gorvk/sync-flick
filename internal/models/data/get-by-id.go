package models

import (
	"database/sql"
	"fmt"

	"github.com/gorvk/starterapp/internal/initializers"
	"github.com/gorvk/starterapp/internal/types"
)

func GetById(id string, userClaim types.UserClaim) (*sql.Rows, error) {
	db := initializers.GetDBInstance()
	if db == nil {
		return nil, fmt.Errorf("nil db instance")
	}
	query := fmt.Sprintf("SELECT * from Data WHERE id = '%v' AND user_id = '%v'", id, userClaim.UserId)
	rows, err := db.Query(query)
	return rows, err
}
