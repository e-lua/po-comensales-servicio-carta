package repositories

import (
	"strconv"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Re_Set_Notified(idbusiness int) error {

	_, err_do := models.RedisCN.Get().Do("SET", strconv.Itoa(idbusiness), strconv.Itoa(idbusiness), "EX", 3600)
	if err_do != nil {
		return err_do
	}

	return nil
}
