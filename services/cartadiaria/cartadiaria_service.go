package carta

import (
	//REPOSITORIES

	"log"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria"
	element_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/element"
	schedule_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/schedule"
)

/*----------------------UDPATE DATA CONSUME----------------------*/

func UpdateElementStock_Service(input_elements models.Pg_ToElement_Mqtt) error {

	//Obtenemos las categorias
	error_add_view := element_repository.Pg_Update_Stock(input_elements)
	if error_add_view != nil {
		return error_add_view
	}

	return nil
}
func UpdateScheduleStock_Service(input_schedule models.Pg_ToSchedule_Mqtt) error {

	//Obtenemos las categorias
	error_add_view := schedule_repository.Pg_Update_Stock(input_schedule)
	if error_add_view != nil {
		return error_add_view
	}

	return nil
}

/*----------------------GET DATA ----------------------*/

func GetBusinessCategory_Service(date string, idbusiness int) (int, bool, string, []models.Pg_Category) {

	//Obtenemos las categorias
	carta_category, error_update := cartadiaria_repository.Pg_Find_Category(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func AddViewInformation_Service(idelement int, idcomensal int) (int, bool, string, string) {

	element_repository.Pg_ExportView(idelement, idcomensal)

	return 200, false, "", "Vista registrada"
}

func GetBusinessElement_Service(date string, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_With_Stock) {

	//Obtenemos las categorias
	carta_elements, error_update := cartadiaria_repository.Pg_Find_Elements(date, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func SearchByNameAndDescription_Service(date string, idbusiness int, text string, limit int, offset int) (int, bool, string, []*models.Mo_Element_With_Stock_Response) {

	//Version PG

	/*carta_elements, error_find := cartadiaria_repository.Pg_Find_Elements_SearchByText(date, idbusiness, text, limit, offset)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar llos elementos, detalles: " + error_find.Error(), carta_elements
	}*/

	//Version MO

	carta_elements, error_find := element_repository.Mo_Search_Name(date, idbusiness, text, int64(limit), int64(offset))
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar llos elementos, detalles: " + error_find.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func GetBusinessSchedule_Service(date string, idbusiness int) (int, bool, string, []models.Pg_ScheduleList) {

	//Obtenemos las categorias
	carta_schedule, error_update := cartadiaria_repository.Pg_Find_ScheduleRange(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_schedule
	}

	return 201, false, "", carta_schedule
}

/*-------------------------------------ELEMENTS-------------------------------------*/

func UpdateCarta_ElementsWithStock_Service(input_mqtt_elements models.Mqtt_Element_With_Stock_Import) error {

	//Insertamos los datos en PG
	error_adelete_update := element_repository.Mo_Delete_Update(input_mqtt_elements)
	if error_adelete_update != nil {
		log.Fatal(error_adelete_update)
	}

	return nil
}
