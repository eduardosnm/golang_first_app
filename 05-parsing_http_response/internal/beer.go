package beerscli

import (
	"encoding/json"
)

// Beer representation of beer into data struct
type Beer struct {
	ProductID int       `json:"id"`
	Name      string    `json:"name"`
	Street    string    `json:"street"`
	City      int       `json:"city"`
	State     string    `json:"state"`
	Type      *BeerType `json:"brewery_type"`
	Phone     string    `json:"phone"`
}

type BeerType int

const (
	Unknown BeerType = iota
	Lager
	Malt
	Ale
	FlavouredMalt
	Stout
	Porter
	NonAlcoholic
)

func (t BeerType) String() string {
	return toString[t]
}

// NewBeerType initialize a type from enum beerTypes
func NewBeerType(t string) *BeerType {
	beerType := toID[t]
	return &beerType
}

var toString = map[BeerType]string{
	Lager:         "Lager",
	Malt:          "Malt",
	Ale:           "Ale",
	FlavouredMalt: "Flavoured Malt",
	Stout:         "Stout",
	Porter:        "Stout",
	NonAlcoholic:  "Non-Alcoholic",
	Unknown:       "unknown",
}

var toID = map[string]BeerType{
	"Lager":          Lager,
	"Malt":           Malt,
	"Ale":            Ale,
	"Flavoured Malt": FlavouredMalt,
	"Stout":          Stout,
	"Porter":         Porter,
	"Non-Alcoholic":  NonAlcoholic,
	"unknown":        Unknown,
}

// UnmarshalJSON convert type from json to beerType
func (t *BeerType) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t = toID[j]
	return nil
}

// BeerRepo definiton of methods to access a data beer
type BeerRepo interface {
	GetBeers() ([]Beer, error)
}

// NewBeer initialize struct beer
func NewBeer(productID int, name, category, brewer, country, price string, beerType *BeerType) (b Beer) {
	b = Beer{
		ProductID: productID,
		Name:      name,
		Category:  category,
		Type:      beerType,
		Brewer:    brewer,
		Country:   country,
		Price:     price,
	}
	return
}
