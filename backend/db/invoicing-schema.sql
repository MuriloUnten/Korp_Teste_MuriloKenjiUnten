CREATE TYPE status AS ENUM ('open', 'closed');

CREATE TABLE invoice (
    number     SERIAL PRIMARY KEY,
    status     status DEFAULT 'open',
    created_at TIMESTAMPZ DEFAULT NOW(),
    closed_at  TIMESTAMPZ DEFAULT NULL,
);

CREATE TABLE invoice_item (
    id              SERIAL PRIMARY KEY,
    invoice_number  INTEGER NOT NULL REFERENCES invoice(number) ON DELETE CASCADE,
    product_id      INTEGER NOT NULL,
    code            VARCHAR(8) NOT NULL,
    description     VARCHAR(255),
    unit_price      NUMERIC(15, 3),
    quantity        INTEGER NOT NULL
);
