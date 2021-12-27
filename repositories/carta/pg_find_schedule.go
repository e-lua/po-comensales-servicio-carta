package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_ScheduleRange(date string, idbusiness int) ([]models.Pg_ScheduleList, error) {

	db := models.Conectar_Pg_DB()

	q := "SELECT ls.idschedule,c.date::date,ls.starttime::time,ls.endtime::time,ls.timezone,ls.maxorders,CONCAT(ls.starttime,' - ',ls.endtime) FROM listschedulerange ls LEFT JOIN carta c ON ls.idcarta=c.idcarta WHERE c.date::date=$1::date AND ls.idbusiness=$2 AND ls.maxorders>0 AND concat(ls.starttime,ls.timezone)::time with time zone>NOW()::time at time zone CONCAT('UTC',(ls.timezone::integer*-1)::varchar(3)) ORDER BY REPLACE(ls.starttime,':','')::int  ASC"
	rows, error_shown := db.Query(context.Background(), q, date, idbusiness)

	//Instanciamos una variable del modelo Pg_TypeFoodXBusiness
	var oScheduleList []models.Pg_ScheduleList

	if error_shown != nil {

		return oScheduleList, error_shown
	}

	//Scaneamos l resultado y lo asignamos a la variable instanciada
	for rows.Next() {
		var oSchedule models.Pg_ScheduleList
		rows.Scan(&oSchedule.IDSchedule, &oSchedule.Starttime, &oSchedule.Endtime, &oSchedule.MaxOrders, &oSchedule.ShowToComensal)
		oScheduleList = append(oScheduleList, oSchedule)
	}

	//Si todo esta bien
	return oScheduleList, nil
}
