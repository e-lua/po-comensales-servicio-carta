package carta

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/labstack/echo/v4"
)

var CartaDiariaRouter_pg *cartaDiariaRouter_pg

type cartaDiariaRouter_pg struct {
}

/*----------------------UDPATE DATA CONSUME----------------------*/

func (cr *cartaDiariaRouter_pg) UpdateElementStock(element_stock models.Pg_ToElement_Mqtt) {

	//Enviamos los datos al servicio
	error_element_stock := UpdateElementStock_Service(element_stock)
	if error_element_stock != nil {
		log.Fatal(error_element_stock)
	}
}

func (cr *cartaDiariaRouter_pg) UpdateScheduleStock(schedule_stock models.Pg_ToSchedule_Mqtt) {

	//Enviamos los datos al servicio
	error_schedule_stock := UpdateScheduleStock_Service(schedule_stock)
	if error_schedule_stock != nil {
		log.Fatal(error_schedule_stock)
	}
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

func GetJWT_Anfitrion(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://a-registro-authenticacion.restoner-api.fun:5000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT_Anfitrion
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IdBusiness
}

/*----------------------EXTERNAL DATA----------------------*/

func (cr *cartaDiariaRouter_pg) GetBusinessInformation(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")
	//idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:5800/v1/business/comensal/bnss/" + idbusiness)
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

func (cr *cartaDiariaRouter_pg) GetBusinessCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)
	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessCategory_Service(date, idbusiness_int)
	results := ResponseCartaCategory{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaDiariaRouter_pg) GetBusinessElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessElement_Service(date, idbusiness_int, idcategory_int)
	results := ResponseCartaElements{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaDiariaRouter_pg) SearchByNameAndDescription(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Recibimos el text
	text := c.Param("text")

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Recibimos el limit
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	/*if offset_int == 0 {
		offset_int = 1
	}*/

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchByNameAndDescription_Service(date, idbusiness_int, text, limit_int, offset_int)
	results := ResponseCartaElements_Searched{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaDiariaRouter_pg) SearchByName_Anfitrion(c echo.Context) error {

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

	//Recibimos el text
	text := c.Param("text")

	//Recibimos el limit
	limit := c.Param("limit")
	limit_int, _ := strconv.Atoi(limit)
	//Recibimos el limit
	offset := c.Param("offset")
	offset_int, _ := strconv.Atoi(offset)

	if offset_int == 0 {
		offset_int = 1
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchByName_Anfitrion_Service(date, data_idbusiness, text, limit_int, offset_int)
	results := ResponseCartaElements_Searched_Mo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cr *cartaDiariaRouter_pg) GetBusinessSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos la fecha de la carta
	date := c.Param("date")

	//Recibimos el id del negocio
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetBusinessSchedule_Service(date, idbusiness_int)
	results := ResponseCartaSchedule{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------ADD VIEW DATA ELEMENT----------------------*/

func (cr *cartaDiariaRouter_pg) AddViewElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del negocio
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Validamos los valores enviados
	if idelement_int < 1 {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio"}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddViewInformation_Service(idelement_int, data_idcomensal)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*===========================*/
/*=========VERSION 2=========*/
/*===========================*/

func (cr *cartaDiariaRouter_pg) GetBusinessInformation_V2(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: "000" + dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "000" + "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")
	//idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:5800/v1/business/comensal/bnss/" + idbusiness)
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

/*-------------------------------------ELEMENTS-------------------------------------*/

func (cr *cartaDiariaRouter_pg) UpdateCarta_ElementsWithStock(inputserialize_elements models.Mqtt_Element_With_Stock_Import) {
	//Enviamos los datos al servicio
	error_r := UpdateCarta_ElementsWithStock_Service(inputserialize_elements)
	if error_r != nil {
		log.Fatal(error_r)
	}
}

func (cr *cartaDiariaRouter_pg) GetElementsByInsumo(c echo.Context) error {

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

	//Recibimos los parametros de filtro
	date := c.Param("date")
	idinsumo := c.Param("idinsumo")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetElementsByInsumo_Service(date, data_idbusiness, idinsumo)
	results := ResponseCartaElements_Searched_Mo{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
