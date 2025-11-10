package product

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/MuriloUnten/Korp_Teste_MuriloKenjiUnten/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetProducts(ctx context.Context, ids []int) ([]internal.Product, error)
	GetProductById(ctx context.Context, id int) (internal.Product, error)
	CreateProduct(ctx context.Context, p internal.CreateProductRequest) (internal.Product, error)
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
	// TODO Implement
	return make([]internal.Product, 0), nil
}

func (s *ProductService) GetProductById(ctx context.Context, id int) (internal.Product, error) {
	// TODO Implement
	return internal.Product{}, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, p internal.CreateProductRequest) (internal.Product, error) {
	// TODO Implement
	return internal.Product{}, nil
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
