package carta_web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var Web_CartaDiariaRouter_pg *webCartaDiariaRouter_pg

type webCartaDiariaRouter_pg struct {
}

/*----------------------EXTERNAL DATA----------------------*/

func (wcr *webCartaDiariaRouter_pg) Web_GetBusinessInformation(c echo.Context) error {

	//Recibimos el id del Business Owner
	uniquename := c.Param("uniquename")
	//idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:80/v1/web/business/comensal/bnss/" + uniquename)
	var get_respuesta ResponseBusiness
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Agregamos la vista del comensal
	//GetBusinessInformation_Service(data_idcomensal, idbusiness_int)

	return c.JSON(200, get_respuesta)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cr *webCartaDiariaRouter_pg) Web_GetBusinessCategory(c echo.Context) error {

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)
	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_GetBusinessCategory_Service(date, idbusiness_int)
	results := ResponseCartaCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *webCartaDiariaRouter_pg) Web_GetBusinessElement(c echo.Context) error {

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_GetBusinessElement_Service(date, idbusiness_int, idcategory_int)
	results := ResponseCartaElements{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *webCartaDiariaRouter_pg) Web_GetBusinessSchedule(c echo.Context) error {

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_GetBusinessSchedule_Service(date, idbusiness_int)
	results := ResponseCartaSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*===========================*/
/*=========VERSION 2=========*/
/*===========================*/

func (cr *webCartaDiariaRouter_pg) Web_GetBusinessInformation_V2(c echo.Context) error {

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")
	//idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:80/v1/business/comensal/bnss/" + idbusiness)
	var get_respuesta ResponseBusiness_V2
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Agregamos la vista del comensal
	//GetBusinessInformation_Service(data_idcomensal, idbusiness_int)

	return c.JSON(200, get_respuesta)
}
