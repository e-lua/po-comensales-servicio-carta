package repositories

import (
	"encoding/json"
	"strconv"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/gomodule/redigo/redis"
)

func Re_Get_Notified(idbusiness int) (int, error) {

	var idbusiness_string string
	var idbusiness_int int

	reply, err := redis.String(models.RedisCN.Get().Do("GET", strconv.Itoa(idbusiness)))
	if err != nil {
		return idbusiness_int, err
	}

	err = json.Unmarshal([]byte(reply), &idbusiness_string)

	idbusiness_int2, _ := strconv.Atoi(idbusiness_string)

	if err != nil {
		return idbusiness_int2, err
	}

	return idbusiness_int2, nil
}
