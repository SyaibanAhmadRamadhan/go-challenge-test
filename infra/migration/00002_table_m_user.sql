-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_user
(
    id           VARCHAR(50) primary key,
    role_id      INT,
    username     VARCHAR(50),
    email        VARCHAR(100),
    password     VARCHAR(255),
    phone_number VARCHAR(15),
    created_at   DECIMAL,
    created_by   VARCHAR(50),
    updated_at   DECIMAL,
    updated_by   VARCHAR(50),
    deleted_at   DECIMAL,
    deleted_by   VARCHAR(50),
    FOREIGN KEY (role_id) REFERENCES m_role (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_user;
-- +goose StatementEnd
