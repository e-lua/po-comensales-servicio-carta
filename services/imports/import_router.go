package imports

import (
	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/labstack/echo/v4"
)

var ImportsRouter_pg *importsRouter_pg

type importsRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (ir *importsRouter_pg) UpdateElementStock(c echo.Context) error {

	var input_elements []models.Mqtt_Import_ElementStock

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&input_elements)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, detalles: " + err.Error(), Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateElementStock_Service(input_elements)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (ir *importsRouter_pg) UpdateScheduleStock(c echo.Context) error {

	var input_schedule []models.Mqtt_Import_SheduleStock

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&input_schedule)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, detalles: " + err.Error(), Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateScheduleStock_Service(input_schedule)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
