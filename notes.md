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
- Stock service (manage products and stock levels)
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

### Stock Service
- /GET  products
- /GET  products/{id}
- /POST products

### Invoicing Service
- /GET  invoices
- /GET  invoices/{id}
- /POST invoices
- /POST invoices/{id}/close
