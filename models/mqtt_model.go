package models

import "time"

type Mqtt_View_Element struct {
	IDElement  int       `bson:"idelement" json:"idelement"`
	IDComensal int       `bson:"idcomensal" json:"idcomensal"`
	Date       time.Time `bson:"date" json:"date"`
}
