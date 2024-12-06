package mock

import (
	"flower-management/contracts"
	persistency "flower-management/internal/persistency/contracts"
)

type DalMock struct {
	Flowers  []*persistency.Flower
	Products []*persistency.Product
	Events   []*persistency.Event
}

func NewDalMock() persistency.DalInterface {
	return &DalMock{
		Flowers:  []*persistency.Flower{},
		Products: []*persistency.Product{},
		Events:   []*persistency.Event{},
	}
}

func (d *DalMock) CreateFlower(flower *persistency.Flower, packingOptions *[]contracts.PackingOptions) error {
	d.Flowers = append(d.Flowers, flower)
	return nil
}

func (d *DalMock) CreateProduct(product *persistency.Product) error {
	d.Products = append(d.Products, product)
	return nil
}

func (d *DalMock) CreateEvent(event *persistency.Event) error {
	d.Events = append(d.Events, event)
	return nil
}

func (d *DalMock) EditFlower(flower *persistency.Flower) error {
	for i, f := range d.Flowers {
		if f.ID == flower.ID {
			d.Flowers[i] = flower
			return nil
		}
	}

	return nil
}

func (d *DalMock) EditProduct(product *persistency.Product) error {
	for i, p := range d.Products {
		if p.ID == product.ID {
			d.Products[i] = product
			return nil
		}
	}

	return nil
}

func (d *DalMock) EditEvent(event *persistency.Event) error {
	for i, e := range d.Events {
		if e.ID == event.ID {
			d.Events[i] = event
			return nil
		}
	}

	return nil
}

func (d *DalMock) DeleteFlower(id string) error {
	for i, f := range d.Flowers {
		if f.ID == id {
			d.Flowers = append(d.Flowers[:i], d.Flowers[i+1:]...)
			return nil
		}
	}

	return nil
}

func (d *DalMock) DeleteProduct(id string) error {
	for i, p := range d.Products {
		if p.ID == id {
			d.Products = append(d.Products[:i], d.Products[i+1:]...)
			return nil
		}
	}

	return nil
}

func (d *DalMock) DeleteEvent(id string) error {
	for i, e := range d.Events {
		if e.ID == id {
			d.Events = append(d.Events[:i], d.Events[i+1:]...)
			return nil
		}
	}

	return nil
}

func (d *DalMock) GetFilteredFlowers(req *contracts.GetFilteredFlowersRequest) ([]*persistency.Flower, error) {
	flowers := []*persistency.Flower{}

	for _, f := range d.Flowers {
		if req.Name != "" && f.Name != req.Name {
			continue
		}

		flowers = append(flowers, f)
	}

	return flowers, nil
}

func (d *DalMock) GetFilteredProducts(req *contracts.GetFilteredProductsRequest) ([]*persistency.Product, error) {
	products := []*persistency.Product{}

	for _, p := range d.Products {
		if req.Name != "" && p.Name != req.Name {
			continue
		}

		if req.Description != "" && p.Description != req.Description {
			continue
		}

		products = append(products, p)
	}

	return products, nil
}

func (d *DalMock) GetFilteredEvents(req *contracts.GetFilteredEventsRequest) ([]*persistency.Event, error) {
	events := []*persistency.Event{}

	for _, e := range d.Events {
		if req.Name != "" && e.Name != req.Name {
			continue
		}

		if !req.Date.IsZero() && e.Date != req.Date {
			continue
		}

		if req.Address != "" && e.Address != req.Address {
			continue
		}

		if req.Description != "" && e.Description != req.Description {
			continue
		}

		events = append(events, e)
	}

	return events, nil
}

func (d *DalMock) GetEvent(id string) (*persistency.Event, error) {
	for _, e := range d.Events {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, nil
}

func (d *DalMock) GetProduct(id string) (*persistency.Product, error) {
	for _, p := range d.Products {
		if p.ID == id {
			return p, nil
		}
	}

	return nil, nil
}

func (d *DalMock) GetFlower(id string) (*persistency.Flower, error) {
	for _, f := range d.Flowers {
		if f.ID == id {
			return f, nil
		}
	}

	return nil, nil
}

func (d *DalMock) AddFlowersToProduct(req *contracts.AddFlowersToProductRequest) error {
	return nil
}

func (d *DalMock) AddProductsToEvent(req *contracts.AddProductsToEventRequest) error {
	return nil
}

func (d *DalMock) EditFlowersInProduct(req *contracts.AddFlowersToProductRequest) error {
	return nil
}

func (d *DalMock) EditProductsInEvent(req *contracts.AddProductsToEventRequest) error {
	return nil
}

func (d *DalMock) GetProductsFromEvent(eventID string) ([]*persistency.EventProduct, error) {
	return nil, nil
}

func (d *DalMock) GetFlowersFromProduct(productID string) ([]*persistency.FlowerInProduct, error) {
	return nil, nil
}
