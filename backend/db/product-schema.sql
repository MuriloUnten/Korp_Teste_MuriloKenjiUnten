CREATE TABLE product (
    product_id  SERIAL PRIMARY KEY,
    code        VARCHAR(8) NOT NULL,
    description VARCHAR(255),
    unit_price  NUMERIC(15, 3),
    stock_level INTEGER NOT NULL,
    reserved    INTEGER DEFAULT 0,
    available   INTEGER GENERATED ALWAYS AS (total - reserved)
);
