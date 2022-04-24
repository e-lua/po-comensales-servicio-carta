package carta

import (
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

type ResponseBusiness_V2 struct {
	Error     bool                  `json:"error"`
	DataError string                `json:"dataError"`
	Data      models.Mo_Business_V2 `json:"data"`
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

type ResponseCartaElements_Searched struct {
	Error     bool                           `json:"error"`
	DataError string                         `json:"dataError"`
	Data      []models.Pg_Element_With_Stock `json:"data"`
}

type ResponseCartaElements_Searched_Mo struct {
	Error     bool                                     `json:"error"`
	DataError string                                   `json:"dataError"`
	Data      []*models.Mo_Element_With_Stock_Response `json:"data"`
}

type ResponseCartaSchedule struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      []models.Pg_ScheduleList `json:"data"`
}

type JWT_Anfitrion struct {
	IdBusiness int `json:"idBusiness"`
	IdWorker   int `json:"idWorker"`
	IdCountry  int `json:"country"`
	IdRol      int `json:"rol"`
}

type ResponseJWT_Anfitrion struct {
	Error     bool          `json:"error"`
	DataError string        `json:"dataError"`
	Data      JWT_Anfitrion `json:"data"`
}

type ResponseCartaCategory_ToCreate struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []models.Pg_Category_ToCreate `json:"data"`
}

type ResponseCartaElements_ToCreate struct {
	Error     bool                         `json:"error"`
	DataError string                       `json:"dataError"`
	Data      []models.Pg_Element_ToCreate `json:"data"`
}

type ResponseCartaSchedule_ToCreate struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []models.Pg_Schedule_ToCreate `json:"data"`
}
