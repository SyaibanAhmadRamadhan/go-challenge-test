-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_cart
(
    id          SERIAL primary key,
    user_id     VARCHAR(50),
    total_price DECIMAL,
    created_at  DECIMAL,
    created_by  VARCHAR(50),
    updated_at  DECIMAL,
    updated_by  VARCHAR(50),
    deleted_at  DECIMAL,
    deleted_by  VARCHAR(50),
    CONSTRAINT fk_m_user FOREIGN KEY (user_id) REFERENCES m_user (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_cart;
-- +goose StatementEnd
