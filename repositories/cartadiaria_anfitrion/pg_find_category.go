package cartadiaria_anfitrion

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_Category(idcarta int, idbusiness int) ([]models.Pg_Category_External, error) {

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

	q := "SELECT idcategory,namecategory,urlphotcategory,COUNT(idelement) FROM Element WHERE idcarta=$1 AND idbusiness=$2 GROUP BY idcategory,namecategory,urlphotcategory ORDER BY namecategory ASC"
	rows, error_shown := db.Query(ctx, q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oLisstCategory []models.Pg_Category_External

	if error_shown != nil {

		return oLisstCategory, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oCategory models.Pg_Category_External
		rows.Scan(&oCategory.IDCategory, &oCategory.Name, &oCategory.UrlPhoto, &oCategory.AmountOfElements)
		oLisstCategory = append(oLisstCategory, oCategory)
	}

	//Si todo esta bien
	return oLisstCategory, nil
}
