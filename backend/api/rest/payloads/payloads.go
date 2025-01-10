package payloads

import (
	"flower-management/contracts"
	"time"
)

type CreateFlowerPayload struct {
	Name           string                     `validate:"required"`
	PackingOptions []contracts.PackingOptions `validate:"required,min=1,dive"`
}

type CreateProductPayload struct {
	Name        string `validate:"required"`
	Description string
}

type CreateEventPayload struct {
	Name        string    `validate:"required"`
	Date        time.Time `validate:"required"`
	Phone       string
	Email       string
	Address     string `validate:"required"`
	Description string `validate:"required"`
}

type EditFlowerPayload struct {
	ID   string `validate:"required,uuid"`
	Name string
}

type EditProductPayload struct {
	ID          string `validate:"required,uuid"`
	Name        string
	Description string
}

type EditEventPayload struct {
	ID          string `validate:"required,uuid"`
	Name        string
	Date        time.Time
	Phone       string
	Email       string
	Address     string
	Description string
}

type DeleteFlowerPayload struct {
	ID string `validate:"required,uuid"`
}

type DeleteProductPayload struct {
	ID string `validate:"required,uuid"`
}

type DeleteEventPayload struct {
	ID string `validate:"required,uuid"`
}

type GetFilteredFlowersPayload struct {
	Name                  string
	NumOfFlowersInPackage int
}

type GetFilteredProductsPayload struct {
	Name        string
	Description string
}

type GetFilteredEventsPayload struct {
	Name        string
	Date        time.Time
	Address     string
	Description string
}

type AddFlowersToProductPayload struct {
	ProductID string                      `json:"product_id" validate:"required,uuid"`
	Flowers   []contracts.FlowerInProduct `json:"flowers" validate:"required,dive"`
}

type AddProductsToEventPayload struct {
	EventID  string                     `json:"event_id" validate:"required,uuid"`
	Products []contracts.ProductInEvent `json:"products" validate:"required,dive"`
}
