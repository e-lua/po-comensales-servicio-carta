package repositories

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Add_ViewBusiness(idcomensal int, idbusiness int) error {

	db_external := models.Conectar_Pg_DB()

	query := `INSERT INTO View(idcomensal,idtype,date,idbusiness) VALUES ($1,$2,$3,$4)`
	_, err := db_external.Query(context.Background(), query, idcomensal, 1, time.Now(), idbusiness)

	if err != nil {
		return err
	}

	return nil
}

func Pg_Add_ViewMenu(idcomensal int) error {

	db_external := models.Conectar_Pg_DB()

	query := `INSERT INTO View(idcomensal,idtype,date) VALUES ($1,$2,$3)`
	_, err := db_external.Query(context.Background(), query, idcomensal, 2, time.Now())

	if err != nil {
		return err
	}

	return nil
}

func Pg_Add_ViewElement(idcomensal int, idelement int) error {

	db_external := models.Conectar_Pg_DB()

	query := `INSERT INTO View(idcomensal,idtype,date,idelement) VALUES ($1,$2,$3,$4)`
	_, err := db_external.Query(context.Background(), query, idcomensal, 3, time.Now(), idelement)

	if err != nil {
		return err
	}

	return nil
}
