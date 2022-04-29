package models

import "time"

type Mqtt_View_Element struct {
	IDElement  int       `bson:"idelement" json:"idelement"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}

type Mqtt_Element_With_Stock struct {
	IDElement        int                     `bson:"id" json:"id"`
	IDCarta          int                     `bson:"idcarta" json:"idcarta"`
	IDBusiness       int                     `bson:"idbusiness" json:"idbusiness"`
	IDCategory       int                     `bson:"idcategory" json:"idcategory"`
	Typefood         string                  `bson:"typefood" json:"typefood"`
	NameCategory     string                  `bson:"namecategory" json:"namecategory"`
	UrlPhotoCategory string                  `bson:"urlphotocategory" json:"urlphotocategory"`
	Name             string                  `bson:"name" json:"name"`
	Price            float32                 `bson:"price" json:"price"`
	Description      string                  `bson:"description" json:"description"`
	TypeMoney        int                     `bson:"typemoney" json:"typemoney"`
	Stock            int                     `bson:"stock" json:"stock"`
	UrlPhoto         string                  `bson:"url" json:"url"`
	AvailableOrders  bool                    `bson:"isavailableorders" json:"isavailableorders"`
	IsExported       bool                    `bson:"isexported" json:"isexported"`
	Date             string                  `bson:"date" json:"date"`
	DeletedDate      time.Time               `bson:"deleteddate" json:"deleteddate"`
	Insumos          []Pg_Mo_Insumo_Elements `bson:"insumos"  json:"insumos"`
	Costo            float64                 `bson:"costo" json:"costo"`
}

type Mqtt_Element_With_Stock_export struct {
	IdCarta             int           `json:"idcarta"`
	Elements_with_stock []interface{} `json:"elementswithstock"`
}

type Mqtt_Element_With_Stock_Import struct {
	IdCarta             int           `json:"idcarta"`
	Elements_with_stock []interface{} `json:"elementswithstock"`
}

/*IMPORT DATA*/

type Mqtt_Import_SheduleStock struct {
	Schedule int64 `json:"idschedule"`
	IDCarta  int   `json:"idcarta"`
	Quantity int   `json:"quantity"`
}

type Mqtt_Import_ElementStock struct {
	IDElement int64                   `json:"idelement"`
	IDCarta   int                     `json:"idcarta"`
	Quantity  int                     `json:"quantity"`
	Insumos   []Pg_Mo_Insumo_Elements `json:"insumos"`
}
