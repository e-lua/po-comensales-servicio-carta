package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Pg_V2_Mo_Insumo_Elements struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Quantity       int                `json:"quantity"`
}

type Pg_V2_Element_With_Stock_External struct {
	IDElement        int                     `json:"id"`
	IDCarta          int                     `json:"idcarta"`
	IDBusiness       int                     `json:"idbusiness"`
	IDCategory       int                     `json:"idcategory"`
	Typefood         string                  `json:"typefood"`
	NameCategory     string                  `json:"namecategory"`
	UrlPhotoCategory string                  `json:"urlphotocategory"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	Stock            int                     `json:"stock"`
	UrlPhoto         string                  `json:"url"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	Date             string                  `json:"date"`
	Costo            float64                 `json:"costo"`
	AvailableOrders  bool                    `json:"availableorders"`
}

type V2_Pg_Category struct {
	IDCategory   int    `json:"idcategory"`
	NameCategory string `json:"namecategory"`
}

type V2_Pg_Categories_Elements struct {
	Category V2_Pg_Category                      `json:"category"`
	Elements []Pg_V2_Element_With_Stock_External `json:"elements"`
}
