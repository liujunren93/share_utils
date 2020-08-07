package helper

import (
	"github.com/go-playground/validator/v10"
	"github.com/shareChina/utils/netHelper"
)

func SerValidator(i StatusI, data interface{}) error {
	if err := validator.New().Struct(data); err != nil {
		netHelper.RpcResponse(i, StatusBadRequest, err.Error(), nil)
		return err
	}
	return nil
}
