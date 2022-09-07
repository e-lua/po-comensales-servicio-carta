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

func Mo_Delete_Update_Schedule(pg_schedule []models.Pg_ScheduleRange_External, idbusiness int) error {

	//Variables para el MQTT
	var lista_total_schedule []interface{}

	//Obtener una fecha
	one_date := ""
	counter_to_skip := 0

	for _, sch := range pg_schedule {

		var list_schedule models.Mo_ListoSchedule_With_Stock

		if "-5" != "" {
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

				var index_pre_fin int
				if len(hora_pre_fin) > 3 {
					index_pre_fin = 2
				} else {
					index_pre_fin = 1
				}

				//Minutos y Horas
				minutos, _ := strconv.Atoi(hora_pre_fin[index_pre_fin:])
				horas, _ := strconv.Atoi(hora_pre_fin[:index_pre_fin])

				//Validamos que no sobrepase los 60 minutos
				var minutos_string string
				if minutos > 59 {
					minutos = 60 - minutos
					if minutos < 10 {
						minutos_string = "0" + strconv.Itoa(minutos)
					} else {
						minutos_string = strconv.Itoa(minutos)
					}
					horas = horas + 1
				} else {
					minutos_string = hora_pre_fin[index_pre_fin:]
				}
				hora_finaliza := strconv.Itoa(horas) + minutos_string

				//TODO SOBRE LA HORA FIN
				var index_fin int
				if len(hora_finaliza) > 3 {
					index_fin = 2
				} else {
					index_fin = 1
				}
				hora_fin_toinsert := hora_finaliza[:index_fin] + ":" + hora_finaliza[index_fin:]

				//Fin de bucle para obtener la hora fin

				//Insertamos los datos en el modelo
				list_schedule.IDSchedule = sch.IDSchedule
				list_schedule.IDCarta = sch.IdCarta
				list_schedule.Date = sch.Date
				list_schedule.IDBusiness = idbusiness
				list_schedule.Starttime = hora_ini_string
				list_schedule.Endtime = hora_fin_toinsert
				list_schedule.MaxOrders = sch.MaxOrders
				list_schedule.Timezone = "-5"

				if counter_to_skip == 0 {
					one_date = sch.Date
				}

				counter_to_skip = counter_to_skip + 1

				lista_total_schedule = append(lista_total_schedule, list_schedule)

				//Nuevo valor de hora de inicio
				new_hora_ini, _ := strconv.Atoi(strconv.Itoa(horas) + minutos_string)
				hora_ini = new_hora_ini
				hora_ini_string = hora_fin_toinsert

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
