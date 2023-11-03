-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_category_product
(
    id         VARCHAR(50) primary key,
    name       VARCHAR(100),
    created_at DECIMAL,
    created_by VARCHAR(50),
    updated_at DECIMAL,
    updated_by VARCHAR(50),
    deleted_at DECIMAL,
    deleted_by VARCHAR(50)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_category_product;
-- +goose StatementEnd
