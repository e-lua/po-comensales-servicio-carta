package models

import "time"

/*------------------------BASIC DATA FOR SEARCH------------------------*/

type Mo_Business struct {
	Description    string            `bson:"description" json:"description"`
	Name           string            `bson:"name" json:"name"`
	TimeZone       string            `bson:"timezone" json:"timezone"`
	DeliveryRange  string            `bson:"deliveryrange" json:"deliveryrange"`
	Delivery       Mo_Delivery       `bson:"delivery" json:"delivery"`
	Contact        []Mo_Contact      `bson:"contact" json:"contact"`
	DailySchedule  []Mo_Day          `bson:"schedule" json:"schedule"`
	Address        Mo_Address        `bson:"address" json:"address"`
	Banner         []Mo_Banner       `bson:"banners" json:"banners"`
	TypeOfFood     []Mo_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services       []Mo_Service      `bson:"services" json:"services"`
	PaymentMethods []Mo_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
	Uniquename     string            `bson:"uniquename" json:"uniquename"`
}

type Mo_Delivery struct {
	Meters        int    `bson:"meters" json:"meters"`
	Details       string `bson:"details" json:"details"`
	IsRestriction bool   `bson:"isrestriction" json:"isrestriction"`
}

type Mo_Banner struct {
	IdBanner int    `bson:"id" json:"id"`
	UrlImage string `bson:"url" json:"url"`
}

type Mo_Address struct {
	Latitude         float64 `bson:"latitude" json:"latitude"`
	Longitude        float64 `bson:"longitude" json:"longitude"`
	FullAddress      string  `bson:"fulladdress" json:"fulladdress"`
	PostalCode       int     `bson:"postalcode" json:"postalcode"`
	State            string  `bson:"state" json:"state"`
	City             string  `bson:"city" json:"city"`
	ReferenceAddress string  `bson:"referenceaddress" json:"referenceaddress"`
}

type Mo_Day struct {
	IDDia      int    `bson:"id" json:"id"`
	StarTime   string `bson:"starttime" json:"starttime"`
	EndTime    string `bson:"endtime" json:"endtime"`
	IsAvaiable bool   `bson:"available" json:"available"`
}

type Mo_TypeFood struct {
	IDTypeFood int    `bson:"id" json:"id"`
	Name       string `bson:"name" json:"name"`
	UrlImage   string `bson:"url" json:"url"`
	IsAvaiable bool   `bson:"available" json:"available"`
}

type Mo_Service struct {
	IDService  int     `bson:"id" json:"id"`
	Name       string  `bson:"name" json:"name"`
	Price      float32 `bson:"price" json:"price"`
	Url        string  `bson:"url" json:"url"`
	TypeMoney  int     `bson:"typemoney" json:"typemoney"`
	IsAvaiable bool    `bson:"available" json:"available"`
}

type Mo_PaymenthMeth struct {
	IDPaymenth  int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	PhoneNumber string `bson:"phonenumber" json:"phonenumber"`
	Url         string `bson:"url" json:"url"`
	HasNumber   bool   `bson:"hasnumber" json:"hasnumber"`
	IsAvaiable  bool   `bson:"available" json:"available"`
}

type Mo_Contact struct {
	IDContact   int    `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	DataContact string `bson:"data" json:"data"`
	IsAvaiable  bool   `bson:"available" json:"available"`
}

/*===========================*/
/*=========VERSION 2=========*/
/*===========================*/

type Mo_Business_V2 struct {
	Description    string            `bson:"description" json:"description"`
	Name           string            `bson:"name" json:"name"`
	TimeZone       string            `bson:"timezone" json:"timezone"`
	DeliveryRange  string            `bson:"deliveryrange" json:"deliveryrange"`
	Delivery       Mo_Delivery       `bson:"delivery" json:"delivery"`
	Contact        []Mo_Contact      `bson:"contact" json:"contact"`
	DailySchedule  []Mo_Day          `bson:"schedule" json:"schedule"`
	Address        Mo_Address        `bson:"address" json:"address"`
	Banner         []Mo_Banner       `bson:"banners" json:"banners"`
	TypeOfFood     []Mo_TypeFood     `bson:"typeoffood" json:"typeoffood"`
	Services       []Mo_Service      `bson:"services" json:"services"`
	PaymentMethods []Mo_PaymenthMeth `bson:"paymentmethods" json:"paymentmethods"`
	Comments       []interface{}     `bson:"comments" json:"comments"`
	Uniquename     string            `bson:"uniquename" json:"uniquename"`
}

/*------------------------BASIC DATA OF CARTA DIARIA------------------------*/

type Mo_CartaDiaria struct {
	IdBusiness      int       `bson:"idbusiness" json:"idbusiness"`
	Date            string    `bson:"date" json:"date"`
	AvailableOrders bool      `bson:"availableorders" json:"availableorders"`
	Visible         bool      `bson:"visible" json:"visible"`
	IsExported      bool      `bson:"isexported" json:"isexported"`
	DeletedDate     time.Time `bson:"deleteddate" json:"deleteddate"`
}

type Mo_Category struct {
	IDCategory       int    `bson:"idcategory" json:"idcategory"`
	Name             string `bson:"namecategory" json:"namecategory"`
	UrlPhoto         string `bson:"urlphotocategory" json:"urlphotocategory"`
	AmountOfElements int    `bson:"elements" json:"elements"`
}

type Mo_Element_With_Stock_Response struct {
	IDElement        int         `bson:"id" json:"id"`
	IDCarta          int         `bson:"idcarta" json:"idcarta"`
	IDBusiness       int         `bson:"idbusiness" json:"idbusiness"`
	IDCategory       int         `bson:"idcategory" json:"idcategory"`
	Typefood         string      `bson:"typefood" json:"typefood"`
	NameCategory     string      `bson:"namecategory" json:"namecategory"`
	UrlPhotoCategory string      `bson:"urlphotocategory" json:"urlphotocategory"`
	Name             string      `bson:"name" json:"name"`
	Price            float32     `bson:"price" json:"price"`
	Description      string      `bson:"description" json:"description"`
	TypeMoney        int         `bson:"typemoney" json:"typemoney"`
	Stock            int         `bson:"stock" json:"stock"`
	UrlPhoto         string      `bson:"url" json:"url"`
	Insumos          []Pg_Insumo `bson:"insumos"  json:"insumos"`
	AvailableOrders  bool        `bson:"availableorders" json:"availableorders"`
	Costo            float64     `bson:"costo" json:"costo"`
}
