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

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT_Anfitrion(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT_Anfitrion
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------EXTERNAL DATA----------------------*/

func (wcr *webCartaDiariaRouter_pg) Web_GetBusinessInformation(c echo.Context) error {

	//Recibimos el id del Business Owner
	uniquename := c.Param("uniquename")
	//idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:80/v1/web/business/comensal/bnss/" + uniquename)
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

func (wcr *webCartaDiariaRouter_pg) Web_GetBusinessPost(c echo.Context) error {

	idbusiness := c.Param("idbusiness")
	limit := c.Param("limit")

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:80/v1/web/business/comensal/bnss/post/" + idbusiness + "/" + limit)
	var get_respuesta ResponsePost
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

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_GetBusinessElement_Service(date, idbusiness_int, limit_int)
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

func (cr *webCartaDiariaRouter_pg) Web_SearchByNameAndDescription(c echo.Context) error {

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_SearchByNameAndDescription_Service(date, idbusiness_int, name, limit_int)
	results := ResponseCartaElements_Searched{Error: boolerror, DataError: dataerror, Data: data}
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

/*==================================================================ANFITRION================================================================*/

/*----------------------GET DATA OF MENU----------------------*/

func (cr *webCartaDiariaRouter_pg) Web_Anfitrion_GetBusinessCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_Anfitrion_GetBusinessCategory_Service(date, data_idbusiness)
	results := ResponseCartaCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *webCartaDiariaRouter_pg) Web_Anfitrion_GetBusinessElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_Anfitrion_GetBusinessElement_Service(date, data_idbusiness, limit_int)
	results := ResponseCartaElements{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *webCartaDiariaRouter_pg) Web_Anfitrion_GetBusinessSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_Anfitrion_GetBusinessSchedule_Service(date, data_idbusiness)
	results := ResponseCartaSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *webCartaDiariaRouter_pg) Web_Anfitrion_SearchByNameAndDescription(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idbusiness <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)

	//Recibimos el nombre
	name := c.Request().URL.Query().Get("name")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := Web_Anfitrion_SearchByNameAndDescription_Service(date, data_idbusiness, name, limit_int)
	results := ResponseCartaElements_Searched{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
