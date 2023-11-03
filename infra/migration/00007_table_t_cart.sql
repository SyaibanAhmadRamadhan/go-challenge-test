-- +goose Up
-- +goose StatementBegin
CREATE TABLE t_cart
(
    id              SERIAL primary key,
    cart_id         INT,
    product_id      VARCHAR(50),
    total           INT,
    sub_total_price DECIMAL,
    created_at      DECIMAL,
    created_by      VARCHAR(50),
    updated_at      DECIMAL,
    updated_by      VARCHAR(50),
    deleted_at      DECIMAL,
    deleted_by      VARCHAR(50),
    CONSTRAINT fk_m_product FOREIGN KEY (product_id) REFERENCES m_product (id),
    CONSTRAINT fk_m_cart FOREIGN KEY (cart_id) REFERENCES m_cart (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS t_cart;
-- +goose StatementEnd
