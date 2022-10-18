package cartadiaria_anfitrion

import (
	"context"
	"time"

	models "github.com/Aphofisis/po-comensales-servicio-carta/models"
)

func Pg_Delete_Update_AutomaticDiscount(pg_automaticdiscount []models.Pg_V2_AutomaticDiscount, idcarta int, idbusiness int) error {

	//Descuentos automaticos
	var groupid_pg []interface{}
	iddiscount_pg, description_pg, discount_pg, typediscount_pg, idbusiness_pg, classdiscount_pg, idcarta_pg := []int{}, []string{}, []float32{}, []int{}, []int{}, []int{}, []int{}

	for _, autodisc := range pg_automaticdiscount {

		iddiscount_pg = append(iddiscount_pg, autodisc.IDAutomaticDiscount)
		description_pg = append(description_pg, autodisc.Description)
		discount_pg = append(discount_pg, autodisc.Discount)
		typediscount_pg = append(typediscount_pg, autodisc.TypeDiscount)
		idbusiness_pg = append(idbusiness_pg, idbusiness)
		classdiscount_pg = append(classdiscount_pg, autodisc.ClassDiscount)
		idcarta_pg = append(idcarta_pg, idcarta)
		groupid_pg = append(groupid_pg, autodisc.Group)
	}

	//Tiempo limite al contexto
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	//defer cancelara el contexto
	defer cancel()
	db_external := models.Conectar_Pg_DB()

	//BEGIN
	tx, error_tx := db_external.Begin(ctx)
	if error_tx != nil {
		return error_tx
	}

	//ELIMINAR LISTA DE RANGOS HORARIOS
	q_delete_list := `DELETE FROM automaticdiscount WHERE idbusiness=$1 AND idcarta=$2`
	if _, err_update := tx.Exec(ctx, q_delete_list, idbusiness, idcarta); err_update != nil {
		tx.Rollback(ctx)
		return err_update
	}

	//AUTOMATIC DISCOUNT
	q_automaticdiscount := `INSERT INTO automaticdiscount(idbusiness,idcarta,iddiscount,description,discount,typediscount,groupid,classdiscount) (SELECT * FROM unnest($1::int[],$2::int[],$3::int[],$4::varchar(30)[],$5::decimal(8,2)[],$6::int[],$7::jsonb[],$8::int[]));`
	if _, err_automaticdiscount := tx.Exec(ctx, q_automaticdiscount, idbusiness_pg, idcarta_pg, iddiscount_pg, description_pg, discount_pg, typediscount_pg, groupid_pg, classdiscount_pg); err_automaticdiscount != nil {
		tx.Rollback(ctx)
		return err_automaticdiscount
	}

	//TERMINAMOS LA TRANSACCION
	err_commit := tx.Commit(ctx)
	if err_commit != nil {
		tx.Rollback(ctx)
		return err_commit
	}

	return nil
}
