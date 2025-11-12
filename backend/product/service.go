package product

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/MuriloUnten/Korp_Teste_MuriloKenjiUnten/internal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetProducts(ctx context.Context, ids []int) ([]internal.Product, error)
	GetProductById(ctx context.Context, id int) (internal.Product, error)
	CreateProduct(ctx context.Context, p internal.CreateProductRequest) (internal.Product, error)
	ReserveProducts(ctx context.Context, req internal.ReserveProductsRequest) ([]internal.Product, error)
	ConsumeProducts(ctx context.Context, req internal.ConsumeProductsRequest) ([]internal.Product, error)
}

type ProductService struct {
	db *pgxpool.Pool
}

func NewProductService() Service {
	s := &ProductService{}
	s.initDB()

	return s
}

func (s *ProductService) GetProducts(ctx context.Context, ids []int) ([]internal.Product, error) {
	q := `SELECT p.product_id, p.code, p.description,
	p.unit_price, p.stock_level, p.available
	FROM product p `
	var rows pgx.Rows
	var err error

	if len(ids) != 0 {
		q += " WHERE p.product_id = ANY($1)"
		rows, err = s.db.Query(ctx, q, ids)
	} else {
		rows, err = s.db.Query(ctx, q)
	}
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return nil, internal.InternalError()
		}
	}

	products := make([]internal.Product, 0)
	for rows.Next() {
		var p internal.Product
		rows.Scan(
			&p.Id, &p.Code, &p.Description,
			&p.UnitPrice, &p.StockLevel, &p.Available,
		)
		products = append(products, p)
	}

	return products, nil
}

func (s *ProductService) GetProductById(ctx context.Context, id int) (internal.Product, error) {
	q := `SELECT p.product_id, p.code, p.description,
	p.unit_price, p.stock_level, p.available
	FROM product p WHERE p.product_id = $1`

	var p internal.Product
	row := s.db.QueryRow(ctx, q, id)
	err := row.Scan(
		&p.Id, &p.Code, &p.Description,
		&p.UnitPrice, &p.StockLevel, &p.Available,
	)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return p, internal.NotFound()
		}

		return p, err
	}

	return p, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, p internal.CreateProductRequest) (internal.Product, error) {
	q := `
	INSERT INTO product(code, description, unit_price, stock_level)
	VALUES($1, $2, $3, $4)
	RETURNING product_id, code, description, unit_price, stock_level, available`

	var product internal.Product
	row := s.db.QueryRow(ctx, q, p.Code, p.Description, p.UnitPrice, p.StockLevel)
	err := row.Scan(
		&product.Id, &product.Code, &product.Description,
		&product.UnitPrice, &product.StockLevel, &product.Available,
	)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *ProductService) ReserveProducts(ctx context.Context, req internal.ReserveProductsRequest) ([]internal.Product, error) {
	// TODO Implement
	return nil, nil
}

func (s *ProductService) ConsumeProducts(ctx context.Context, req internal.ConsumeProductsRequest) ([]internal.Product, error) {
	// TODO Implement
	return nil, nil
}

func (s *ProductService) initDB() {
	dbURL := os.Getenv("DB_URL")

	if dbURL == "" {
		log.Fatal("unable to retreive database url from enviromnent variables")
	}

	var err error
	s.db, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("error creating database connection pool", err)
	}

	if err = s.db.Ping(context.Background()); err != nil {
		log.Fatal("error conecting to database", err)
	}

	fmt.Println("database connected")
}
