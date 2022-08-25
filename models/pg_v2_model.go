package models

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
