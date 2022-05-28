package element

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Search_Name_Comensales(date string, idbusiness int, text string, limit int64, offset int64) ([]*models.Pg_Element_ToCreate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_cartadiaria")
	col := db.Collection("elements")

	var resultado []*models.Pg_Element_ToCreate

	condicion := bson.M{
		"idbusiness": idbusiness,
		"date":       date,
		"name": primitive.Regex{
			Pattern: text,
			Options: "i",
		},
	}

	opciones := options.Find()
	/*Indicar como ira ordenado*/
	opciones.SetSort(bson.D{{Key: "name", Value: 1}})
	opciones.SetSkip((offset - 1) * limit)

	/*Cursor es como una tabla de base de datos donde se van a grabar los resultados
	y podre ir recorriendo 1 a la vez*/
	cursor, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return resultado, err
	}

	//contexto, en este caso, me crea un contexto vacio
	for cursor.Next(context.TODO()) {
		/*Aca trabajare con cada Tweet. El resultado lo grabará en registro*/
		var registro models.Pg_Element_ToCreate
		err := cursor.Decode(&registro)
		if err != nil {
			return resultado, err
		}
		/*Recordar que Append sirve para añadir un elemento a un slice*/
		resultado = append(resultado, &registro)
	}

	return resultado, nil
}
