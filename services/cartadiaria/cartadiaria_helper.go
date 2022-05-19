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
	Data      []*models.Pg_Element_With_Stock_External `json:"data"`
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

/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*-------------------------------------CARTA DIARIA ANFITRIONES-------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/

type Carta struct {
	Date      string `json:"date"`
	WannaCopy bool   `json:"wannacopy"`
	FromCarta string `json:"fromcarta"`
}

type CartaStatus struct {
	IDCarta   int  `json:"idcarta"`
	Available bool `json:"available"`
	Visible   bool `json:"visible"`
}

type CartaSchedule struct {
	IDCarta        int                                `json:"idcarta"`
	ScheduleRanges []models.Pg_ScheduleRange_External `json:"schedule"`
}

type ResponseInt struct {
	Error     bool   `json:"error"`
	DataError string `json:"dataError"`
	Data      int    `json:"data"`
}

type ResponseCartaBasicData struct {
	Error     bool                     `json:"error"`
	DataError string                   `json:"dataError"`
	Data      models.Pg_Carta_External `json:"data"`
}

type ResponseCartaCategory_Ext struct {
	Error     bool                          `json:"error"`
	DataError string                        `json:"dataError"`
	Data      []models.Pg_Category_External `json:"data"`
}

type ResponseCartaElements_Ext struct {
	Error     bool                                    `json:"error"`
	DataError string                                  `json:"dataError"`
	Data      []models.Pg_Element_With_Stock_External `json:"data"`
}

type ResponseCartaSchedule_Ext struct {
	Error     bool                               `json:"error"`
	DataError string                             `json:"dataError"`
	Data      []models.Pg_ScheduleRange_External `json:"data"`
}

type ResponseCartas struct {
	Error     bool                    `json:"error"`
	DataError string                  `json:"dataError"`
	Data      []models.Pg_Carta_Found `json:"data"`
}

//ADDRESS
type ResponseAddress struct {
	Error     bool      `json:"error"`
	DataError string    `json:"dataError"`
	Data      B_Address `json:"data"`
}

type B_Address struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

/*===============TESTEANDO===============*/

type CartaElements_WithAction struct {
	IDCarta            int                                     `json:"idcarta"`
	ElementsWithAction []models.Pg_Element_With_Stock_External `json:"elements"`
}
