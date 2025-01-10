package rest

import (
	"flower-management/api/rest/payloads"
	"flower-management/contracts"
	"flower-management/internal/core/servicecore"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func createFlower(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var createFlowerPayload payloads.CreateFlowerPayload

	if err := c.BodyParser(&createFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(createFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createFlowerRequest := &contracts.CreateFlowerRequest{
		Name:           createFlowerPayload.Name,
		PackingOptions: &createFlowerPayload.PackingOptions,
	}

	flowerID, err := service.CreateFlower(createFlowerRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Status(fiber.StatusCreated)
	return c.SendString(flowerID)
}

func createProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var createProductPayload payloads.CreateProductPayload

	if err := c.BodyParser(&createProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(createProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createProductRequest := &contracts.CreateProductRequest{
		Name:        createProductPayload.Name,
		Description: createProductPayload.Description,
	}

	productID, err := service.CreateProduct(createProductRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Status(fiber.StatusCreated)
	return c.SendString(productID)
}

func createEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var createEventPayload payloads.CreateEventPayload

	if err := c.BodyParser(&createEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(createEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	createEventRequest := &contracts.CreateEventRequest{
		Name:        createEventPayload.Name,
		Date:        createEventPayload.Date,
		Phone:       createEventPayload.Phone,
		Email:       createEventPayload.Email,
		Address:     createEventPayload.Address,
		Description: createEventPayload.Description,
	}

	eventID, err := service.CreateEvent(createEventRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Status(fiber.StatusCreated)
	return c.SendString(eventID)
}

func editFlower(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var editFlowerPayload payloads.EditFlowerPayload

	if err := c.BodyParser(&editFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(editFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editFlowerRequest := &contracts.EditFlowerRequest{
		ID:   editFlowerPayload.ID,
		Name: editFlowerPayload.Name,
	}

	err := service.EditFlower(editFlowerRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Flower updated successfully")
}

func editProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var editProductPayload payloads.EditProductPayload

	if err := c.BodyParser(&editProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(editProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editProductRequest := &contracts.EditProductRequest{
		ID:          editProductPayload.ID,
		Name:        editProductPayload.Name,
		Description: editProductPayload.Description,
	}

	err := service.EditProduct(editProductRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Product updated successfully")
}

func editEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var editEventPayload payloads.EditEventPayload

	if err := c.BodyParser(&editEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(editEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editEventRequest := &contracts.EditEventRequest{
		ID:          editEventPayload.ID,
		Name:        editEventPayload.Name,
		Date:        editEventPayload.Date,
		Phone:       editEventPayload.Phone,
		Email:       editEventPayload.Email,
		Address:     editEventPayload.Address,
		Description: editEventPayload.Description,
	}

	err := service.EditEvent(editEventRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Event updated successfully")
}

func deleteFlower(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var deleteFlowerPayload payloads.DeleteFlowerPayload

	if err := c.BodyParser(&deleteFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(deleteFlowerPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := service.DeleteFlower(deleteFlowerPayload.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Flower deleted successfully")
}

func deleteProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var deleteProductPayload payloads.DeleteProductPayload

	if err := c.BodyParser(&deleteProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(deleteProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := service.DeleteProduct(deleteProductPayload.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Product deleted successfully")
}

func deleteEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var deleteEventPayload payloads.DeleteEventPayload

	if err := c.BodyParser(&deleteEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(deleteEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err := service.DeleteEvent(deleteEventPayload.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Event deleted successfully")
}

func getFilteredFlowers(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var getFilteredFlowersPayload payloads.GetFilteredFlowersPayload

	if err := c.BodyParser(&getFilteredFlowersPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(getFilteredFlowersPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	getFilteredFlowersRequest := &contracts.GetFilteredFlowersRequest{
		Name: getFilteredFlowersPayload.Name,
	}

	flowers, err := service.GetFilteredFlowers(getFilteredFlowersRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(flowers)
}

func getFilteredProducts(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var getFilteredProductsPayload payloads.GetFilteredProductsPayload

	if err := c.BodyParser(&getFilteredProductsPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(getFilteredProductsPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	getFilteredProductsRequest := &contracts.GetFilteredProductsRequest{
		Name: getFilteredProductsPayload.Name,
	}

	products, err := service.GetFilteredProducts(getFilteredProductsRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(products)
}

func getFilteredEvents(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var getFilteredEventsPayload payloads.GetFilteredEventsPayload

	if err := c.BodyParser(&getFilteredEventsPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(getFilteredEventsPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	getFilteredEventsRequest := &contracts.GetFilteredEventsRequest{
		Name: getFilteredEventsPayload.Name,
	}

	events, err := service.GetFilteredEvents(getFilteredEventsRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(events)
}

func getFlower(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	flowerID := c.Params("flowerID")
	_, err := uuid.Parse(flowerID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid flower ID")
	}

	flower, err := service.GetFlower(flowerID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(flower)
}

func getProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	productID := c.Params("productID")
	_, err := uuid.Parse(productID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid product ID")
	}

	product, err := service.GetProduct(productID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(product)
}

func getEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	eventID := c.Params("eventID")
	_, err := uuid.Parse(eventID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid event ID")
	}

	event, err := service.GetEvent(eventID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(event)
}

func addFlowersToProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var addFlowersToProductPayload payloads.AddFlowersToProductPayload

	if err := c.BodyParser(&addFlowersToProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(addFlowersToProductPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	addFlowersToProductRequest := &contracts.AddFlowersToProductRequest{
		ProductID: addFlowersToProductPayload.ProductID,
		Flowers:   &addFlowersToProductPayload.Flowers,
	}

	err := service.AddFlowersToProduct(addFlowersToProductRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Flowers added to product successfully")
}

func addProductsToEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var addProductsToEventPayload payloads.AddProductsToEventPayload

	if err := c.BodyParser(&addProductsToEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(addProductsToEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	addProductsToEventRequest := &contracts.AddProductsToEventRequest{
		EventID:  addProductsToEventPayload.EventID,
		Products: &addProductsToEventPayload.Products,
	}

	err := service.AddProductsToEvent(addProductsToEventRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Flowers added to product successfully")
}

func editFlowersInProduct(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var editFlowersInProduct payloads.AddFlowersToProductPayload

	if err := c.BodyParser(&editFlowersInProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(editFlowersInProduct); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editFlowersInProductRequest := &contracts.AddFlowersToProductRequest{
		ProductID: editFlowersInProduct.ProductID,
		Flowers:   &editFlowersInProduct.Flowers,
	}

	err := service.EditFlowersInProduct(editFlowersInProductRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("Flowers in product updated successfully")
}

func editProductsInEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	var editProductsInEventPayload payloads.AddProductsToEventPayload

	if err := c.BodyParser(&editProductsInEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(editProductsInEventPayload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editProductsInEventRequest := &contracts.AddProductsToEventRequest{
		EventID:  editProductsInEventPayload.EventID,
		Products: &editProductsInEventPayload.Products,
	}

	err := service.AddProductsToEvent(editProductsInEventRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendString("products in event updated successfully")
}

func getFlowersInEvent(c *fiber.Ctx, service *servicecore.ServiceCore) error {
	eventID := c.Params("eventID")
	_, err := uuid.Parse(eventID)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid event ID")
	}

	flowers, err := service.GetFlowersInEvent(eventID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(flowers)
}
