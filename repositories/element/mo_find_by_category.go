package element

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Find_ByCategory(date string, idbusiness int, idcategory int) ([]*models.Mo_Element_With_Stock_Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_inventory")
	col := db.Collection("storehouse")

	/*Aca pude haber hecho un make, es decir, resultado:=make([]...)*/
	var resultado []*models.Mo_Element_With_Stock_Response

	condicion := bson.M{
		"idbusiness": idbusiness,
		"date":       date,
		"idcategory": idcategory,
	}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "namecategory", Value: 1}})

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Mo_Element_With_Stock_Response
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
