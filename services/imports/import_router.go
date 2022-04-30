package imports

import (
	"log"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

var ImportsRouter_pg *importsRouter_pg

type importsRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (ir *importsRouter_pg) UpdateElementStock(input_elements []models.Mqtt_Import_ElementStock) {

	//Enviamos los datos al servicio
	error_update := UpdateElementStock_Service(input_elements)
	if error_update != nil {
		log.Println("Error al intentar actualizar el stock de elementos, detalles ", error_update.Error())
	}
}

func (ir *importsRouter_pg) UpdateScheduleStock(input_schedule []models.Mqtt_Import_SheduleStock) {

	//Enviamos los datos al servicio
	error_update := UpdateScheduleStock_Service(input_schedule)
	if error_update != nil {
		log.Println("Error al intentar actualizar el stock de schedule, detalles ", error_update.Error())
	}
}
