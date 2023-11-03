-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_product
(
    id                  VARCHAR(50) primary key,
    category_product_id VARCHAR(50),
    name                VARCHAR(100),
    price               DECIMAL,
    description         TEXT,
    created_at          DECIMAL,
    created_by          VARCHAR(50),
    updated_at          DECIMAL,
    updated_by          VARCHAR(50),
    deleted_at          DECIMAL,
    deleted_by          VARCHAR(50),
    CONSTRAINT fk_m_category_product FOREIGN KEY (category_product_id) REFERENCES m_category_product (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_product;
-- +goose StatementEnd
