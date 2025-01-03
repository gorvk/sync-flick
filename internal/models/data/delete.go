package models

import (
	"fmt"

	"github.com/gorvk/starterapp/internal/initializers"
	"github.com/gorvk/starterapp/internal/types"
)

func Delete(id string, userClaim types.UserClaim) error {
	db := initializers.GetDBInstance()
	if db == nil {
		return fmt.Errorf("nil db instance")
	}

	stmt, err := db.Prepare("DELETE FROM Data WHERE id = $1 AND user_id = $2")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id, userClaim.UserId)

	return err
}
