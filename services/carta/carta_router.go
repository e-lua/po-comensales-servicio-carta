package carta

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Aphofisis/po-comensales-servicio-carta/models"
	"github.com/labstack/echo/v4"
)

var CartaRouter_pg *cartaRouter_pg

type cartaRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

func GetJWT(jwt string) (int, bool, string, int) {
	//Obtenemos los datos del auth
	respuesta, _ := http.Get("http://143.110.145.136:3000/v1/trylogin?jwt=" + jwt)
	var get_respuesta ResponseJWT
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		return 500, true, "Error en el sevidor interno al intentar decodificar la autenticacion, detalles: " + error_decode_respuesta.Error(), 0
	}
	return 200, false, "", get_respuesta.Data.IDComensal
}

/*----------------------EXTERNAL DATA----------------------*/

func (cr *cartaRouter_pg) GetBusinessInformation(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: ""}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: true, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del Business Owner
	idbusiness := c.Param("idbusiness")
	idbusiness_int, _ := strconv.Atoi(idbusiness)

	//Enviamos los datos al servicio de anfitriones para obtener los datos completos
	respuesta, _ := http.Get("http://137.184.74.10:5800/v1/business/comensal/bnss/" + idbusiness)
	var get_respuesta models.Mo_Business
	error_decode_respuesta := json.NewDecoder(respuesta.Body).Decode(&get_respuesta)
	if error_decode_respuesta != nil {
		results := Response{Error: true, DataError: "El valor ingresado no cumple con la regla de negocio", Data: ""}
		return c.JSON(403, results)
	}

	//Agregamos la vista del comensal
	GetBusinessInformation_Service(data_idcomensal, idbusiness_int)

	//Enviamos el resultado de la consulta del endpoint
	results := ResponseBusiness{Error: boolerror, DataError: dataerror, Data: get_respuesta}

	return c.JSON(status, results)
}

/*----------------------GET DATA OF MENU----------------------*/

func (cr *cartaRouter_pg) GetBusinessCategory(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetBusinessElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) GetBusinessSchedule(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
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

func (cr *cartaRouter_pg) AddViewElement(c echo.Context) error {

	//Obtenemos los datos del auth
	status, boolerror, dataerror, data_idcomensal := GetJWT(c.Request().Header.Get("Authorization"))
	if dataerror != "" {
		results := Response{Error: boolerror, DataError: dataerror, Data: dataerror}
		return c.JSON(status, results)
	}
	if data_idcomensal <= 0 {
		results := Response{Error: boolerror, DataError: "Token incorrecto", Data: ""}
		return c.JSON(400, results)
	}

	//Recibimos el id del negocio
	idelement := c.Param("idelement")
	idelement_int, _ := strconv.Atoi(idelement)

	//Enviamos los datos al servicio
	status, boolerror, dataerror, data := AddViewElement_Service(data_idcomensal, idelement_int)
	results := Response{Error: boolerror, DataError: dataerror, Data: data}
	return c.JSON(status, results)
}
