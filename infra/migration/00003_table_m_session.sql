-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_session
(
    id         SERIAL primary key,
    user_id    VARCHAR(50),
    token      TEXT,
    device     VARCHAR(255),
    login_at   DECIMAL,
    ip         VARCHAR(25),
    created_at DECIMAL,
    created_by VARCHAR(50),
    updated_at DECIMAL,
    updated_by VARCHAR(50),
    deleted_at DECIMAL,
    deleted_by VARCHAR(50),
    CONSTRAINT fk_m_user FOREIGN KEY (user_id) REFERENCES m_user (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_session;
-- +goose StatementEnd
