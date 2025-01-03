package models

import (
	"github.com/gorvk/starterapp/internal/initializers"
	"github.com/gorvk/starterapp/internal/types"
)

func Update(data types.Data, userClaim types.UserClaim) error {
	db := initializers.GetDBInstance()
	stmt, err := db.Prepare("UPDATE Data SET id = $1 WHERE id = $1 AND user_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data.Id, userClaim.UserId)
	return err
}
