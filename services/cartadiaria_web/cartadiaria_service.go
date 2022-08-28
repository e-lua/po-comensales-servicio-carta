package carta_web

import (
	//REPOSITORIES

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria"
	//element_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/element"
)

/*----------------------GET DATA ----------------------*/

func Web_GetBusinessCategory_Service(date string, idbusiness int) (int, bool, string, []models.Pg_Category) {

	//Obtenemos las categorias
	carta_category, error_update := cartadiaria_repository.Pg_Find_Category(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func Web_GetBusinessElement_Service(date string, idbusiness int, limit int) (int, bool, string, []models.V2_Pg_Categories_Elements) {

	carta_elements, error_update := cartadiaria_repository.V2_Pg_Web_Find_Elements(date, idbusiness, limit)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func Web_GetBusinessSchedule_Service(date string, idbusiness int) (int, bool, string, []models.Pg_ScheduleList) {

	//Obtenemos las categorias
	carta_schedule, error_update := cartadiaria_repository.Pg_Find_ScheduleRange(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_schedule
	}

	return 201, false, "", carta_schedule
}

func Web_SearchByNameAndDescription_Service(date string, idbusiness int, name string, limit int) (int, bool, string, []models.Pg_Element_ToCreate) {

	carta_elements, error_find := cartadiaria_repository.Pg_Web_Find_Elements_SearchByText(date, idbusiness, name, limit)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar llos elementos, detalles: " + error_find.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}
