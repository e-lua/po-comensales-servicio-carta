package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pg_R_PaymentMethod struct {
	IDPaymenth  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	HasNumber   bool   `json:"hasnumber"`
	IsAvailable bool   `json:"available"`
}

type Pg_PaymentMethod_X_Business struct {
	IDPaymenth  int
	IDBusiness  int
	IsAvailable bool
}

type Pg_R_Service struct {
	IDservice   int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	IsAvailable bool   `json:"available"`
}

type Pg_R_TypeFood struct {
	IDTypefood  int    `json:"id"`
	Name        string `json:"name"`
	Url         string `json:"url"`
	IsAvailable bool   `json:"available"`
}

type Pg_Category struct {
	IDCarta          int    `json:"idcarta"`
	IDCategory       int    `json:"idcategory"`
	Name             string `json:"namecategory"`
	UrlPhoto         string `json:"urlphotocategory"`
	AmountOfElements int    `json:"elements"`
}

type Pg_Element_With_Stock struct {
	IDElement        int         `json:"id"`
	IDBusiness       int         `json:"idbusiness"`
	IDCategory       int         `json:"idcategory"`
	NameCategory     string      `json:"namecategory"`
	TypeFood         string      `json:"typefood"`
	UrlPhotoCategory string      `json:"urlphotocategory"`
	Name             string      `json:"name"`
	Price            float32     `json:"price"`
	Description      string      `json:"description"`
	TypeMoney        int         `json:"typemoney"`
	Stock            int         `json:"stock"`
	UrlPhoto         string      `json:"url"`
	Insumos          []Pg_Insumo `json:"insumos"`
	Costo            float64     `json:"costo"`
	AvailableOrders  bool        `json:"availableorders"`
}

type Pg_ScheduleList struct {
	IDSchedule     int    `json:"idschedule"`
	Date           string `json:"date"`
	Starttime      string `json:"starttime"`
	Endtime        string `json:"endtime"`
	TimeZone       string `json:"timezone"`
	MaxOrders      int    `json:"maxorders"`
	ShowToComensal string `json:"showtocomensal"`
}

type Pg_ToElement_Mqtt struct {
	IDElement []int `json:"idElement"`
	IDCarta   []int `json:"idCarta"`
	Quantity  []int `json:"Quantity"`
}

type Pg_ToElement_Mqtt_Obj struct {
	IDElement int `json:"idElement"`
	IDCarta   int `json:"idCarta"`
	Quantity  int `json:"Quantity"`
}

type Pg_ToSchedule_Mqtt struct {
	IDSchedule int64 `json:"idSchedule"`
	IDCarta    int   `json:"idCarta"`
	Quantity   int   `json:"Quantity"`
}

type Pg_Insumo struct {
	Insumo   Mo_Insumo_Response `json:"insumo"`
	Quantity int                `json:"quantity"`
}

type Mo_Insumo_Response struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Available      bool               `json:"available"`
	SendToDelete   time.Time          `json:"sendtodelete"`
}

type Mo_Stock struct {
	Price        float64   `json:"price"`
	IdProvider   string    `json:"idprovider"`
	TimeZone     string    `json:"timezone"`
	CreatedDate  time.Time `json:"createdDate"`
	Quantity     int       `json:"quantity"`
	ProviderName string    `json:"providername"`
}
