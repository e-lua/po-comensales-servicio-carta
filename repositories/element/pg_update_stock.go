package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock(elements_stock models.Pg_ToElement_Mqtt) error {

	db := models.Conectar_Pg_DB()

	/*for i := 0; i < elements_stock.IDElement[len(elements_stock.IDElement)-1]; i++ {
		query := `UPDATE Element SET stock=stock-$1 WHERE idelement=$2 AND idcarta=$3`
		if _, err := db.Exec(context.Background(), query, elements_stock.Quantity[i], elements_stock.IDElement[i], elements_stock.IDCarta[i]); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}*/

	query := `UPDATE Element SET stock=stock-ex.stck FROM (select * from  unnest($1::int[], $2::int[],$3::int[])) as ex(stck,idelemt,idcrta) WHERE idelement=ex.idelemt AND idcarta=ex.idcrta`
	if _, err := db.Exec(context.Background(), query, elements_stock.Quantity, elements_stock.IDElement, elements_stock.IDCarta); err != nil {
		return err
	}

	return nil
}
