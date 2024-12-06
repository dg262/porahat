package rest

import (
	"flower-management/internal/core/config"
	"flower-management/internal/core/servicecore"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RestServer struct {
	port        int
	sizeLimit   int
	headerSize  int
	app         *fiber.App
	service     *servicecore.ServiceCore
	idleTimeout time.Duration
}

func NewRestServer(cfg *config.RestConfig, service *servicecore.ServiceCore) *RestServer { //nolint:lll
	return &RestServer{
		port:        cfg.Port,
		sizeLimit:   cfg.SizeLimit,
		headerSize:  cfg.HeaderSize,
		idleTimeout: time.Duration(cfg.IdleTimeout) * time.Second,
		app:         fiber.New(),
		service:     service,
	}
}

func (r *RestServer) Start() error {
	defineRoutes(r.app, r.service)

	go func() {
		if err := r.app.Listen(":" + strconv.Itoa(r.port)); err != nil {
			panic(err)
		}
	}()

	return nil
}

func defineRoutes(app *fiber.App, service *servicecore.ServiceCore) {
	app.Post("/flower", func(c *fiber.Ctx) error {
		return createFlower(c, service)
	})

	app.Post("/product", func(c *fiber.Ctx) error {
		return createProduct(c, service)
	})

	app.Post("/event", func(c *fiber.Ctx) error {
		return createEvent(c, service)
	})

	app.Put("/flower", func(c *fiber.Ctx) error {
		return editFlower(c, service)
	})

	app.Put("/product", func(c *fiber.Ctx) error {
		return editProduct(c, service)
	})

	app.Put("/event", func(c *fiber.Ctx) error {
		return editEvent(c, service)
	})

	app.Delete("/flower", func(c *fiber.Ctx) error {
		return deleteFlower(c, service)
	})

	app.Delete("/product", func(c *fiber.Ctx) error {
		return deleteProduct(c, service)
	})

	app.Delete("/event", func(c *fiber.Ctx) error {
		return deleteEvent(c, service)
	})

	app.Get("/flowers", func(c *fiber.Ctx) error {
		return getFilteredFlowers(c, service)
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		return getFilteredProducts(c, service)
	})

	app.Get("/events", func(c *fiber.Ctx) error {
		return getFilteredEvents(c, service)
	})

	app.Get("/event/:eventID", func(c *fiber.Ctx) error {
		return getEvent(c, service)
	})

	app.Get("/product/:productID", func(c *fiber.Ctx) error {
		return getProduct(c, service)
	})

	app.Get("/flower/:flowerID", func(c *fiber.Ctx) error {
		return getFlower(c, service)
	})

	app.Post("/product/flowers", func(c *fiber.Ctx) error {
		return addFlowersToProduct(c, service)
	})

	app.Post("/event/products", func(c *fiber.Ctx) error {
		return addProductsToEvent(c, service)
	})

	app.Put("/product/flowers", func(c *fiber.Ctx) error {
		return editFlowersInProduct(c, service)
	})

	app.Put("/event/products", func(c *fiber.Ctx) error {
		return editProductsInEvent(c, service)
	})

	// app.Get("/event/flowers/:eventID", func(c *fiber.Ctx) error {
	// 	return getFlowersInEvent(c, service)
	// })

}
