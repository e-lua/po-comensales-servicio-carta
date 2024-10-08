package api

import (
	"bytes"
	"encoding/json"
	"log"

	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/rs/cors"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
	cartadiaria "github.com/Aphofisis/po-comensales-servicio-carta/services/cartadiaria"
	cartadiaria_web "github.com/Aphofisis/po-comensales-servicio-carta/services/cartadiaria_web"
	imports "github.com/Aphofisis/po-comensales-servicio-carta/services/imports"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	/*e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAccessControlAllowOrigin},
	}))*/

	//Consumidor-MQTT
	//go Consumer_Element_Stock()
	go Consumer_Schedule_Stock()
	go Delete_Vencidas()
	//go Notify_NoCarta()

	e.GET("/", index)

	/*---------------------------------------------------------------------------------------------------------------------------------*/

	//VERSION WEB
	version_1_web := e.Group("/v1/web")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_business_web := version_1_web.Group("/business/data")
	router_business_web.GET("/:uniquename/information", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessInformation)
	router_business_web.GET("/:idbusiness/post/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessPost)
	router_business_web.GET("/:idbusiness/menu/:date/categories", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessCategory)
	router_business_web.GET("/:idbusiness/menu/:date/elements/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessElement)
	router_business_web.GET("/:idbusiness/menu/:date/scheduleranges", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessSchedule)
	router_business_web.GET("/:idbusiness/menu/:date/search/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_SearchByNameAndDescription)

	/*===========CARTA ANFITRION===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_anfitrion_web := version_1_web.Group("/anfitrion/menu/createorder")
	router_anfitrion_web.GET("/:date/category", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_Anfitrion_GetBusinessCategory)
	router_anfitrion_web.GET("/:date/elements/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_Anfitrion_GetBusinessElement)
	router_anfitrion_web.GET("/:date/scheduleranges", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_Anfitrion_GetBusinessSchedule)
	router_anfitrion_web.GET("/:date/search/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_Anfitrion_SearchByNameAndDescription)

	/*---------------------------------------------------------------------------------------------------------------------------------*/

	//VERSION
	version_1 := e.Group("/v1")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_business := version_1.Group("/business/data")
	router_business.GET("/:idbusiness/information", cartadiaria.CartaDiariaRouter_pg.GetBusinessInformation)
	router_business.GET("/:idbusiness/post/:limit", cartadiaria_web.Web_CartaDiariaRouter_pg.Web_GetBusinessPost)
	router_business.GET("/:idbusiness/menu/:date/category", cartadiaria.CartaDiariaRouter_pg.GetBusinessCategory)
	router_business.GET("/:idbusiness/menu/:date/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetBusinessElement)
	router_business.GET("/:idbusiness/menu/:date/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetBusinessSchedule)
	router_business.GET("/:idbusiness/menu/:date/search/:text/:limit/:offset", cartadiaria.CartaDiariaRouter_pg.SearchByNameAndDescription)
	//Ver la lista de elementos por categorias
	router_business.GET("/:idbusiness/menu/:date/elements/:limit", cartadiaria.CartaDiariaRouter_pg.GetBusinessElement_ListByCategory)

	/*to create an order - ANFITRION*/
	router_anfitrion_menu := version_1.Group("/anfitrion/menu")
	router_anfitrion_menu.POST("", cartadiaria.CartaDiariaRouter_pg.AddCarta)
	router_anfitrion_menu.PUT("", cartadiaria.CartaDiariaRouter_pg.UpdateCartaStatus)
	router_anfitrion_menu.GET("", cartadiaria.CartaDiariaRouter_pg.GetCartas)
	router_anfitrion_menu.DELETE("", cartadiaria.CartaDiariaRouter_pg.DeleteCarta)
	router_anfitrion_menu.GET("/:date", cartadiaria.CartaDiariaRouter_pg.GetCartaBasicData)
	router_anfitrion_menu.GET("/:idcarta/category", cartadiaria.CartaDiariaRouter_pg.GetCartaCategory)
	router_anfitrion_menu.GET("/:idcarta/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetCartaElementsByCarta)
	router_anfitrion_menu.PUT("/elements", cartadiaria.CartaDiariaRouter_pg.UpdateCartaElements)
	router_anfitrion_menu.GET("/:idcarta/elements", cartadiaria.CartaDiariaRouter_pg.GetCartaElements)
	router_anfitrion_menu.PUT("/onelement", cartadiaria.CartaDiariaRouter_pg.UpdateCartaOneElement)
	router_anfitrion_menu.PUT("/scheduleranges", cartadiaria.CartaDiariaRouter_pg.UpdateCartaScheduleRanges)
	router_anfitrion_menu.GET("/:idcarta/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetCartaScheduleRanges)
	router_anfitrion_menu.GET("/carta/:date/insumo/:idinsumo/elements", cartadiaria.CartaDiariaRouter_pg.GetElementsByInsumo)
	router_anfitrion_menu.PUT("/automaticdiscounts", cartadiaria.CartaDiariaRouter_pg.UpdateCartaAutomaticDiscounts)
	router_anfitrion_menu.GET("/:idcarta/automaticdiscounts", cartadiaria.CartaDiariaRouter_pg.GetCartAutomaticDiscounts)

	/*Create order*/
	router_anfitrion_menu_createorder := version_1.Group("/anfitrion/menu/createorder")
	router_anfitrion_menu_createorder.GET("/:date/category", cartadiaria.CartaDiariaRouter_pg.GetCategories_ToCreateOrder)
	router_anfitrion_menu_createorder.GET("/:date/category/:idcategory/elements", cartadiaria.CartaDiariaRouter_pg.GetElements_ToCreateOrder)
	router_anfitrion_menu_createorder.GET("/:date/scheduleranges", cartadiaria.CartaDiariaRouter_pg.GetSchedule_ToCreateOrder)
	router_anfitrion_menu_createorder.GET("/:date/search/:text/:limit/:offset", cartadiaria.CartaDiariaRouter_pg.SearchByName_Anfitrion)
	//router_anfitrion_menu_createorder.GET("/:date/automaticdiscounts", cartadiaria.CartaDiariaRouter_pg.GetAutomaticDiscount_ToCreateOrder)

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

	//e.Start(":6500")
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

	msgs, err_consume := ch.Consume("anfitrion/stock_element_add", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop := make(chan bool)

	go func() {
		for d := range msgs {
			var element_stock []models.Mqtt_Import_ElementStock
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&element_stock)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			imports.ImportsRouter_pg.UpdateElementStock(element_stock)
		}
	}()

	<-noStop

}

func Consumer_Schedule_Stock() {

	ch, error_conection := models.MqttCN.Channel()
	if error_conection != nil {
		log.Fatal("Error connection canal " + error_conection.Error())
	}

	msgs, err_consume := ch.Consume("anfitrion/stock_schedule_add", "", true, false, false, false, nil)
	if err_consume != nil {
		log.Fatal("Error connection cola " + err_consume.Error())
	}

	noStop2 := make(chan bool)

	go func() {
		for d := range msgs {
			var schedule_stock []models.Mqtt_Import_SheduleStock
			buf := bytes.NewBuffer(d.Body)
			decoder := json.NewDecoder(buf)
			err_consume := decoder.Decode(&schedule_stock)
			if err_consume != nil {
				log.Fatal("Error decoding")
			}
			imports.ImportsRouter_pg.UpdateScheduleStock(schedule_stock)
		}
	}()

	<-noStop2

}

func Delete_Vencidas() {

	noStopDeleteV := make(chan bool)

	go func() {
		for {
			time.Sleep(12 * time.Hour)
			cartadiaria.CartaDiariaRouter_pg.Delete_Vencidas()
		}
	}()

	<-noStopDeleteV
}

func Notify_NoCarta() {

	noStopNotify_NoCarta := make(chan bool)

	go func() {
		for {
			time.Sleep(1 * time.Hour)
			cartadiaria.CartaDiariaRouter_pg.Find__Notify_NoCarta()
		}
	}()

	<-noStopNotify_NoCarta
}
