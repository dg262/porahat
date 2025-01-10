package contracts

import (
	"flower-management/contracts"
	"time"
)

type Flower struct {
	ID   string
	Name string
}

type FlowerInProduct struct {
	FlowerID     string
	ProductID    string
	NumOfFlowers int
}

type FlowerPackageOptions struct {
	FlowerID     string
	NumOfFlowers int
	Price        float64
}

type Product struct {
	ID          string
	Name        string
	Description string
}

type Event struct {
	ID          string
	Name        string
	Date        time.Time
	Phone       string
	Email       string
	Address     string
	Description string
}

type EventProduct struct {
	EventID   string
	ProductID string
	Quantity  int
}

type DalInterface interface {
	CreateFlower(flower *Flower, packingOptions *[]contracts.PackingOptions) error
	CreateProduct(product *Product) error
	CreateEvent(event *Event) error
	EditFlower(flower *Flower) error
	EditProduct(product *Product) error
	EditEvent(event *Event) error
	DeleteFlower(id string) error
	DeleteProduct(id string) error
	DeleteEvent(id string) error
	GetFilteredFlowers(req *contracts.GetFilteredFlowersRequest) ([]*Flower, error)
	GetFilteredProducts(req *contracts.GetFilteredProductsRequest) ([]*Product, error)
	GetFilteredEvents(req *contracts.GetFilteredEventsRequest) ([]*Event, error)
	GetEvent(id string) (*Event, error)
	GetProduct(id string) (*Product, error)
	GetFlower(id string) (*Flower, error)
	AddFlowersToProduct(req *contracts.AddFlowersToProductRequest) error
	AddProductsToEvent(req *contracts.AddProductsToEventRequest) error
	EditFlowersInProduct(req *contracts.AddFlowersToProductRequest) error
	EditProductsInEvent(req *contracts.AddProductsToEventRequest) error
	GetProductsFromEvent(eventID string) ([]*EventProduct, error)
	GetFlowersFromProduct(productID string) ([]*FlowerInProduct, error)
	GetFlowerPackingOptions(flowerID string) ([]*FlowerPackageOptions, error)
}
