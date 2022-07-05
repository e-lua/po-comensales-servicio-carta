package cartadiaria_anfitrion

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_SearchToNotify() ([]int, int, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT  car.idbusiness FROM carta as car JOIN element as ele ON car.idcarta=ele.idcarta JOIN schedulerange as sch ON car.idcarta=sch.idcarta WHERE (car.date)::date=(now() at time zone CONCAT('utc',(sch.timezone::integer*-1)::varchar(5)))::date GROUP BY car.idbusiness"
	rows, error_shown := db.Query(ctx, q)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListBusiness []int
	quantity := 0

	if error_shown != nil {

		return oListBusiness, quantity, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oBusiness int
		rows.Scan(&oBusiness)
		oListBusiness = append(oListBusiness, oBusiness)
		quantity = quantity + 1
	}

	//Si todo esta bien
	return oListBusiness, quantity, nil

}
