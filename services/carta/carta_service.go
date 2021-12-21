package carta

import (
	//REPOSITORIES

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	carta_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/carta"
	view_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/view"
)

func GetBusinessInformation_Service(idcomensal int, idbusiness int) error {

	//Obtenemos las categorias
	error_add_view := view_repository.Pg_Add_ViewBusiness(idcomensal, idbusiness)
	if error_add_view != nil {
		return error_add_view
	}

	return nil
}

func GetBusinessCategory_Service(date string, idbusiness int) (int, bool, string, []models.Pg_Category) {

	//Obtenemos las categorias
	carta_category, error_update := carta_repository.Pg_Find_Category(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func GetBusinessElement_Service(date string, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_With_Stock) {

	//Obtenemos las categorias
	carta_elements, error_update := carta_repository.Pg_Find_Elements(date, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func GetBusinessSchedule_Service(date string, idbusiness int) (int, bool, string, []models.Pg_ScheduleList) {

	//Obtenemos las categorias
	carta_schedule, error_update := carta_repository.Pg_Find_ScheduleRange(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_schedule
	}

	return 201, false, "", carta_schedule
}

/*----------------------ADD VIEW DATA ELEMENT----------------------*/

func AddViewElement_Service(idcomensal int, idelement int) (int, bool, string, string) {

	//Obtenemos las categorias
	view_repository.Pg_Add_ViewElement(idcomensal, idelement)

	return 201, false, "", "Vista registrada"
}
