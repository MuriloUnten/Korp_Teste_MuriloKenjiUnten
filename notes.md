# Project Spec

## Functionalities

### Register Product
- code
- description
- stock level

### Register Invoice
- Sequential numbering
- Status (open | closed)
- Products (multiple products with their quantities)

- Create with sequential numbering and initial status = open

### Issue Invoice
- Update invoice status to closed
- Now allow issuing of closed invoices
- Update product stock level based on invoice

## UI
- Built with Angular
- Clear button to issue invoice
- Show processing indicator when issuing invoice

## Requirements

### Microsservice architecture
At least these 2 microsservices:
- Product service (manage products and stock levels)
- Billing service (manage invoices)

### Failure handling
- Cenario where a microsservice fails
- Recover from failure
- Provide user with feedback about the failure

## Optional requirements

- Concurrency management
- Use of AI
- Idempotence

## Endpoints

### Product Service
- GET  /products - // may contain one or more `id` query params
- GET  /products/{id}
- POST /products
- PUT /products/reserve - // reserve products in batches
- PUT /products/consume - // consume products in batches

### Invoicing Service
- GET  /invoices
- GET  /invoices/{id}
- POST /invoices
- PUT /invoices/{id}/close
