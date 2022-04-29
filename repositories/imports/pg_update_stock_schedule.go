package imports

import (
	"context"
	"time"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock_Schedule(input_schedules []models.Mqtt_Import_SheduleStock) error {

	idschedule_pg, idcarta_pg, quantity_pg := []int64{}, []int{}, []int{}

	for _, schedule := range input_schedules {
		idschedule_pg = append(idschedule_pg, schedule.Schedule)
		idcarta_pg = append(idcarta_pg, schedule.IDCarta)
		quantity_pg = append(quantity_pg, schedule.Quantity)
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE listschedulerange SET maxorders=maxorders-ex.stck FROM (select * from  unnest($1::bigint[],$2::int[],$3::int[])) as ex(idsche,idcrta,stck) WHERE idschedule=ex.idsche AND idcarta=ex.idcrta`
	if _, err := db.Exec(ctx, query, idschedule_pg, idcarta_pg, quantity_pg); err != nil {
		return err
	}

	return nil
}
