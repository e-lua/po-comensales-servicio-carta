package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock(schedule_stock models.Pg_ToSchedule_Mqtt) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE listschedulerange SET maxorders=maxorders-1 WHERE idschedule=$1 AND idcarta=$2`
	if _, err := db.Exec(ctx, query, schedule_stock.IDSchedule, schedule_stock.IDCarta); err != nil {
		return err
	}

	return nil
}
