package carta

import (
	"time"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
)

type Response struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      string `json:"data"`
}

type ResponseBusiness struct {
	Error     bool               `json:"error"`
	DataError string             `json:"dataError"`
	Data      models.Mo_Business `json:"data"`
}

type ResponseJWT struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      JWT    `json:"data"`
}

type JWT struct {
	Phone      int    `json:"phone"`
	Country    int    `json:"country"`
	IDComensal int    ` json:"comensal"`
	Name       string ` json:"name"`
	LastName   string ` json:"lastName"`
}

type ResponseCartaCategory struct {
	Error     bool                 `json:"error"`
	DataError string               `json:"dataError"`
	Data      []models.Pg_Category `json:"data"`
}

type ResponseCartaElements struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.Pg_Element_With_Stock `json:"data"`
}

type ResponseCartaSchedule struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.Pg_ScheduleList `json:"data"`
}

type Send_View_Element struct {
	IDElement  int       `bson:"idelement" json:"idelement"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}
