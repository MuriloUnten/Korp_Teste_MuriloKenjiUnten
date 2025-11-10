package invoicing

import (
	"context"
	"os"
	"log"
	"fmt"

	"github.com/MuriloUnten/Korp_Teste_MuriloKenjiUnten/internal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service interface {
	GetInvoices(ctx context.Context) ([]internal.Invoice, error)
	GetInvoiceById(ctx context.Context, id int) (internal.Invoice, error)
	CreateInvoice(ctx context.Context, invoice internal.CreateInvoiceRequest) (internal.Invoice, error)
	CloseInvoice(ctx context.Context, id int) (internal.Invoice, error)
}

type InvoicingService struct {
	db *pgxpool.Pool
	ProductServiceUrl string
}

func NewInvoicingService(productServiceUrl string) Service {
	s := &InvoicingService{
		ProductServiceUrl: productServiceUrl,
	}
	s.initDB()

	return s
}

func (s *InvoicingService) GetInvoices(ctx context.Context) ([]internal.Invoice, error) {
	// TODO Implement
	return make([]internal.Invoice, 0), nil
}

func (s *InvoicingService) GetInvoiceById(ctx context.Context, id int) (internal.Invoice, error) {
	// TODO Implement
	return internal.Invoice{}, nil
}

func (s *InvoicingService) CreateInvoice(ctx context.Context, invoice internal.CreateInvoiceRequest) (internal.Invoice, error) {
	// TODO Implement
	return internal.Invoice{}, nil
}

func (s *InvoicingService) CloseInvoice(ctx context.Context, id int) (internal.Invoice, error) {
	// TODO Implement
	return internal.Invoice{}, nil
}

func (s *InvoicingService) initDB() {
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
