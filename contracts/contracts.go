package contracts

import (
	"time"
)

type FlowerInProduct struct {
	FlowerID     string `json:"flower_id"`
	NumOfFlowers int    `json:"num_of_flowers"`
}

type ProductInEvent struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type CreateFlowerRequest struct {
	Name                  string
	NumOfFlowersInPackage int
}

type CreateProductRequest struct {
	Name        string
	Description string
}

type CreateEventRequest struct {
	Name        string
	Date        time.Time
	Phone       string
	Email       string
	Address     string
	Description string
}

type EditFlowerRequest struct {
	ID                    string
	Name                  string
	NumOfFlowersInPackage int
}

type EditProductRequest struct {
	ID          string
	Name        string
	Flowers     *[]FlowerInProduct
	Description string
}

type EditEventRequest struct {
	ID          string
	Name        string
	Date        time.Time
	Phone       string
	Email       string
	Address     string
	Description string
}

type GetFilteredFlowersRequest struct {
	Name                  string
	NumOfFlowersInPackage int
}

type GetFilteredProductsRequest struct {
	Name        string
	Description string
}

type GetFilteredEventsRequest struct {
	Name        string
	Date        time.Time
	Address     string
	Description string
}

type AddFlowersToProductRequest struct {
	ProductID string
	Flowers   *[]FlowerInProduct
}

type AddProductsToEventRequest struct {
	EventID  string
	Products *[]ProductInEvent
}

type EditProductsInEventRequest struct {
	EventID  string
	Products *[]ProductInEvent
}

type GetFlowersInEventResponse struct {
	FlowerID      string
	NumOfFlowers  int
	NumOfPackages int
	RemindFlowers int
}
