package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_Elements(date string, idbusiness int, idcategory int) ([]models.Pg_Element_With_Stock, error) {

	db := models.Conectar_Pg_DB()

	q := "SELECT e.idelement,e.idbusiness,e.idcategory,e.namecategory,e.urlphotcategory,e.name,e.price,e.description,e.urlphoto,e.typemoney,e.stock FROM element e LEFT JOIN carta c ON e.idbusiness=c.idbusiness WHERE c.date=$1 AND e.idbusiness=$2 AND e.idcategory=$3 ORDER BY stock ASC"
	rows, error_shown := db.Query(context.Background(), q, date, idbusiness, idcategory)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_With_Stock

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_With_Stock
		rows.Scan(&oElement.IDElement, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
