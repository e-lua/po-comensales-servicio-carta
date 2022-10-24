package cartadiaria_anfitrion

import (
	"context"
	"strconv"
	"strings"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Copy_Carta(pg_automaticdiscount []models.Pg_V2_AutomaticDiscount, pg_schedule []models.Pg_ScheduleRange_External, pg_element_external []models.Pg_Element_With_Stock_External, idbusiness int, date string, idcarta int) (int, error) {

	//Elementos
	idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg, costo_pg, discount_pg, latitude_pg, longitude_pg, typefood_pg := []int{}, []int{}, []int{}, []string{}, []string{}, []string{}, []float32{}, []string{}, []string{}, []int{}, []int{}, []int{}, []float64{}, []float32{}, []float32{}, []float32{}, []string{}
	var insumos_pg []interface{}
	var additionals_pg []interface{}

	for _, e := range pg_element_external {
		idelement_pg = append(idelement_pg, e.IDElement)
		idcarta_pg = append(idcarta_pg, idcarta)
		idcategory_pg = append(idcategory_pg, e.IDCategory)
		namecategory_pg = append(namecategory_pg, e.NameCategory)
		urlphotocategory_pg = append(urlphotocategory_pg, e.UrlPhotoCategory)
		name_pg = append(name_pg, e.Name)
		price_pg = append(price_pg, e.Price)
		description_pg = append(description_pg, e.Description)
		urlphot_pg = append(urlphot_pg, e.UrlPhoto)
		typem_pg = append(typem_pg, e.TypeMoney)
		stock_pg = append(stock_pg, e.Stock)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
		insumos_pg = append(insumos_pg, e.Insumos)
		costo_pg = append(costo_pg, e.Costo)
		discount_pg = append(discount_pg, e.Discount)
		additionals_pg = append(additionals_pg, e.Additionals)
		latitude_pg = append(latitude_pg, e.Latitude)
		longitude_pg = append(longitude_pg, e.Longitude)
		typefood_pg = append(typefood_pg, e.Typefood)
	}

	//Rango horarios
	timezone_pg_2, idschedule_pg_2, idcartamain_pg_2, idbusinessmain_pg_2, name_pg_2, description_pg_2, minutesperfraction_pg_2, numberfractions_pg_2, start_pg_2, end_pg_2, maxorders_pg_2 := []string{}, []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

		idschedule_pg_2 = append(idschedule_pg_2, sch.IDSchedule)
		idcartamain_pg_2 = append(idcartamain_pg_2, idcarta)
		idbusinessmain_pg_2 = append(idbusinessmain_pg_2, idbusiness)
		name_pg_2 = append(name_pg_2, sch.Name)
		description_pg_2 = append(description_pg_2, sch.Description)
		minutesperfraction_pg_2 = append(minutesperfraction_pg_2, sch.MinutePerFraction)
		numberfractions_pg_2 = append(numberfractions_pg_2, sch.NumberOfFractions)
		start_pg_2 = append(start_pg_2, sch.StartTime)
		end_pg_2 = append(end_pg_2, sch.EndTime)
		maxorders_pg_2 = append(maxorders_pg_2, sch.MaxOrders)
		timezone_pg_2 = append(timezone_pg_2, "-5")

	}

	//Lista de actualizacion de rangos horarios
	timezone_pg_3, idschedulerange_pg_3, idcarta_pg_3, idbusiness_pg_3, startime_pg_3, endtime_pg_3, max_orders_3 := []string{}, []int64{}, []int{}, []int{}, []string{}, []string{}, []int{}

	for _, sch := range pg_schedule {

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
			/*var horas_string string
			if horas > 23 {
				horas = 24 - horas
				horas_string = "0" + strconv.Itoa(horas)
			} else {
				if horas == 0 {
					horas_string = "00"
				} else {
					horas_string = hora_pre_fin[:index_pre_fin]
				}
			}*/
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
			idschedulerange_pg_3 = append(idschedulerange_pg_3, sch.IDSchedule)
			idcarta_pg_3 = append(idcarta_pg_3, idcarta)
			idbusiness_pg_3 = append(idbusiness_pg_3, idbusiness)
			startime_pg_3 = append(startime_pg_3, hora_ini_string)
			endtime_pg_3 = append(endtime_pg_3, hora_fin_toinsert)
			max_orders_3 = append(max_orders_3, 20)
			timezone_pg_3 = append(timezone_pg_3, "-5")

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

	//Descuentos automaticos
	var groupid_pg_4 []interface{}
	iddiscount_pg_4, description_pg_4, discount_pg_4, typediscount_pg_4, idbusiness_pg_4, classdiscount_pg_4, idcarta_pg_4 := []int{}, []string{}, []float32{}, []int{}, []int{}, []int{}, []int{}

	for _, autodisc := range pg_automaticdiscount {

		iddiscount_pg_4 = append(iddiscount_pg_4, autodisc.IDAutomaticDiscount)
		description_pg_4 = append(description_pg_4, autodisc.Description)
		discount_pg_4 = append(discount_pg_4, autodisc.Discount)
		typediscount_pg_4 = append(typediscount_pg_4, autodisc.TypeDiscount)
		idbusiness_pg_4 = append(idbusiness_pg_4, idbusiness)
		classdiscount_pg_4 = append(classdiscount_pg_4, autodisc.ClassDiscount)
		idcarta_pg_4 = append(idcarta_pg_4, idcarta)
		groupid_pg_4 = append(groupid_pg_4, autodisc.Group)
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()

	db_external := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return 0, error_tx
	}

	//INSERTAR ELEMENTO
	q_element := `INSERT INTO element(idelement,idcarta,idcategory,namecategory,urlphotcategory,name,price,description,urlphoto,typemoney,stock,idbusiness,insumos,costo,additionals,discount,latitude,longitude,typefood) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(100)[],$5::varchar(230)[],$6::varchar(100)[],$7::decimal(8,2)[],$8::varchar(250)[],$9::varchar(230)[],$10::int[],$11::int[],$12::int[],$13::jsonb[],$14::real[],$15::jsonb[],$16::decimal(8,2)[],$17::real[],$18::real[],$19::varchar(100)[]));`
	if _, err_insert_element := tx.Exec(ctx, q_element, idelement_pg, idcarta_pg, idcategory_pg, namecategory_pg, urlphotocategory_pg, name_pg, price_pg, description_pg, urlphot_pg, typem_pg, stock_pg, idbusiness_pg, insumos_pg, costo_pg, additionals_pg, discount_pg, latitude_pg, longitude_pg, typefood_pg); err_insert_element != nil {
		tx.Rollback(ctx)
		return 0, err_insert_element
	}

	//INSERTAR RANGO HORARIO
	q_schedulerange := `INSERT INTO ScheduleRange(idScheduleRange,idbusiness,idcarta,name,description,minuteperfraction,numberfractions,startTime,endTime,maxOrders,timezone) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(12)[],$5::varchar(60)[],$6::int[],$7::int[],$8::varchar(10)[],$9::varchar(10)[],$10::int[],$11::varchar(3)[]));`
	if _, err_insert_schedulerange := tx.Exec(ctx, q_schedulerange, idschedule_pg_2, idbusinessmain_pg_2, idcartamain_pg_2, name_pg_2, description_pg_2, minutesperfraction_pg_2, numberfractions_pg_2, start_pg_2, end_pg_2, maxorders_pg_2, timezone_pg_2); err_insert_schedulerange != nil {
		tx.Rollback(ctx)
		return 0, err_insert_schedulerange
	}

	//INSERTAR LISTAS DE RANGOS HORARIOS
	q_listschedulerange := `INSERT INTO ListScheduleRange(idcarta,idschedulemain,idbusiness,starttime,endtime,maxorders,timezone) (select * from unnest($1::int[],$2::int[],$3::int[],$4::varchar(6)[],$5::varchar(6)[],$6::int[],$7::varchar(3)[]))`
	if _, err_listschedulerange := tx.Exec(ctx, q_listschedulerange, idcarta_pg_3, idschedulerange_pg_3, idbusiness_pg_3, startime_pg_3, endtime_pg_3, max_orders_3, timezone_pg_3); err_listschedulerange != nil {
		tx.Rollback(ctx)
		return 0, err_listschedulerange
	}

	//INSERTAR LOS DESCUENTOS AUTOMATICOS
	q_automaticdiscount := `INSERT INTO AutomaticDiscount(idbusiness,idcarta,iddiscount,description,discount,typediscount,groupid,classdiscount) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(30)[],$5::decimal(8,2)[],$6::int[],$7::jsonb[],$8::int[]));`
	if _, err_automaticdiscount := tx.Exec(ctx, q_automaticdiscount, idbusiness_pg_4, idcarta_pg_4, iddiscount_pg_4, description_pg_4, discount_pg_4, typediscount_pg_4, groupid_pg_4, classdiscount_pg_4); err_automaticdiscount != nil {
		tx.Rollback(ctx)
		return 0, err_automaticdiscount
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return 0, err_commit
	}

	return idcarta, nil
}
