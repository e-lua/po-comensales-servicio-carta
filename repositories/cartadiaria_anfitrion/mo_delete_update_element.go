package cartadiaria_anfitrion

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Mo_Delete_Update_Elements(pg_element_withaction_external []models.Pg_Element_With_Stock_External, idcarta int, idbusiness int) error {

	//Variables para el MQTT
	var elements_mqtt []interface{}

	//Repartiendo los datos
	for _, e := range pg_element_withaction_external {

		//Variables MQTT
		var one_element_mqtt models.Mqtt_Element_With_Stock
		one_element_mqtt.AvailableOrders = e.AvailableOrders
		one_element_mqtt.Date = e.Date
		one_element_mqtt.DeletedDate = time.Now().AddDate(0, 0, 5)
		one_element_mqtt.Description = e.Description
		one_element_mqtt.IDBusiness = idbusiness
		one_element_mqtt.IDCarta = idcarta
		one_element_mqtt.IDCategory = e.IDCategory
		one_element_mqtt.IDElement = e.IDElement
		one_element_mqtt.IsExported = false
		one_element_mqtt.Name = e.Name
		one_element_mqtt.NameCategory = e.NameCategory
		one_element_mqtt.Price = e.Price
		one_element_mqtt.Stock = e.Stock
		one_element_mqtt.TypeMoney = e.TypeMoney
		one_element_mqtt.Typefood = e.Typefood
		one_element_mqtt.UrlPhoto = e.UrlPhoto
		one_element_mqtt.UrlPhotoCategory = e.UrlPhotoCategory
		one_element_mqtt.Insumos = e.Insumos
		one_element_mqtt.Costo = e.Costo
		elements_mqtt = append(elements_mqtt, one_element_mqtt)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_cartadiaria")
	col := db.Collection("elements")

	// transaction
	err_transaction := db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		_, err_delete := col.DeleteMany(ctx, bson.M{"idcarta": idcarta})
		if err_delete != nil {
			return err_delete
		}

		_, err_insermany := col.InsertMany(ctx, elements_mqtt)
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
