CREATE TYPE status AS ENUM ('open', 'closed');

CREATE TABLE invoice (
    invoice_id SERIAL PRIMARY KEY,
    status     status DEFAULT 'open',
    created_at TIMESTAMPZ DEFAULT NOW(),
    issued_at  TIMESTAMPZ DEFAULT NULL,
);

CREATE TABLE invoice_item (
    id          SERIAL PRIMARY KEY,
    invoice_id  INTEGER NOT NULL REFERENCES invoice(invoice_id) ON DELETE CASCADE,
    product_id  INTEGER NOT NULL,
    code        VARCHAR(8) NOT NULL,
    description VARCHAR(255),
    unit_price  NUMERIC(15, 3),
    quantity    INTEGER NOT NULL
);
