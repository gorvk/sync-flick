package models

import (
	"github.com/gorvk/starterapp/internal/initializers"
	"github.com/gorvk/starterapp/internal/types"
)

func Create(data types.Data, userClaim types.UserClaim) error {
	db := initializers.GetDBInstance()

	stmt, err := db.Prepare("INSERT INTO Data VALUES($1, $2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		data.Id,
		userClaim.UserId,
	)

	return err
}
