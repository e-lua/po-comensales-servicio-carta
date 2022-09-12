package cartadiaria_anfitrion

import (
	"context"
	"strconv"
	"strings"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Mo_Delete_Update_Schedule(pg_schedule []models.Pg_ScheduleRange_External, idcarta int, idbusiness int) error {

	//Variables para el MQTT
	var lista_total_schedule []interface{}

	//Obtener una fecha
	one_date := ""
	counter_to_skip := 0

	for _, sch := range pg_schedule {

		var list_schedule models.Mo_ListoSchedule_With_Stock

		arr := strings.Split(sch.StartTime, ":")
		hora_ini := 0
		hora_ini_string := sch.StartTime

		for i := 0; i < sch.NumberOfFractions; i++ {

			if i == 0 {
				//TODO SOBRE LA HORA DE INICIO
				hora_ini_c, _ := strconv.Atoi(arr[0] + arr[1][:2])
				hora_ini = hora_ini_c
			}

			//TODO SOBRE LA HORA PRE FIN
			hora_pre_fin := strconv.Itoa(hora_ini + sch.MinutePerFraction)

			//Definimos los minutos y horas
			var index_pre_fin, minutos, horas int
			switch len(hora_pre_fin) {
			case 3:
				index_pre_fin = 1
				//Minutos y Horas
				minutos, _ = strconv.Atoi(hora_pre_fin[index_pre_fin:])
				horas, _ = strconv.Atoi(hora_pre_fin[:index_pre_fin])
			case 4:
				index_pre_fin = 2
				//Minutos y Horas
				minutos, _ = strconv.Atoi(hora_pre_fin[index_pre_fin:])
				horas, _ = strconv.Atoi(hora_pre_fin[:index_pre_fin])
			default:
				index_pre_fin = 0
				//Minutos y Horas
				minutos, _ = strconv.Atoi(hora_pre_fin[index_pre_fin:])
				horas = 0
			}

			//Validamos que no sobrepase los 60 minutos
			var minutos_string string
			if minutos > 59 {
				minutos = 60 - minutos
				if minutos < 10 && minutos >= 0 {
					minutos_string = "0" + strconv.Itoa(minutos)
				} else {
					if minutos < 0 {
						minutos_string = strconv.Itoa(minutos * -1)
					} else {
						minutos_string = strconv.Itoa(minutos)
					}
				}
				horas = horas + 1
			} else {
				minutos_string = hora_pre_fin[index_pre_fin:]
			}

			//Validamos que no sobrepase las 24 horas
			var horas_string string
			if horas > 23 {
				horas = 24
				horas_string = "24"
			} else {
				if horas == 0 {
					horas_string = "00"
				} else {
					horas_string = strconv.Itoa(horas)
				}
			}

			//Hora que finaliza
			hora_finaliza := horas_string + minutos_string

			//TODO SOBRE LA HORA FIN
			var index_fin int
			if len(hora_finaliza) > 3 {
				index_fin = 2
			} else {
				index_fin = 1
			}

			//Le pondremos un 0 al comienzo de un numero si es necesario, para horas como las 3 de la mañana= 03
			var hora_fin_toinsert string
			if len(hora_finaliza[:index_fin]) == 1 {
				hora_fin_toinsert = "0" + strconv.Itoa(horas) + ":" + hora_finaliza[index_fin:]
			} else {
				hora_fin_toinsert = strconv.Itoa(horas) + ":" + hora_finaliza[index_fin:]
			}

			//Insertamos los datos en el modelo
			list_schedule.IDSchedule = sch.IDSchedule
			list_schedule.IDCarta = idcarta
			list_schedule.Date = sch.Date
			list_schedule.IDBusiness = idbusiness
			list_schedule.Starttime = hora_ini_string
			list_schedule.Endtime = hora_fin_toinsert
			list_schedule.MaxOrders = sch.MaxOrders
			list_schedule.Timezone = sch.TimeZone

			if counter_to_skip == 0 {
				one_date = sch.Date
			}
			counter_to_skip = counter_to_skip + 1

			//Nuevo valor de hora de inicio
			new_hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + hora_finaliza[index_fin:])
			hora_ini = new_hora_ini
			hora_ini_string = hora_fin_toinsert

			//Si supera la media noche se termina el bucle
			if horas_string == "24" {
				break
			}

		}

	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_cartadiaria")
	col := db.Collection("schedule_list")

	// transaction
	err_transaction := db.Client().UseSession(ctx, func(sessionContext mongo.SessionContext) error {
		err := sessionContext.StartTransaction()
		if err != nil {
			return err
		}

		_, err_delete := col.DeleteMany(ctx, bson.M{"date": one_date})
		if err_delete != nil {
			return err_delete
		}

		_, err_insermany := col.InsertMany(ctx, lista_total_schedule)
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
