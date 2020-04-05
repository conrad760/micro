package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product is the structure of the data I am playing with
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreateOn    string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeleteOn    string  `json:"-"`
}

// FromJSON will catch any request made that need to be formed into an JSON Object
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validSku)
	return validate.Struct(p)
}

func validSku(fl validator.FieldLevel) bool {
	// three letters three numbers
	re := regexp.MustCompile(`\b[A-z]{3}[0-9]{3}\b`)
	matches := re.FindAllString(fl.Field().String(), -1)
	if len(matches) != 1 {
		return false
	}
	return true
}

// Products is a slice made of Products
type Products []*Product

// ToJSON exports our objects in a JSON format
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of all the products and information about each in a json format
func GetProducts() Products {
	return productList
}

// AddProduct appends new products
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

// UpdateProduct changes existing products
func UpdateProduct(id int, p *Product) error {
	_, position, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	p.UpdateOn = time.Now().UTC().String()
	productList[position] = p
	return nil
}

// ErrProductNotFound returns an Error when a product is not found
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	listprod := productList[len(productList)-1]
	return listprod.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "A coffee drink made with espresso and steamed milk.",
		Price:       3.97,
		SKU:         "lta001",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	}, &Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Italian style coffee drink.",
		Price:       2.50,
		SKU:         "esp002",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	}, &Product{
		ID:          3,
		Name:        "Double Espresso",
		Description: "Double the volume of a single espresso",
		Price:       2.97,
		SKU:         "dou003",
		CreateOn:    time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
	},
}
