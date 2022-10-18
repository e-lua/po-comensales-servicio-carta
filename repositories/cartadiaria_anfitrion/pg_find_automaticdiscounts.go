package cartadiaria_anfitrion

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_AutomaticDiscounts(idcarta int, idbusiness int) ([]models.Pg_V2_AutomaticDiscount, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT idbusiness,idcarta,iddiscount,description,discount,typediscount,groupid,classdiscount FROM automaticdiscount WHERE idcarta=$1 AND idbusiness=$2 ORDER BY discount ASC"
	rows, error_shown := db.Query(ctx, q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListAutomaticDiscount []models.Pg_V2_AutomaticDiscount

	if error_shown != nil {
		return oListAutomaticDiscount, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oAutomaticDiscount models.Pg_V2_AutomaticDiscount
		rows.Scan(&oAutomaticDiscount.IDBusiness, &oAutomaticDiscount.IDCarta, &oAutomaticDiscount.IDAutomaticDiscount, &oAutomaticDiscount.Description, &oAutomaticDiscount.Discount, &oAutomaticDiscount.TypeDiscount, &oAutomaticDiscount.Group, &oAutomaticDiscount.ClassDiscount)
		oListAutomaticDiscount = append(oListAutomaticDiscount, oAutomaticDiscount)
	}

	//Si todo esta bien
	return oListAutomaticDiscount, nil
}
