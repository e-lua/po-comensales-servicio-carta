package carta

import (
	//REPOSITORIES

	"log"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria"
	cartadiaria_anfitrion_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria_anfitrion"
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

func SearchByNameAndDescription_Service(date string, idbusiness int, text string, limit int, offset int) (int, bool, string, []models.Pg_Element_With_Stock) {

	//Version PG

	carta_elements, error_find := cartadiaria_repository.Pg_Find_Elements_SearchByText(date, idbusiness, text, limit, offset)
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar llos elementos, detalles: " + error_find.Error(), carta_elements
	}

	//Version MO

	/*carta_elements, error_find := element_repository.Mo_Search_Name(date, idbusiness, text, int64(limit), int64(offset))
	if error_find != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar llos elementos, detalles: " + error_find.Error(), carta_elements
	}*/

	return 201, false, "", carta_elements
}

func SearchByName_Anfitrion_Service(date string, idbusiness int, text string, limit int, offset int) (int, bool, string, []*models.Mo_Element_With_Stock_Response) {

	//Version MO

	carta_elements, error_find := element_repository.Mo_Search_Name_Anfitriones(date, idbusiness, text, int64(limit), int64(offset))
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

func GetElementsByInsumo_Service(date string, idbusiness int, idinsumo string) (int, bool, string, []*models.Mo_Element_With_Stock_Response) {

	//Obtenemos los elementos
	elementos, error_update := element_repository.Mo_Find_ByInsumo(date, idbusiness, idinsumo)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos por el insumo indicado, detalles: " + error_update.Error(), elementos
	}

	return 201, false, "", elementos
}

/*-------------------------------------CARTA DIARIA ANFITRIONES-------------------------------------*/

func GetCategories_ToCreateOrder_Service(date string, idbusiness int) (int, bool, string, []models.Pg_Category_ToCreate) {

	//Obtenemos las categorias
	category_tocreate, error_update := cartadiaria_anfitrion_repository.Pg_Find_Category_ToCreate(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), category_tocreate
	}

	return 201, false, "", category_tocreate
}

func GetElements_ToCreateOrder_Service(date string, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_ToCreate) {

	//Obtenemos las categorias
	elements_tocreate, error_update := cartadiaria_anfitrion_repository.Pg_Find_Elements_ToCreate(date, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update.Error(), elements_tocreate
	}

	return 201, false, "", elements_tocreate
}

func GetSchedule_ToCreateOrder_Service(date string, idbusiness int) (int, bool, string, []models.Pg_Schedule_ToCreate) {

	//Obtenemos las categorias
	schedule_tocreate, error_update := cartadiaria_anfitrion_repository.Pg_Find_ScheduleRange_ToCreate(date, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), schedule_tocreate
	}

	return 201, false, "", schedule_tocreate
}
