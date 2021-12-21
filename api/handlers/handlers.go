package api

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/cors"

	carta "github.com/Aphofisis/po-comensales-servicio-carta/services/carta"
)

func Manejadores() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", index)
	//VERSION
	version_1 := e.Group("/v1")

	/*===========CARTA===========*/
	//V1 FROM V1 TO ...TO ENTITY MENU
	router_business := version_1.Group("/business/data")
	router_business.GET("/:idbusiness", carta.CartaRouter_pg.GetBusinessInformation)
	router_business.GET("/:idbusiness/menu/:date/category", carta.CartaRouter_pg.GetBusinessCategory)
	router_business.GET("/:idbusiness/menu/:date/category/:idcategory/elements", carta.CartaRouter_pg.GetBusinessElement)
	router_business.GET("/:idbusiness/menu/:date/scheduleranges", carta.CartaRouter_pg.GetBusinessSchedule)

	//V1 FROM V1 TO ...TO VIEW
	router_view := version_1.Group("/view")
	router_view.POST("/:idelement", carta.CartaRouter_pg.AddViewElement)

	/*===========CARTA===========*/

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
