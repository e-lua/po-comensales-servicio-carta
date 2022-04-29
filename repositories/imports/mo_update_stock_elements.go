package imports

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Mo_Update_Many(input_elements []models.Mqtt_Import_ElementStock) error {

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)

	//defer cancelara el contexto
	defer cancel()

	db := models.MongoCN.Database("comensal_by_anfitrion")
	col := db.Collection("comensales")

	models := []mongo.WriteModel{}

	for _, element := range input_elements {

		updtString := bson.M{
			"$inc": bson.M{
				"stock": -element.Quantity,
			},
		}
		models = append(models,
			mongo.NewUpdateOneModel().SetFilter(
				bson.M{
					"idcarta": element.IDCarta,
					"id":      element.IDElement,
				},
			).
				SetUpdate(updtString).SetUpsert(true),
		)

	}

	opts := options.BulkWrite().SetOrdered(true)

	_, error_update := col.BulkWrite(ctx, models, opts)

	if error_update != nil {
		return error_update
	}

	return nil
}
