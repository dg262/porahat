package dal

import (
	"context"
	"flower-management/contracts"
	"flower-management/internal/core/config"
	persistency "flower-management/internal/persistency/contracts"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(ctx context.Context, dalConfig *config.DalConfig) (persistency.DalInterface, error) {
	// connect to the database
	pool, err := pgxpool.New(ctx, dalConfig.Url)
	if err != nil {
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return &Dal{
		pool: pool,
	}, nil
}

func (d *Dal) CreateFlower(flower *persistency.Flower, packingOptions *[]contracts.PackingOptions) error {
	tx, err := d.pool.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	flower.ID = uuid.New().String()
	flowerQueryEnumerator, flowerParameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	flowerParameterEnumerator.AppendParameter("id", flower.ID)
	flowerParameterEnumerator.AppendParameter("name", flower.Name)

	// Construct the SQL query
	query := fmt.Sprintf(
		"INSERT INTO flowers (%s) VALUES (%s)",
		flowerParameterEnumerator.GetColumns(),
		flowerParameterEnumerator.GetParameters(),
	)

	// Execute the query
	_, err = tx.Exec(context.Background(), query, flowerQueryEnumerator.args...)
	if err != nil {
		tx.Rollback(context.Background())
		return fmt.Errorf("failed to create flower: %w", err)
	}

	for _, packingOption := range *packingOptions {
		packingOptionQueryEnumerator, packingOptionParameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
		packingOptionParameterEnumerator.AppendParameter("flower_id", flower.ID)
		packingOptionParameterEnumerator.AppendParameter("num_of_flowers", packingOption.Quantity)
		packingOptionParameterEnumerator.AppendParameter("price", packingOption.Price)

		// Construct the SQL query
		packingOptionQuery := fmt.Sprintf(
			"INSERT INTO flower_package_options (%s) VALUES (%s)",
			packingOptionParameterEnumerator.GetColumns(),
			packingOptionParameterEnumerator.GetParameters(),
		)

		// Execute the query
		_, err = tx.Exec(context.Background(), packingOptionQuery, packingOptionQueryEnumerator.args...)
		if err != nil {
			tx.Rollback(context.Background())
			return fmt.Errorf("failed to create packing option: %w", err)
		}
	}

	err = tx.Commit(context.Background())
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (d *Dal) CreateProduct(product *persistency.Product) error {
	product.ID = uuid.New().String()
	queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	parameterEnumerator.AppendParameter("id", product.ID)
	parameterEnumerator.AppendParameter("name", product.Name)
	parameterEnumerator.AppendParameter("description", product.Description)

	// Construct the SQL query
	query := fmt.Sprintf(
		"INSERT INTO products (%s) VALUES (%s)",
		parameterEnumerator.GetColumns(),
		parameterEnumerator.GetParameters(),
	)

	// Execute the query
	_, err := d.pool.Exec(context.Background(), query, queryEnumerator.args...)
	if err != nil {
		return fmt.Errorf("failed to create product: %w", err)
	}

	return nil
}

func (d *Dal) CreateEvent(event *persistency.Event) error {
	event.ID = uuid.New().String()
	queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	parameterEnumerator.AppendParameter("id", event.ID)
	parameterEnumerator.AppendParameter("name", event.Name)
	parameterEnumerator.AppendParameter("date", event.Date)
	parameterEnumerator.AppendParameter("phone", event.Phone)
	parameterEnumerator.AppendParameter("email", event.Email)
	parameterEnumerator.AppendParameter("address", event.Address)
	parameterEnumerator.AppendParameter("description", event.Description)

	// Construct the SQL query
	query := fmt.Sprintf(
		"INSERT INTO events (%s) VALUES (%s)",
		parameterEnumerator.GetColumns(),
		parameterEnumerator.GetParameters(),
	)

	// Execute the query
	_, err := d.pool.Exec(context.Background(), query, queryEnumerator.args...)
	if err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}

	return nil
}

func (d *Dal) EditFlower(flower *persistency.Flower) error {
	queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	flowerIDParameter := queryEnumerator.Enumerate(flower.ID)

	// Append parameters to the enumerator
	parameterEnumerator.AppendParameter("name", flower.Name)

	// Construct the SQL query
	query := fmt.Sprintf(
		"UPDATE flowers SET %s WHERE id = %s",
		parameterEnumerator.GetAssignedParameters(),
		flowerIDParameter)

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, queryEnumerator.args...)
	if err != nil {
		return fmt.Errorf("failed to edit flower: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("flower with ID %s does not exist", flower.ID)
	}

	return nil
}

func (d *Dal) EditProduct(product *persistency.Product) error {
	queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	productIDParameter := queryEnumerator.Enumerate(product.ID)

	// Append parameters to the enumerator
	parameterEnumerator.AppendParameter("name", product.Name)
	parameterEnumerator.AppendParameter("description", product.Description)

	// Construct the SQL query
	query := fmt.Sprintf(
		"UPDATE products SET %s WHERE id = %s",
		parameterEnumerator.GetAssignedParameters(),
		productIDParameter)

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, queryEnumerator.args...)
	if err != nil {
		return fmt.Errorf("failed to edit product: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product with ID %s does not exist", product.ID)
	}

	return nil
}

func (d *Dal) EditEvent(event *persistency.Event) error {
	queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
	eventIDParameter := queryEnumerator.Enumerate(event.ID)

	// Append parameters to the enumerator
	parameterEnumerator.AppendParameter("name", event.Name)
	parameterEnumerator.AppendParameter("date", event.Date)
	parameterEnumerator.AppendParameter("phone", event.Phone)
	parameterEnumerator.AppendParameter("email", event.Email)
	parameterEnumerator.AppendParameter("address", event.Address)
	parameterEnumerator.AppendParameter("description", event.Description)

	// Construct the SQL query
	query := fmt.Sprintf(
		"UPDATE events SET %s WHERE id = %s",
		parameterEnumerator.GetAssignedParameters(),
		eventIDParameter)

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, queryEnumerator.args...)
	if err != nil {
		return fmt.Errorf("failed to edit event: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("event with ID %s does not exist", event.ID)
	}

	return nil
}

func (d *Dal) DeleteFlower(id string) error {
	query := "DELETE FROM flowers WHERE id = $1"

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete flower: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("flower with ID %s does not exist", id)
	}

	return nil
}

func (d *Dal) DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id = $1"

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product with ID %s does not exist", id)
	}

	return nil
}

func (d *Dal) DeleteEvent(id string) error {
	query := "DELETE FROM events WHERE id = $1"

	// Execute the query
	result, err := d.pool.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete event: %w", err)
	}

	// Check if any rows were affected
	if result.RowsAffected() == 0 {
		return fmt.Errorf("event with ID %s does not exist", id)
	}

	return nil
}

func (d *Dal) GetFilteredFlowers(req *contracts.GetFilteredFlowersRequest) ([]*persistency.Flower, error) {
	query := "SELECT id, name, num_of_flowers_in_package FROM flowers WHERE 1=1"
	enumerator := &parameterEnumerate{}

	query += enumerator.CreateLikeCondition("name", req.Name)

	// Prepare the query with parameters
	rows, err := d.pool.Query(context.Background(), query, enumerator.args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get filtered flowers: %w", err)
	}
	defer rows.Close()

	var flowers []*persistency.Flower

	// Scan the results into a slice of Flower
	for rows.Next() {
		var flower persistency.Flower
		if err := rows.Scan(&flower.ID, &flower.Name); err != nil {
			return nil, fmt.Errorf("failed to scan flower: %w", err)
		}
		flowers = append(flowers, &flower)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over flowers: %w", err)
	}

	return flowers, nil

}

func (d *Dal) GetFilteredProducts(req *contracts.GetFilteredProductsRequest) ([]*persistency.Product, error) {
	query := "SELECT id, name, description FROM products WHERE 1=1"
	enumerator := &parameterEnumerate{}

	query += enumerator.CreateLikeCondition("name", req.Name)
	query += enumerator.CreateLikeCondition("description", req.Description)

	// Prepare the query with parameters
	rows, err := d.pool.Query(context.Background(), query, enumerator.args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get filtered products: %w", err)
	}
	defer rows.Close()

	var products []*persistency.Product

	// Scan the results into a slice of Product
	for rows.Next() {
		var product persistency.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Description); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, &product)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over products: %w", err)
	}

	return products, nil
}

func (d *Dal) GetFilteredEvents(req *contracts.GetFilteredEventsRequest) ([]*persistency.Event, error) {
	query := "SELECT id, name, date, phone, email, address, description FROM events WHERE 1=1"
	enumerator := &parameterEnumerate{}

	query += enumerator.CreateLikeCondition("name", req.Name)
	query += enumerator.CreateExactCondition("date", req.Date)
	query += enumerator.CreateLikeCondition("address", req.Address)
	query += enumerator.CreateLikeCondition("description", req.Description)

	// Prepare the query with parameters
	rows, err := d.pool.Query(context.Background(), query, enumerator.args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get filtered events: %w", err)
	}
	defer rows.Close()

	var events []*persistency.Event

	// Scan the results into a slice of Event
	for rows.Next() {
		var event persistency.Event
		if err := rows.Scan(&event.ID, &event.Name, &event.Date, &event.Phone, &event.Email, &event.Address, &event.Description); err != nil {
			return nil, fmt.Errorf("failed to scan event: %w", err)
		}
		events = append(events, &event)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over events: %w", err)
	}

	return events, nil
}

func (d *Dal) GetFlower(id string) (*persistency.Flower, error) {
	query := "SELECT id, name FROM flowers WHERE id = $1"

	// Execute the query
	row := d.pool.QueryRow(context.Background(), query, id)

	// Create a Flower instance to hold the result
	var flower persistency.Flower

	// Scan the result into the flower instance
	err := row.Scan(&flower.ID, &flower.Name)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("flower with ID %s does not exist", id)
		}
		return nil, fmt.Errorf("failed to get flower: %w", err)
	}

	return &flower, nil
}

func (d *Dal) GetEvent(id string) (*persistency.Event, error) {
	query := "SELECT id, name, date, phone, email, address, description FROM events WHERE id = $1"

	// Execute the query
	row := d.pool.QueryRow(context.Background(), query, id)

	// Create an Event instance to hold the result
	var event persistency.Event

	// Scan the result into the event instance
	err := row.Scan(&event.ID, &event.Name, &event.Date, &event.Phone, &event.Email, &event.Address, &event.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("event with ID %s does not exist", id)
		}
		return nil, fmt.Errorf("failed to get event: %w", err)
	}

	return &event, nil
}

func (d *Dal) GetProduct(id string) (*persistency.Product, error) {
	query := "SELECT id, name, description FROM products WHERE id = $1"

	// Execute the query
	row := d.pool.QueryRow(context.Background(), query, id)

	// Create a Product instance to hold the result
	var product persistency.Product

	// Scan the result into the product instance
	err := row.Scan(&product.ID, &product.Name, &product.Description)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("product with ID %s does not exist", id)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	return &product, nil
}

func (d *Dal) AddFlowersToProduct(req *contracts.AddFlowersToProductRequest) error {
	ctx := context.Background()
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	for _, flower := range *req.Flowers {
		queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
		parameterEnumerator.AppendParameter("product_id", req.ProductID)
		parameterEnumerator.AppendParameter("flower_id", flower.FlowerID)
		parameterEnumerator.AppendParameter("num_of_flowers", flower.NumOfFlowers)

		// Construct the SQL query
		query := fmt.Sprintf(
			"INSERT INTO flower_in_product (%s) VALUES (%s)",
			parameterEnumerator.GetColumns(),
			parameterEnumerator.GetParameters(),
		)

		// Execute the query within the transaction
		_, err = tx.Exec(context.Background(), query, queryEnumerator.args...)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to add flower to product: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil

}

func (d *Dal) AddProductsToEvent(req *contracts.AddProductsToEventRequest) error {
	ctx := context.Background()
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	for _, product := range *req.Products {
		queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
		parameterEnumerator.AppendParameter("event_id", req.EventID)
		parameterEnumerator.AppendParameter("product_id", product.ProductID)
		parameterEnumerator.AppendParameter("quantity", product.Quantity)

		// Construct the SQL query
		query := fmt.Sprintf(
			"INSERT INTO event_product (%s) VALUES (%s)",
			parameterEnumerator.GetColumns(),
			parameterEnumerator.GetParameters(),
		)

		// Execute the query within the transaction
		_, err = tx.Exec(context.Background(), query, queryEnumerator.args...)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to add product to event: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (d *Dal) EditFlowersInProduct(req *contracts.AddFlowersToProductRequest) error {
	ctx := context.Background()
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	for _, flower := range *req.Flowers {
		queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
		parameterEnumerator.AppendParameter("product_id", req.ProductID)
		parameterEnumerator.AppendParameter("flower_id", flower.FlowerID)
		parameterEnumerator.AppendParameter("num_of_flowers", flower.NumOfFlowers)

		// Construct the SQL query
		query := fmt.Sprintf(
			"UPDATE flowers_in_products SET num_of_flowers = %s WHERE product_id = %s AND flower_id = %s",
			parameterEnumerator.GetParameters(),
			queryEnumerator.Enumerate(req.ProductID),
			queryEnumerator.Enumerate(flower.FlowerID),
		)

		// Execute the query within the transaction
		_, err = tx.Exec(context.Background(), query, queryEnumerator.args...)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to edit flower in product: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (d *Dal) EditProductsInEvent(req *contracts.AddProductsToEventRequest) error {
	ctx := context.Background()
	tx, err := d.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	for _, product := range *req.Products {
		queryEnumerator, parameterEnumerator := new(parameterEnumerate).WithParameterEnumerate()
		parameterEnumerator.AppendParameter("event_id", req.EventID)
		parameterEnumerator.AppendParameter("product_id", product.ProductID)
		parameterEnumerator.AppendParameter("quantity", product.Quantity)

		// Construct the SQL query
		query := fmt.Sprintf(
			"UPDATE products_in_events SET quantity = %s WHERE event_id = %s AND product_id = %s",
			parameterEnumerator.GetParameters(),
			queryEnumerator.Enumerate(req.EventID),
			queryEnumerator.Enumerate(product.ProductID),
		)

		// Execute the query within the transaction
		_, err = tx.Exec(context.Background(), query, queryEnumerator.args...)
		if err != nil {
			tx.Rollback(ctx)
			return fmt.Errorf("failed to edit product in event: %w", err)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (d *Dal) GetProductsFromEvent(eventID string) ([]*persistency.EventProduct, error) {
	query := `SELECT event_id, product_id, quantity FROM event_product WHERE event_id = $1`

	rows, err := d.pool.Query(context.Background(), query, eventID)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from event: %w", err)
	}
	defer rows.Close()

	var ProductsInEvent []*persistency.EventProduct

	// Scan the results into a slice of EventProduct
	for rows.Next() {
		var productInEvent persistency.EventProduct
		if err := rows.Scan(&productInEvent.EventID, &productInEvent.ProductID, &productInEvent.Quantity); err != nil {
			return nil, fmt.Errorf("failed to scan EventProduct: %w", err)
		}
		ProductsInEvent = append(ProductsInEvent, &productInEvent)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over products in event: %w", err)
	}

	return ProductsInEvent, nil
}

func (d *Dal) GetFlowersFromProduct(productID string) ([]*persistency.FlowerInProduct, error) {
	query := `SELECT flower_id, product_id, num_of_flowers FROM flower_in_product WHERE product_id = $1`

	rows, err := d.pool.Query(context.Background(), query, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to get flowers from product: %w", err)
	}
	defer rows.Close()

	var FlowersInProduct []*persistency.FlowerInProduct

	// Scan the results into a slice of FlowerInProduct
	for rows.Next() {
		var flowerInProduct persistency.FlowerInProduct
		if err := rows.Scan(&flowerInProduct.FlowerID, &flowerInProduct.ProductID, &flowerInProduct.NumOfFlowers); err != nil {
			return nil, fmt.Errorf("failed to scan FlowerInProduct: %w", err)
		}
		FlowersInProduct = append(FlowersInProduct, &flowerInProduct)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over flowers in product: %w", err)
	}

	return FlowersInProduct, nil
}

func (d *Dal) GetFlowerPackingOptions(flowerID string) ([]*persistency.FlowerPackageOptions, error) {
	query := `SELECT flower_id, num_of_flowers, price FROM flower_package_options WHERE flower_id = $1`

	rows, err := d.pool.Query(context.Background(), query, flowerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get flower packing options: %w", err)
	}
	defer rows.Close()

	var FlowerPackageOptions []*persistency.FlowerPackageOptions

	// Scan the results into a slice of FlowerPackageOptions
	for rows.Next() {
		var flowerPackageOption persistency.FlowerPackageOptions
		if err := rows.Scan(&flowerPackageOption.FlowerID, &flowerPackageOption.NumOfFlowers, &flowerPackageOption.Price); err != nil {
			return nil, fmt.Errorf("failed to scan FlowerPackageOptions: %w", err)
		}
		FlowerPackageOptions = append(FlowerPackageOptions, &flowerPackageOption)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred while iterating over flower packing options: %w", err)
	}

	return FlowerPackageOptions, nil
}
