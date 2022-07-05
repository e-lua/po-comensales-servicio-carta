package cartadiaria_anfitrion

import (
	"context"
	"math/rand"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

func Pg_Find_Elements(idcarta int, idbusiness int) ([]models.Pg_Element_With_Stock_External, error) {

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

	q := "SELECT idelement,idcarta,idbusiness,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,typefood,insumos,availableorders,costo FROM element WHERE idcarta=$1 AND idbusiness=$2 ORDER BY stock ASC"
	rows, error_shown := db.Query(ctx, q, idcarta, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oListElementsWithStock []models.Pg_Element_With_Stock_External

	if error_shown != nil {

		return oListElementsWithStock, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oElement models.Pg_Element_With_Stock_External
		rows.Scan(&oElement.IDElement, &oElement.IDCarta, &oElement.IDBusiness, &oElement.IDCategory, &oElement.NameCategory, &oElement.UrlPhotoCategory, &oElement.Name, &oElement.Price, &oElement.Description, &oElement.UrlPhoto, &oElement.TypeMoney, &oElement.Stock, &oElement.Typefood, &oElement.Insumos, &oElement.AvailableOrders, &oElement.Costo)
		oListElementsWithStock = append(oListElementsWithStock, oElement)
	}

	//Si todo esta bien
	return oListElementsWithStock, nil
}
