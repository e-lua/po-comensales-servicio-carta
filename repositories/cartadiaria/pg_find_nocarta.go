package cartadiaria

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_NoCarta() ([]int, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT idbusiness FROM carta WHERE date=now()::date"
	rows, error_shown := db.Query(ctx, q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListIdBusiness []int
	quantity := 0

	if error_shown != nil {
		quantity = quantity + 1
		return oListIdBusiness, quantity, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oIdBusiness int
		rows.Scan(&oIdBusiness)
		oListIdBusiness = append(oListIdBusiness, oIdBusiness)
	}

	//Si todo esta bien
	return oListIdBusiness, quantity, nil
}
