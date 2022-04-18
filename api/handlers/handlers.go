package api

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria "github.com/Aphofisis/po-comensales-servicio-carta/services/cartadiaria"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Consumidor-MQTT
	go Consumer_Element_Stock()
	go Consumer_Schedule_Stock()
	go Consumer_Elements_With_Stock()

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_business := version_1.Group("/business/data")
	router_business.GET("/:idbusiness/information", cartadiaria.CartaDiariaRouter_pg.GetBusinessInformation)
	router_business.GET("/:idbusiness/menu/:date/category", cartadiaria.CartaDiariaRouter_pg.GetBusinessCategory)
	router_business.GET("/:idbusiness/menu/:date/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetBusinessElement)
	router_business.GET("/:idbusiness/menu/:date/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetBusinessSchedule)
	router_business.GET("/:idbusiness/menu/:date/search/:text/:limit/:offset", cartadiaria.CartaDiariaRouter_pg.SearchByNameAndDescription)

	//V1 FROM V1 TO ...TO VIEW
	router_view := version_1.Group("/view")
	router_view.POST("/:idelement", cartadiaria.CartaDiariaRouter_pg.AddViewElement)

	/*===========================*/
	/*=========VERSION 2=========*/
	/*===========================*/

	version_2 := e.Group("/v2")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_business2 := version_2.Group("/business/data")
	router_business2.GET("/:idbusiness/information", cartadiaria.CartaDiariaRouter_pg.GetBusinessInformation_V2)

	//Abrimos el puerto
	PORT := os.Getenv("PORT")
	//Si dice que existe PORT
	if PORT == "" {
		PORT = "6500"
	}

	//cors son los permisos que se le da a la API
	//para que sea accesibl esde cualquier lugar
	handler := cors.AllowAll().Handler(e)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

func index(c echo.Context) error {
	return c.JSON(401, "Acceso no autorizado")
}

func Consumer_Element_Stock() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/stock", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop := make(chan bool)

	go func() {
		for d := range msgs {
			var element_stock models.Pg_ToElement_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&element_stock)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			cartadiaria.CartaDiariaRouter_pg.UpdateElementStock(element_stock)
		}
	}()

	<-noStop

}

func Consumer_Schedule_Stock() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/schedulestock", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop2 := make(chan bool)

	go func() {
		for d := range msgs {
			var schedule_stock models.Pg_ToSchedule_Mqtt
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&schedule_stock)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			cartadiaria.CartaDiariaRouter_pg.UpdateScheduleStock(schedule_stock)
		}
	}()

	<-noStop2

}

func Consumer_Elements_With_Stock() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/cartadiaria_elements", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStopElementsWithStock := make(chan bool)

	go func() {
		for d := range msgs {
			var elements_stock models.Mqtt_Element_With_Stock_Import
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&elements_stock)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			cartadiaria.CartaDiariaRouter_pg.UpdateCarta_ElementsWithStock(elements_stock)
		}
	}()

	<-noStopElementsWithStock
}
