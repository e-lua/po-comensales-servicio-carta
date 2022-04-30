package imports

import (
	//REPOSITORIES

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	imports_repository "github.com/Aphofisis/po-comensales-servicio-carta/repositories/imports"
)

func UpdateElementStock_Service(input_elements []models.Mqtt_Import_ElementStock) error {

	error_update := imports_repository.Pg_Update_Stock_Element(input_elements)
	if error_update != nil {
		return error_update
	}

	error_update_mo := imports_repository.Mo_Update_Many(input_elements)
	if error_update != nil {
		return error_update_mo
	}

	return nil
}

func UpdateScheduleStock_Service(input_schedule []models.Mqtt_Import_SheduleStock) error {

	error_update := imports_repository.Pg_Update_Stock_Schedule(input_schedule)
	if error_update != nil {
		return error_update
	}

	return nil
}
