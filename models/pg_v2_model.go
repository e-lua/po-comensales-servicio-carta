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
	Discount         float32                 `json:"discount"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	Additionals      []Pg_Additionals        `json:"additionals"`
	Date             string                  `json:"date"`
	Costo            float64                 `json:"costo"`
	AvailableOrders  bool                    `json:"availableorders"`
}

type Pg_V2_GroupDataDiscount struct {
	ID       int     `json:"id"`
	Quantity float32 `json:"quantity"`
}

type Pg_V2_AutomaticDiscount struct {
	IDCarta             int                    `json:"idcarta"`
	IDAutomaticDiscount int                    `json:"id"`
	Date                string                 `json:"date"`
	IDBusiness          int                    `json:"business"`
	Description         string                 `json:"description"`
	Discount            float32                `json:"discount"`
	TypeDiscount        int                    `json:"type"`
	Group               []Pg_GroupDataDiscount `json:"group"`
	ClassDiscount       int                    `json:"class"`
}

type V2_Pg_Category struct {
	IDCategory   int    `json:"idcategory"`
	NameCategory string `json:"namecategory"`
}

type V2_Pg_Categories_Elements_AutomaticDiscounts struct {
	Category           V2_Pg_Category                      `json:"category"`
	Elements           []Pg_V2_Element_With_Stock_External `json:"elements"`
	AutomaticDiscounts Pg_V2_AutomaticDiscount             `json:"automaticdiscounts"`
}
