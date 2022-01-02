package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock(elements_stock models.Pg_ToElement_Mqtt) error {

	db := models.Conectar_Pg_DB()

	query := `UPDATE Element SET stock=stock-ex.stck FROM (select * from  unnest($1::int[], $2::int[],$3::int[])) as ex(stck,idelemt,idcrta) WHERE idelement=ex.idelemt AND idcarta=ex.idcrta`
	if _, err := db.Exec(context.Background(), query, elements_stock.Quantity, elements_stock.IDElement, elements_stock.IDCarta); err != nil {
		return err
	}

	return nil
}
