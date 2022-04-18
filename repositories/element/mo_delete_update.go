package element

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Mo_Delete_Update(input_mqtt_elements models.Mqtt_Element_With_Stock_Import) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	var array_elements []interface{}
	longitude_of_array := len(input_mqtt_elements.Elements_with_stock)

	for i := 0; i <= longitude_of_array; i++ {
		array_elements[i] = input_mqtt_elements.Elements_with_stock[i]
	}

	db := models.MongoCN.Database("restoner_cartadiaria")
	col := db.Collection("elements")

	// transaction
	err_transaction := db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		_, err_delete := col.DeleteMany(ctx, bson.M{"idcarta": input_mqtt_elements.IdCarta})
		if err_delete != nil {
			return err_delete
		}

		_, err_insermany := col.InsertMany(ctx, array_elements)
		if err_insermany != nil {
			return err_insermany
		}

		if err = sessionContext.CommitTransaction(sessionContext); err != nil {
			return err
		}
		return nil
	})

	if err_transaction != nil {
		return err_transaction
	}

	return nil
}
