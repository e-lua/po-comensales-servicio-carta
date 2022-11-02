package carta

import (
	//REPOSITORIES

	"bytes"
	"encoding/json"
	"net/http"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria"
	cartadiaria_anfitrion_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/cartadiaria_anfitrion"
	element_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/element"
	schedule_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/schedule"
)

/*----------------------AUTOMATIC NOTIFY----------------------*/

func Find__Notify_NoCarta_Service() error {

	ahora := time.Now()

	if ahora.Hour() == 14 {

		//Obtenemos las categorias
		/*list_idbusiness, quantity, error_nocarta := cartadiaria_repository.Pg_Find_NoCarta()
		if error_nocarta != nil {
			return error_nocarta
		}

		if quantity > 0 {

			notification := map[string]interface{}{
				"message":      "No olvide programar la carta para el día de hoy en la sección Carta Diaria",
				"multipleuser": list_idbusiness,
				"typeuser":     6,
				"priority":     1,
				"title":        "Restoner anfitriones",
			}
			json_data, _ := json.Marshal(notification)
			http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))

			return nil
		}*/

		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "No olvide programar la carta para el día de hoy en la sección Carta Diaria",
			"iduser":   0,
			"typeuser": 4,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
		return nil

	}

	return nil
}

/*---------------------------------------------------------------*/

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

func GetBusinessElement_Service(date string, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_ToCreate) {

	//Obtenemos las categorias
	carta_elements, error_update := cartadiaria_repository.Pg_Find_Elements(date, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func SearchByNameAndDescription_Service(date string, idbusiness int, text string, limit int, offset int) (int, bool, string, []models.Pg_Element_ToCreate) {

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

func GetBusinessElement_ListByCategory_Service(date string, idbusiness int, limit int) (int, bool, string, []interface{}) {

	carta_elements, error_update := cartadiaria_repository.V2_Pg_Web_Find_Elements(date, idbusiness, limit)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_elements
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

/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*-------------------------------------CARTA DIARIA ANFITRIONES-------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/

func SearchByName_Anfitrion_Service(date string, idbusiness int, text string, limit int, offset int) (int, bool, string, []models.Pg_Element_ToCreate) {
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

func GetElementsByInsumo_Service(date string, idbusiness int, idinsumo string) (int, bool, string, []*models.Mqtt_Element_With_Stock) {

	//Obtenemos los elementos
	elementos, error_update := element_repository.Mo_Find_ByInsumo(date, idbusiness, idinsumo)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos por el insumo indicado, detalles: " + error_update.Error(), elementos
	}

	return 201, false, "", elementos
}

func AddCarta_Service(input_carta Carta, idbusiness int) (int, bool, string, int) {

	//Insertamos los datos en Mo
	idcarta, error_add_carta := cartadiaria_anfitrion_repository.Pg_Add(idbusiness, input_carta.Date)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	go func() {

		var nameday string

		t, _ := time.Parse("2006-01-02", input_carta.Date)

		switch int(t.Weekday()) {
		case 0:
			nameday = "Domingo "
		case 1:
			nameday = "Lunes "
		case 2:
			nameday = "Martes "
		case 3:
			nameday = "Miercoles "
		case 4:
			nameday = "Jueves "
		case 5:
			nameday = "Viernes "
		default:
			nameday = "Sabado "
		}

		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se creó una nueva carta para el día " + nameday + string(input_carta.Date[8:]) + "/" + string(input_carta.Date[5]) + string(input_carta.Date[6]) + "/" + string(input_carta.Date[:4]),
			"iduser":   idbusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 201, false, "", idcarta
}

func UpdateCartaStatus_Service(carta_status CartaStatus, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := cartadiaria_anfitrion_repository.Pg_Update_Available_Visible(carta_status.Available, carta_status.Visible, carta_status.IDCarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar la visibilidad y disponibilidad de la carta , detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "La disponibilidad y visibilidad se actualizaron correctamente"
}

func UpdateCartaOneElement_Service(stock int, idelement int, idcarta int, idbusiness int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_update := cartadiaria_anfitrion_repository.Pg_Update_One_Element(stock, idelement, idcarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el elemento, detalles: " + error_update.Error(), ""
	}

	return 201, false, "", "Elemento actualizado correctamente"
}

func UpdateCartaElements_Service(carta_elements CartaElements_WithAction, idbusiness int, latitude float64, longitude float64) (int, bool, string, string) {

	error_update := cartadiaria_anfitrion_repository.Pg_Delete_Update_Element(carta_elements.ElementsWithAction, carta_elements.IDCarta, idbusiness, latitude, longitude)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los elementos, detalles: " + error_update.Error(), ""
	}

	//Registramos los datos en Mongo DB
	error_update_mo := cartadiaria_anfitrion_repository.Mo_Delete_Update_Elements(carta_elements.ElementsWithAction, carta_elements.IDCarta, idbusiness)
	if error_update_mo != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los elementos, detalles: " + error_update_mo.Error(), ""
	}

	return 201, false, "", "Los elementos se actualizaron correctamente"
}

func UpdateCartaScheduleRanges_Service(carta_schedule CartaSchedule, idbusiness int) (int, bool, string, string) {

	error_update := cartadiaria_anfitrion_repository.Pg_Delete_Update_ScheduleRange(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update.Error(), ""
	}

	//Registramos los datos en Mongo DB
	/*error_update_mo := cartadiaria_anfitrion_repository.Mo_Delete_Update_Schedule(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
	if error_update_mo != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update_mo.Error(), ""
	}*/

	return 201, false, "", "Los rangos horario se actualizaron correctamente"
}

func UpdateCartaAutomaticDiscounts_Service(carta_automaticdiscount CartaAutomaticDiscount, idbusiness int) (int, bool, string, string) {

	error_update := cartadiaria_anfitrion_repository.Pg_Delete_Update_AutomaticDiscount(carta_automaticdiscount.AutomaticDiscount, carta_automaticdiscount.IDCarta, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los descuentos automaticos, detalles: " + error_update.Error(), ""
	}

	//Registramos los datos en Mongo DB
	/*error_update_mo := cartadiaria_anfitrion_repository.Mo_Delete_Update_Schedule(carta_schedule.ScheduleRanges, carta_schedule.IDCarta, idbusiness)
	if error_update_mo != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update_mo.Error(), ""
	}*/

	return 201, false, "", "Los descuentos automaticos, se actualizaron correctamente"
}

/*----------------------GET DATA OF MENU----------------------*/

func GetCartaBasicData_Service(date string, idbusiness int) (int, bool, string, models.Pg_Carta_External) {

	//Insertamos los datos en Mo
	carta_ini_values, error_show := cartadiaria_anfitrion_repository.Pg_Find_IniData(date, idbusiness)
	if error_show != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar la informacion basica de la carta, detalles: " + error_show.Error(), carta_ini_values
	}

	return 201, false, "", carta_ini_values
}

func GetCartaCategory_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_Category_External) {

	//Obtenemos las categorias
	carta_category, error_update := cartadiaria_anfitrion_repository.Pg_Find_Category(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar las categorias de la carta, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func GetCartaElementsByCarta_Service(idcarta_int int, idbusiness int, idcategory int) (int, bool, string, []models.Pg_Element_With_Stock_External) {

	//Obtenemos las categorias
	carta_category, error_update := cartadiaria_anfitrion_repository.Pg_Find_Elements_ByCategory(idcarta_int, idbusiness, idcategory)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la categoria seleccionada, detalles: " + error_update.Error(), carta_category
	}

	return 201, false, "", carta_category
}

func GetCartaElements_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_Element_With_Stock_External) {

	//Insertamos los datos en Mo
	carta_elements, error_update := cartadiaria_anfitrion_repository.Pg_Find_Elements(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update.Error(), carta_elements
	}

	return 201, false, "", carta_elements
}

func GetCartaScheduleRanges_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_ScheduleRange_External) {

	//Insertamos los datos en Mo
	carta_scheduleranges, error_update := cartadiaria_anfitrion_repository.Pg_Find_ScheduleRanges(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update.Error(), carta_scheduleranges
	}

	return 201, false, "", carta_scheduleranges
}

func GetCartAutomaticDiscounts_Service(idcarta_int int, idbusiness int) (int, bool, string, []models.Pg_V2_AutomaticDiscount) {

	//Insertamos los datos en Mo
	carta_automaticdiscounts, error_update := cartadiaria_anfitrion_repository.Pg_Find_AutomaticDiscounts(idcarta_int, idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los descuentos automaticos de la carta, detalles: " + error_update.Error(), carta_automaticdiscounts
	}

	return 201, false, "", carta_automaticdiscounts
}

func GetCartas_Service(idbusiness int) (int, bool, string, []models.Pg_Carta_Found) {

	//Insertamos los datos en Mo
	carta_found, error_update := cartadiaria_anfitrion_repository.Pg_Find_Cartas(idbusiness)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update.Error(), carta_found
	}

	return 201, false, "", carta_found
}

/*----------------------COPY BETWEEN MENUS----------------------*/

func AddCartaFromOther_Service(input_carta Carta, idbusiness int) (int, bool, string, int) {

	//Buscamos la carta
	idcarta_int, error_add_carta := cartadiaria_anfitrion_repository.Pg_Find_IniData(input_carta.FromCarta, idbusiness)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	carta_elements, error_update_element := cartadiaria_anfitrion_repository.Pg_Find_Elements(idcarta_int.IDCarta, idbusiness)
	if error_update_element != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los elementos de la carta, detalles: " + error_update_element.Error(), 0
	}

	carta_scheduleranges, error_update_schedule := cartadiaria_anfitrion_repository.Pg_Find_ScheduleRanges(idcarta_int.IDCarta, idbusiness)
	if error_update_schedule != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los rangos horarios de la carta, detalles: " + error_update_schedule.Error(), 0
	}

	carta_automaticdiscounts, error_update_schedule := cartadiaria_anfitrion_repository.Pg_Find_AutomaticDiscounts(idcarta_int.IDCarta, idbusiness)
	if error_update_schedule != nil {
		return 500, true, "Error en el servidor interno al intentar encontrar los los descuentos automaticos de la carta, detalles: " + error_update_schedule.Error(), 0
	}

	//Creamos la carta
	idcarta, error_add_carta := cartadiaria_anfitrion_repository.Pg_Add(idbusiness, input_carta.Date)
	if error_add_carta != nil {
		return 500, true, "Error en el servidor interno al intentar crear la carta, detalles: " + error_add_carta.Error(), 0
	}

	//Transaccion
	id_carta, error_update_schedulelist := cartadiaria_anfitrion_repository.Pg_Copy_Carta(carta_automaticdiscounts, carta_scheduleranges, carta_elements, idbusiness, input_carta.Date, idcarta)
	if error_update_schedulelist != nil {
		return 500, true, "Error en el servidor interno al intentar agregar los elementos y rangos horarios, detalles: " + error_update_schedulelist.Error(), 0
	}

	//Registramos los datos en Mongo DB
	error_update_mo := cartadiaria_anfitrion_repository.Mo_Delete_Update_Elements(carta_elements, 0, idbusiness)
	if error_update_mo != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los elementos, detalles: " + error_update_mo.Error(), 0
	}

	//Registramos los datos en Mongo DB
	error_update_mo_sch := cartadiaria_anfitrion_repository.Mo_Delete_Update_Schedule(carta_scheduleranges, 0, idbusiness)
	if error_update_mo != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar los rangos horarios, detalles: " + error_update_mo_sch.Error(), 0
	}

	/*--SENT NOTIFICATION--*/
	var nameday string

	t, _ := time.Parse("2006-01-02", input_carta.Date)

	switch int(t.Weekday()) {
	case 0:
		nameday = "Domingo "
	case 1:
		nameday = "Lunes "
	case 2:
		nameday = "Martes "
	case 3:
		nameday = "Miercoles "
	case 4:
		nameday = "Jueves "
	case 5:
		nameday = "Viernes "
	default:
		nameday = "Sabado "
	}

	notification := map[string]interface{}{
		"message":  "Se copió la carta del " + string(input_carta.FromCarta[8:]) + "/" + string(input_carta.FromCarta[5]) + string(input_carta.FromCarta[6]) + "/" + string(input_carta.FromCarta[:4]) + " para crear una nueva carta para el día " + nameday + string(input_carta.Date[8:]) + "/" + string(input_carta.Date[5]) + string(input_carta.Date[6]) + "/" + string(input_carta.Date[:4]),
		"iduser":   idbusiness,
		"typeuser": 1,
		"priority": 1,
		"title":    "Restoner anfitriones",
	}
	json_data, _ := json.Marshal(notification)
	http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
	/*---------------------*/

	return 201, false, "", id_carta
}

/*----------------------DELETE MENU----------------------*/

func DeleteCarta_Service(idbusiness int, idcarta int) (int, bool, string, string) {

	//Insertamos los datos en Mo
	error_delete := cartadiaria_anfitrion_repository.Pg_Delete(idbusiness, idcarta)
	if error_delete != nil {
		return 500, true, "Error en el servidor interno al intentar eliminar la carta, detalles: " + error_delete.Error(), ""
	}

	go func() {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Se eliminó una carta",
			"iduser":   idbusiness,
			"typeuser": 1,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}()

	return 201, false, "", "Eliminado correctamente"
}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func SearchToNotifyCarta_Service() (int, bool, string, []int) {

	//Agregamos la categoria
	all_business, quantity, error_add := cartadiaria_anfitrion_repository.Pg_SearchToNotify()
	if error_add != nil {
		return 500, true, "Error en el servidor interno al intentar listar los negocios con datos de la carta a no notificar, detalles: " + error_add.Error(), all_business
	}

	if quantity > 0 {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":      "Recuerde crear la carta para el día hoy. Cree la carta desde cero, o copie la del día anterior, pero tenga en cuenta revisar el stock de elementos disponibles de la nueva carta, y habilitar la visibilidad",
			"multipleuser": all_business,
			"typeuser":     6,
			"priority":     1,
			"title":        "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	} else {
		/*--SENT NOTIFICATION--*/
		notification := map[string]interface{}{
			"message":  "Recuerde crear la carta para el día hoy. Cree la carta desde cero, o copie la del día anterior, pero tenga en cuenta revisar el stock de elementos disponibles de la nueva carta, y habilitar la visibilidad",
			"typeuser": 4,
			"priority": 1,
			"title":    "Restoner anfitriones",
		}
		json_data, _ := json.Marshal(notification)
		http.Post("http://c-a-notificacion-tip.restoner-api.fun:5800/v1/notification", "application/json", bytes.NewBuffer(json_data))
		/*---------------------*/
	}

	return 201, false, "", all_business
}

/*----------------------DELETE----------------------*/

func Delete_Vencidas_Service() (string, string) {

	error_update := cartadiaria_anfitrion_repository.Pg_Delete_Vencidas()
	if error_update != nil {
		return "Error en el servidor interno al intentar eliminar las cartas vencidas, detalles: " + error_update.Error(), ""
	}

	return "", "Cartas vencidas eliminadas correctamente"
}

/*-------------------------------------CARTA DIARIA ANFITRIONES <==> CREAR ORDEN-------------------------------------*/

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
