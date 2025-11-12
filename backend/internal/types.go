package internal

type Product struct {
	Id          int     `json:"id"`
	Code        string  `json:"code"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unitPrice"`
	StockLevel  int     `json:"stockLevel"`
	Available   int     `json:"available"`
}

type CreateProductRequest struct {
	Code        string  `json:"code"`
	Description string  `json:"description"`
	UnitPrice   float32 `json:"unitPrice"`
	StockLevel  int     `json:"stockLevel"`
}

type ReserveProductsRequest struct {
	Products []struct{
		Id     int `json:"id"`
		Amount int `json:"amount"`
	} `json:"products"`
}

type ConsumeProductsRequest struct {
	Products []struct{
		Id     int `json:"id"`
		Amount int `json:"amount"`
	} `json:"products"`
}

type Invoice struct {
	Number int                 `json:"number"`
	Status Status              `json:"status"`
	Items  []InvoiceItemOutput `json:"items"`
}

type InvoiceItemOutput struct {
	InvoiceItemInput
	Code        string `json:"code"`
	Description string `json:"description"`
}

type CreateInvoiceRequest struct {
	Number  int                 `json:"number"`
	Status  Status              `json:"status"`
	Items   []InvoiceItemInput `json:"items"`
}

type InvoiceItemInput struct {
	ProductId int `json:"id"`
	Quantity  int `json:"quantity"`
}

type Status string

const (
	Open   Status = "open"
	Closed Status = "closed"
)
