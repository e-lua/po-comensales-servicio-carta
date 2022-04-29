package imports

import (
	//REPOSITORIES

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	imports_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/imports"
)

func UpdateElementStock_Service(input_elements []models.Mqtt_Import_ElementStock) (int, bool, string, string) {

	error_update := imports_repository.Pg_Update_Stock_Element(input_elements)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el stock" + error_update.Error(), ""
	}

	error_update_mo := imports_repository.Mo_Update_Many(input_elements)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el stock_mo" + error_update_mo.Error(), ""
	}

	return 200, false, "", "Actualización correcta"
}

func UpdateScheduleStock_Service(input_schedule []models.Mqtt_Import_SheduleStock) (int, bool, string, string) {

	error_update := imports_repository.Pg_Update_Stock_Schedule(input_schedule)
	if error_update != nil {
		return 500, true, "Error en el servidor interno al intentar actualizar el stock" + error_update.Error(), ""
	}

	return 200, false, "", "Actualización correcta"
}
