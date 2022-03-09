package models

/*------------------------BASIC DATA FOR SEARCH------------------------*/

type Mo_Business struct {
	Name           string            `bson:"name" json:"name"`
	TimeZone       string            `bson:"timezone" json:"timezone"`
	DeliveryRange  string            `bson:"deliveryrange" json:"deliveryrange"`
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
