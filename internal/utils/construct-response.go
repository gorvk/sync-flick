package utils

import (
	"encoding/json"

	"github.com/gorvk/starterapp/internal/types"
)

func ConstructResponse(isSuccess bool, result any) ([]byte, error) {
	response := types.DATA_RESPONSE_PARAMETERS{
		RESPONSE_COMMON_PARAMETERS: types.RESPONSE_COMMON_PARAMETERS{
			IsSuccess: isSuccess,
		},
	}
	data, err := json.Marshal(response)
	return data, err
}
