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

type Pg_Category_ToCreate struct {
	IDCarta          int    `json:"idcarta"`
	IDCategory       int    `json:"idcategory"`
	Name             string `json:"namecategory"`
	UrlPhoto         string `json:"urlphotocategory"`
	AmountOfElements int    `json:"elements"`
}

type Pg_Mo_Insumo_Elements struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Name           string             `json:"name"`
	Measure        string             `json:"measure"`
	IDStoreHouse   string             `json:"idstorehouse"`
	NameStoreHouse string             `json:"namestorehouse"`
	Description    string             `json:"description"`
	Stock          []*Mo_Stock        `json:"stock"`
	Quantity       int                `json:"quantity"`
}

type Pg_Element_ToCreate struct {
	IDElement        int                     `json:"id"`
	IDBusiness       int                     `json:"idbusiness"`
	IDCarta          int                     `json:"idcarta"`
	IDCategory       int                     `json:"idcategory"`
	NameCategory     string                  `json:"namecategory"`
	TypeFood         string                  `json:"typefood"`
	UrlPhotoCategory string                  `json:"urlphotocategory"`
	Name             string                  `json:"name"`
	Price            float32                 `json:"price"`
	Description      string                  `json:"description"`
	TypeMoney        int                     `json:"typemoney"`
	Stock            int                     `json:"stock"`
	UrlPhoto         string                  `json:"url"`
	Insumos          []Pg_Mo_Insumo_Elements `json:"insumos"`
	Costo            float64                 `json:"costo"`
}

type Pg_Schedule_ToCreate struct {
	IDSchedule     int    `json:"idschedule"`
	Date           string `json:"date"`
	Starttime      string `json:"starttime"`
	Endtime        string `json:"endtime"`
	TimeZone       string `json:"timezone"`
	MaxOrders      int    `json:"maxorders"`
	ShowToComensal string `json:"showtocomensal"`
}

/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*-------------------------------------CARTA DIARIA ANFITRIONES-------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/
/*---------------------------------------------------------------------------------------*/

type Pg_ScheduleRange_External struct {
	IDSchedule        int64  `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	MinutePerFraction int    `json:"minutesperfraction"`
	StartTime         string `json:"starttime"`
	EndTime           string `json:"endtime"`
	TimeZone          string `json:"timezone"`
	NumberOfFractions int    `json:"numberfractions"`
	MaxOrders         int    `json:"maxOrders"`
}

type Pg_Element_With_Stock_External struct {
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

type Pg_Carta_Found struct {
	Date     time.Time `json:"date"`
	Elements int       `json:"elements"`
}

type Pg_Category_External struct {
	IDCategory       int    `json:"idcategory"`
	Name             string `json:"namecategory"`
	UrlPhoto         string `json:"urlphotocategory"`
	AmountOfElements int    `json:"elements"`
}

type Pg_Carta_External struct {
	IDCarta            int       `json:"id"`
	Date               time.Time `json:"date"`
	AvailableForOrders bool      `json:"availablefororders"`
	Visible            bool      `json:"visible"`
	Elements           int       `json:"elements"`
	ScheduleRanges     int       `json:"scheduleranges"`
}

/*IMPORT DATA*/

type Pg_Schedule struct {
	IDSchedule        int    `json:"idschedule"`
	IDCarta           int    `json:"idcarta"`
	DateRequired      string `json:"daterequired"`
	TimeStartRequired string `json:"starttime"`
	TimeEndRequired   string `json:"endtime"`
	TimeZone          string `json:"timezone"`
}

type V2_Pg_Element struct {
	IDElement   int                     `json:"idelement"`
	IDBusiness  int                     `json:"idbusiness"`
	IDCarta     int                     `json:"idcarta"`
	NameE       string                  `json:"name"`
	IdCategory  int                     `json:"idcategory"`
	Category    string                  `json:"category"`
	TypeFood    string                  `json:"typefood"`
	URLPhoto    string                  `json:"url"`
	Description string                  `json:"description"`
	TypeMoney   int                     `json:"typemoney"`
	UnitPrice   float64                 `json:"unitprice"`
	Quantity    int                     `json:"quantity"`
	Discount    float32                 `json:"discount"`
	Insumos     []Pg_Mo_Insumo_Elements `json:"insumos"`
	Costo       float64                 `json:"costo"`
}

type Import_Data struct {
	Schedule []Pg_Schedule     `json:"schedule"`
	Elements [][]V2_Pg_Element `json:"elements"`
}
