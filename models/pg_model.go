package models

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
	IDElement        int     `json:"id"`
	IDBusiness       int     `json:"idbusiness"`
	IDCategory       int     `json:"idcategory"`
	NameCategory     string  `json:"namecategory"`
	UrlPhotoCategory string  `json:"urlphotocategory"`
	Name             string  `json:"name"`
	Price            float32 `json:"price"`
	Description      string  `json:"description"`
	TypeMoney        int     `json:"typemoney"`
	Stock            int     `json:"stock"`
	UrlPhoto         string  `json:"url"`
}

type Pg_ScheduleList struct {
	IDSchedule     int    `json:"idschedule"`
	Starttime      string `json:"starttime"`
	Endtime        string `json:"endtime"`
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
