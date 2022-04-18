package cartadiaria

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_Category(date string, idbusiness int) ([]models.Pg_Category, error) {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db := models.Conectar_Pg_DB()

	q := "SELECT c.idcarta,c.availableorders,e.idcategory,e.namecategory,e.urlphotcategory,COUNT(e.idelement) FROM Element e LEFT JOIN Carta c ON e.idcarta=c.idcarta WHERE c.date=$1 AND e.idbusiness=$2 GROUP BY c.idcarta,e.idcategory,e.namecategory,e.urlphotcategory,c.availableorders ORDER BY e.namecategory ASC"
	rows, error_shown := db.Query(ctx, q, date, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oLisstCategory []models.Pg_Category

	if error_shown != nil {

		return oLisstCategory, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oCategory models.Pg_Category
		rows.Scan(&oCategory.IDCarta, &oCategory.Availableorders, &oCategory.IDCategory, &oCategory.Name, &oCategory.UrlPhoto, &oCategory.AmountOfElements)
		oLisstCategory = append(oLisstCategory, oCategory)
	}

	//Si todo esta bien
	return oLisstCategory, nil
}
