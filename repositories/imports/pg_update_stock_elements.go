package imports

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Update_Stock_Element(input_elements []models.Mqtt_Import_ElementStock) error {

	idelement_pg, idcarta_pg, quantity_pg := []int64{}, []int{}, []int{}

	for _, element := range input_elements {
		idelement_pg = append(idelement_pg, element.IDElement)
		idcarta_pg = append(idcarta_pg, element.IDCarta)
		quantity_pg = append(quantity_pg, element.Quantity)
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	query := `UPDATE Element SET stock=stock-ex.stck FROM (select * from  unnest($1::bigint[],$2::int[],$3::int[])) as ex(idelemt,idcrta,stck) WHERE idelement=ex.idelemt AND idcarta=ex.idcrta`
	if _, err := db.Exec(ctx, query, idelement_pg, idcarta_pg, quantity_pg); err != nil {
		return err
	}

	return nil
}
