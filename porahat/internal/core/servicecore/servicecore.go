package servicecore

import (
	"flower-management/contracts"
	persistency "flower-management/internal/persistency/contracts"
)

type ServiceCore struct {
	DalInstance persistency.DalInterface
}

func NewServiceCore(dalInstance persistency.DalInterface) *ServiceCore {
	return &ServiceCore{
		DalInstance: dalInstance,
	}
}

func (s *ServiceCore) CreateFlower(createFlowerRequest *contracts.CreateFlowerRequest) (string, error) {
	flower := &persistency.Flower{
		Name:                  createFlowerRequest.Name,
		NumOfFlowersInPackage: createFlowerRequest.NumOfFlowersInPackage,
	}
	err := s.DalInstance.CreateFlower(flower)

	return flower.ID, err
}

func (s *ServiceCore) CreateProduct(createProductRequest *contracts.CreateProductRequest) (string, error) {
	product := &persistency.Product{
		Name:        createProductRequest.Name,
		Description: createProductRequest.Description,
	}
	err := s.DalInstance.CreateProduct(product)

	return product.ID, err
}

func (s *ServiceCore) CreateEvent(createEventRequest *contracts.CreateEventRequest) (string, error) {
	event := &persistency.Event{
		Name:        createEventRequest.Name,
		Date:        createEventRequest.Date,
		Phone:       createEventRequest.Phone,
		Email:       createEventRequest.Email,
		Address:     createEventRequest.Address,
		Description: createEventRequest.Description,
	}
	err := s.DalInstance.CreateEvent(event)

	return event.ID, err
}

func (s *ServiceCore) EditFlower(editFlowerRequest *contracts.EditFlowerRequest) error {
	flower := &persistency.Flower{
		ID:                    editFlowerRequest.ID,
		Name:                  editFlowerRequest.Name,
		NumOfFlowersInPackage: editFlowerRequest.NumOfFlowersInPackage,
	}

	return s.DalInstance.EditFlower(flower)
}

func (s *ServiceCore) EditProduct(editProductRequest *contracts.EditProductRequest) error {
	product := &persistency.Product{
		ID:          editProductRequest.ID,
		Name:        editProductRequest.Name,
		Description: editProductRequest.Description,
	}
	err := s.DalInstance.EditProduct(product)

	return err
}

func (s *ServiceCore) EditEvent(editEventRequest *contracts.EditEventRequest) error {
	event := &persistency.Event{
		ID:          editEventRequest.ID,
		Name:        editEventRequest.Name,
		Date:        editEventRequest.Date,
		Phone:       editEventRequest.Phone,
		Email:       editEventRequest.Email,
		Address:     editEventRequest.Address,
		Description: editEventRequest.Description,
	}

	return s.DalInstance.EditEvent(event)
}

func (s *ServiceCore) DeleteFlower(id string) error {
	return s.DalInstance.DeleteFlower(id)
}

func (s *ServiceCore) DeleteProduct(id string) error {
	return s.DalInstance.DeleteProduct(id)
}

func (s *ServiceCore) DeleteEvent(id string) error {
	return s.DalInstance.DeleteEvent(id)
}

func (s *ServiceCore) GetFilteredFlowers(req *contracts.GetFilteredFlowersRequest) ([]*persistency.Flower, error) {
	return s.DalInstance.GetFilteredFlowers(req)
}

func (s *ServiceCore) GetFilteredProducts(req *contracts.GetFilteredProductsRequest) ([]*persistency.Product, error) {
	return s.DalInstance.GetFilteredProducts(req)
}

func (s *ServiceCore) GetFilteredEvents(req *contracts.GetFilteredEventsRequest) ([]*persistency.Event, error) {
	return s.DalInstance.GetFilteredEvents(req)
}

func (s *ServiceCore) GetEvent(id string) (*persistency.Event, error) {
	return s.DalInstance.GetEvent(id)
}

func (s *ServiceCore) GetProduct(id string) (*persistency.Product, error) {
	return s.DalInstance.GetProduct(id)
}

func (s *ServiceCore) GetFlower(id string) (*persistency.Flower, error) {
	return s.DalInstance.GetFlower(id)
}

func (s *ServiceCore) AddFlowersToProduct(req *contracts.AddFlowersToProductRequest) error {
	// check if the product exists
	_, err := s.DalInstance.GetProduct(req.ProductID)
	if err != nil {
		return err
	}

	// check if the flowers exist
	for _, flowerInProduct := range *req.Flowers {
		_, err := s.DalInstance.GetFlower(flowerInProduct.FlowerID)
		if err != nil {
			return err
		}
	}
	return s.DalInstance.AddFlowersToProduct(req)
}

func (s *ServiceCore) AddProductsToEvent(req *contracts.AddProductsToEventRequest) error {
	// check if the event exists
	_, err := s.DalInstance.GetEvent(req.EventID)
	if err != nil {
		return err
	}

	// check if the products exist
	for _, productInEvent := range *req.Products {
		_, err := s.DalInstance.GetProduct(productInEvent.ProductID)
		if err != nil {
			return err
		}
	}
	return s.DalInstance.AddProductsToEvent(req)
}

func (s *ServiceCore) EditFlowersInProduct(req *contracts.AddFlowersToProductRequest) error {
	// check if the product exists
	_, err := s.DalInstance.GetProduct(req.ProductID)
	if err != nil {
		return err
	}

	// check if the flowers exist
	for _, flowerInProduct := range *req.Flowers {
		_, err := s.DalInstance.GetFlower(flowerInProduct.FlowerID)
		if err != nil {
			return err
		}
	}
	return s.DalInstance.EditFlowersInProduct(req)
}

func (s *ServiceCore) EditProductsInEvent(req *contracts.AddProductsToEventRequest) error {
	// check if the event exists
	_, err := s.DalInstance.GetEvent(req.EventID)
	if err != nil {
		return err
	}

	// check if the products exist
	for _, productInEvent := range *req.Products {
		_, err := s.DalInstance.GetProduct(productInEvent.ProductID)
		if err != nil {
			return err
		}
	}
	return s.DalInstance.EditProductsInEvent(req)
}

func (s *ServiceCore) GetFlowersInEvent(eventID string) ([]*contracts.GetFlowersInEventResponse, error) {
	// check if the event exists
	_, err := s.DalInstance.GetEvent(eventID)
	if err != nil {
		return nil, err
	}

	products, err := s.DalInstance.GetProductsFromEvent(eventID)
	if err != nil {
		return nil, err
	}

	var flowersInEvent = make(map[string]int)

	for _, product := range products {
		flowers, err := s.DalInstance.GetFlowersFromProduct(product.ProductID)
		if err != nil {
			return nil, err
		}

		for _, flower := range flowers {
			flowersInEvent[flower.FlowerID] = flower.NumOfFlowers
		}
	}

	var response []*contracts.GetFlowersInEventResponse
	for flowerID, numOfFlowers := range flowersInEvent {
		flower, err := s.DalInstance.GetFlower(flowerID)
		if err != nil {
			return nil, err
		}
		numOfPackages := numOfFlowers / flower.NumOfFlowersInPackage
		remindFlowers := numOfFlowers % flower.NumOfFlowersInPackage
		response = append(response, &contracts.GetFlowersInEventResponse{
			FlowerID:      flowerID,
			NumOfFlowers:  numOfFlowers,
			NumOfPackages: numOfPackages,
			RemindFlowers: remindFlowers,
		})
	}

	return response, nil

}
