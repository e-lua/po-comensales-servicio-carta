package cartadiaria_anfitrion

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Delete_Vencidas() error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAMOS EL ELEMENTO
	query_element := `DELETE FROM element USING carta WHERE element.idcarta = carta.idcarta AND carta.deleteddate<=NOW()`
	_, error_element := tx.Exec(ctx, query_element)
	if error_element != nil {
		tx.Rollback(ctx)
		return error_element
	}

	//ELIMINAMOS LA LISTA DE RANGO HORARIOS
	query_listschedulerange := `DELETE FROM listschedulerange USING carta WHERE listschedulerange.idcarta = carta.idcarta AND carta.deleteddate<=NOW()`
	_, error_listschedulerange := tx.Exec(ctx, query_listschedulerange)
	if error_listschedulerange != nil {
		tx.Rollback(ctx)
		return error_listschedulerange
	}

	//ELIMINAMOS EL RANGO HORARIO
	query_schedulerange := `DELETE FROM schedulerange USING carta WHERE schedulerange.idcarta = carta.idcarta AND carta.deleteddate<=NOW()`
	_, error_schedule := tx.Exec(ctx, query_schedulerange)
	if error_schedule != nil {
		tx.Rollback(ctx)
		return error_schedule
	}

	//ELIMINAMOS LA CARTA
	query_carta := `DELETE FROM carta WHERE carta.deleteddate<=NOW()`
	_, error_carta := tx.Exec(ctx, query_carta)
	if error_carta != nil {
		tx.Rollback(ctx)
		return error_carta
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
