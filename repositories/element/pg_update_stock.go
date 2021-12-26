package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock(elements_stock models.Pg_ToElement_Mqtt) error {

	db := models.Conectar_Pg_DB()

	for i := 0; i < elements_stock.IDElement[len(elements_stock.IDElement)-1]; i++ {
		query := `UPDATE OrderDetails SET quantity=quantity-$1 WHERE idelement=$2 AND idcarta=$3`
		if _, err := db.Exec(context.Background(), query, elements_stock.Quantity[i], elements_stock.IDElement[i], elements_stock.IDCarta[i]); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}
