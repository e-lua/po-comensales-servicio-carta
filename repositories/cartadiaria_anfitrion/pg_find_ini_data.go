package cartadiaria_anfitrion

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_IniData(date string, idbusiness int) (models.Pg_Carta_External, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	var carta_ini_data models.Pg_Carta_External

	var db *pgxpool.Pool

	random := rand.Intn(4)
	if random%2 == 0 {
		db = models.Conectar_Pg_DB()
	} else {
		db = models.Conectar_Pg_DB_Slave()
	}

	q := "SELECT idcarta,date,availableorders,visible,(SELECT COUNT(*) FROM Element e JOIN Carta c on e.idcarta=c.idcarta WHERE c.date=$1 AND c.idbusiness=$2),(SELECT COUNT(*) FROM ScheduleRange sr JOIN Carta c on sr.idcarta=c.idcarta WHERE c.date=$1 AND c.idbusiness=$2) FROM Carta WHERE date=$1 AND idbusiness=$2"
	error_shown := db.QueryRow(ctx, q, date, idbusiness).Scan(&carta_ini_data.IDCarta, &carta_ini_data.Date, &carta_ini_data.AvailableForOrders, &carta_ini_data.Visible, &carta_ini_data.Elements, &carta_ini_data.ScheduleRanges)

	if error_shown != nil {
		return carta_ini_data, error_shown
	}

	//Si todo esta bien
	return carta_ini_data, nil
}
