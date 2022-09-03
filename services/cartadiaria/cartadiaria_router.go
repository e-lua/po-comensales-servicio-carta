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

func GetJWT_Comensal(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://c-registro-authenticacion.restoner-api.fun:80/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

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

func (cr *cartaDiariaRouter_pg) GetBusinessInformation(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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
	respuesta, _ := http.Get("http://a-informacion.restoner-api.fun:80/v1/business/comensal/bnss/" + idbusiness)
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
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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

func (cr *cartaDiariaRouter_pg) GetBusinessSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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
	status, boolerror, dataerror, data_idcomensal := GetJWT_Comensal(c.Request().Header.Get("Authorization"))
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

/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*-------------------------------------CARTA DIARIA ANFITRIONES-------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/

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

	/*if offset_int == 0 {
		offset_int = 1
	}*/

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := SearchByName_Anfitrion_Service(date, data_idbusiness, text, limit_int, offset_int)
	results := ResponseCartaElements_Searched_Pg{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) AddCarta(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idbusiness := GetJWT_Anfitrion(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := ResponseInt{Error: boolerror, DataError: "000" + dataerror, Data: 0}
		return c.JSON(status, results)

	}
	if data_idbusiness <= 0 {
		results := ResponseInt{Error: boolerror, DataError: "000" + "Token incorrecto", Data: 0}
		return c.JSON(400, results)
	}

	//Instanciamos una variable del modelo Carta
	var carta Carta

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta)
	if err != nil {
		results := ResponseInt{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio, detalles: " + err.Error(), Data: 0}
		return c.JSON(403, results)
	}

	if !carta.WannaCopy {
		//Enviamos los datos al servicio
		status, boolerror, dataerror, data := AddCarta_Service(carta, data_idbusiness)
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	} else {
		//Enviamos los datos al servicio
		status, boolerror, dataerror, data := AddCartaFromOther_Service(carta, data_idbusiness)
		results := ResponseInt{Error: boolerror, DataError: dataerror, Data: data}
		return c.JSON(status, results)
	}

}

func (cdr *cartaDiariaRouter_pg) UpdateCartaStatus(c echo.Context) error {

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

	//Instanciamos una variable del modelo Category
	var carta_status CartaStatus

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_status)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaStatus_Service(carta_status, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaElements(c echo.Context) error {

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

	//Instanciamos una variable del modelo CartaElements
	var carta_elements CartaElements_WithAction

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_elements)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaElements_Service(carta_elements, data_idbusiness, 0, 0)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaOneElement(c echo.Context) error {

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

	//Recibimos el JWT
	stock_string := c.Request().URL.Query().Get("stock")
	idelement_string := c.Request().URL.Query().Get("idelement")
	idcarta_string := c.Request().URL.Query().Get("idcarta")

	//Convertimos a int
	stock, _ := strconv.Atoi(stock_string)
	idcarta, _ := strconv.Atoi(idelement_string)
	idelement, _ := strconv.Atoi(idcarta_string)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaOneElement_Service(stock, idelement, idcarta, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) UpdateCartaScheduleRanges(c echo.Context) error {

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

	//Instanciamos una variable del modelo CartaSchedule
	var carta_schedule CartaSchedule

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_schedule)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := UpdateCartaScheduleRanges_Service(carta_schedule, data_idbusiness)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cdr *cartaDiariaRouter_pg) GetCartaBasicData(c echo.Context) error {

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

	//Recibimos el limit
	date := c.Param("date")

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaBasicData_Service(date, data_idbusiness)
	results := ResponseCartaBasicData{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaCategory(c echo.Context) error {

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

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaCategory_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaCategory_Ext{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaElementsByCarta(c echo.Context) error {

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

	//Recibimos el idcarta
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Recibimos el idcategory
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaElementsByCarta_Service(idcarta_int, data_idbusiness, idcategory_int)
	results := ResponseCartaElements_Ext{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaElements(c echo.Context) error {

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

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaElements_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaElements_Ext{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartaScheduleRanges(c echo.Context) error {

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

	//Recibimos el limit
	idcarta := c.Param("idcarta")
	idcarta_int, _ := strconv.Atoi(idcarta)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartaScheduleRanges_Service(idcarta_int, data_idbusiness)
	results := ResponseCartaSchedule_Ext{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetCartas(c echo.Context) error {

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

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetCartas_Service(data_idbusiness)
	results := ResponseCartas{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cdr *cartaDiariaRouter_pg) DeleteCarta(c echo.Context) error {

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

	//Instanciamos una variable del modelo Category
	var carta_status CartaStatus

	//Agregamos los valores enviados a la variable creada
	err := c.Bind(&carta_status)
	if err != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := DeleteCarta_Service(data_idbusiness, carta_status.IDCarta)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

/*----------------------OBTENER TODOS LOS DATOS NEGOCIOS PARA NOTIFICARLOS----------------------*/

func (cdr *cartaDiariaRouter_pg) SearchToNotifyCarta() {

	//Enviamos los datos al servicio
	status, _, dataerror, _ := SearchToNotifyCarta_Service()
	log.Println(strconv.Itoa(status) + " " + dataerror)
}

/*----------------------DELETE----------------------*/

func (cdr *cartaDiariaRouter_pg) Delete_Vencidas() {

	error_delete, data := Delete_Vencidas_Service()
	log.Println(error_delete, data)

}

/*-------------------------------------CARTA DIARIA ANFITRIONES <==> CREAR ORDEN-------------------------------------*/

func (cdr *cartaDiariaRouter_pg) GetCategories_ToCreateOrder(c echo.Context) error {

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
	status, boolerror, dataerror, data := GetCategories_ToCreateOrder_Service(date, data_idbusiness)
	results := ResponseCartaCategory_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetElements_ToCreateOrder(c echo.Context) error {

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

	//Recibimos el id de la categoria
	idcategory := c.Param("idcategory")
	idcategory_int, _ := strconv.Atoi(idcategory)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := GetElements_ToCreateOrder_Service(date, data_idbusiness, idcategory_int)
	results := ResponseCartaElements_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) GetSchedule_ToCreateOrder(c echo.Context) error {

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
	status, boolerror, dataerror, data := GetSchedule_ToCreateOrder_Service(date, data_idbusiness)
	results := ResponseCartaSchedule_ToCreate{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}

func (cdr *cartaDiariaRouter_pg) Find__Notify_NoCarta() {

	//Enviamos los datos al servicio
	error_notify := Find__Notify_NoCarta_Service()

	if error_notify != nil {
		log.Println(error_notify)

	}

	log.Println("Notificado con exito")
}
