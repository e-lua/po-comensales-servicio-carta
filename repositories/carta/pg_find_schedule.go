package repositories

import (
	"context"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Find_ScheduleRange(date string, idbusiness int) ([]models.Pg_ScheduleList, error) {

	db := models.Conectar_Pg_DB()

	q := "SELECT ls.idschedule,CONCAT(c.date::timestamp::date,' ',ls.starttime)::timestamp,CONCAT(c.date::timestamp::date,' ',ls.endtime)::timestamp,ls.maxorders,CONCAT(ls.starttime,' - ',ls.endtime) FROM listschedulerange ls LEFT JOIN carta c ON ls.idcarta=c.idcarta WHERE c.date::timestamp::date=$1::timestamp::date AND ls.idbusiness=$2 AND ls.maxorders>0 AND ls.starttime::time>NOW()::timestamp::time ORDER BY REPLACE(ls.starttime,':','')::int  ASC"
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
