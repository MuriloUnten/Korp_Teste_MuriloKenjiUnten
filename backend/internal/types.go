package internal

type Product struct {
	Id          int     `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	StockLevel  float32 `json:"stockLevel"`
	Available   float32 `json:"available"`
}

type CreateProductRequest struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	StockLevel  float32 `json:"stockLevel"`
}

type Invoice struct {
	Number int           `json:"number"`
	Status Status        `json:"status"`
	Items  []InvoiceItem `json:"items"`
}

type InvoiceItem struct {
	ProductId   int     `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	Quantity    float32 `json:"quantity"`
}

type CreateInvoiceRequest struct {
	Number  int                 `json:"number"`
	Status  Status              `json:"status"`
	Items   []CreateInvoiceItem `json:"items"`
}

type CreateInvoiceItem  struct {
	ProductId int     `json:"id"`
	Quantity  float32 `json:"quantity"`
}

type Status string

const (
	Open   Status = "open"
	Closed Status = "closed"
)
