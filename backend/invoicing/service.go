package invoicing

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

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
	q := `
	SELECT i.number, i.status, i.created_at, i.closed_at,
		item.product_id, item.code, item.description, item.unit_price, item.quantity
	FROM invoice i LEFT JOIN invoice_item item ON i.number = item.invoice_number
	ORDER BY i.number
	`

	rows, err := s.db.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	invoices := make(map[int]*internal.Invoice)
	for rows.Next() {
		var (
			number    int
			status    internal.Status
			createdAt time.Time
			closedAt  time.Time

			productId   *int
			code        *string
			description *string
			unit_price  *float32
			quantity    *int
		)

		err := rows.Scan(
			&number, &status, &createdAt, &closedAt,
			&productId, &code, &description, &unit_price, &quantity,
		)
		if err != nil {
			return nil, err
		}

		// populate the map if its a new invoice
		inv, exists := invoices[number]
		if !exists {
			inv = &internal.Invoice{
				Number: number,
				Status: status,
				CreatedAt: createdAt,
				ClosedAt: &closedAt,
				Items: []internal.InvoiceItemOutput{},
			}
			invoices[number] = inv
		}

		if productId == nil {
			continue
		}
		inv.Items = append(inv.Items, internal.InvoiceItemOutput{
			InvoiceItemInput: internal.InvoiceItemInput{
				ProductId: *productId,
				Quantity: *quantity,
			},
			Code: *code,
			Description: *description,
		})
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	// convert invoice map to slice
	invoiceList := make([]internal.Invoice, 0, len(invoices))
	for _, inv := range invoices {
		invoiceList = append(invoiceList, *inv)
	}

	return invoiceList, nil
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
