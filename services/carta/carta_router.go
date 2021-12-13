package carta

var CartaRouter_pg *cartaRouter_pg

type cartaRouter_pg struct {
}

/*----------------------TRAEMOS LOS DATOS DEL AUTENTICADOR----------------------*/

/*
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
*/

/*----------------------CONSUMER----------------------*/
