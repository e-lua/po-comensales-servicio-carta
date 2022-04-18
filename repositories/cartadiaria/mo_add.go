package cartadiaria

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Mo_Add(input_carta models.Mo_CartaDiaria) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	db := models.MongoCN.Database("restoner_cartadiaria")
	col := db.Collection("cartadiaria")

	_, err := col.InsertOne(ctx, input_carta)
	if err != nil {
		return err
	}

	return nil
}
