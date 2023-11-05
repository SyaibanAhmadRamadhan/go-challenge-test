-- +goose Up
-- +goose StatementBegin
CREATE TABLE m_role
(
    id          SERIAL primary key,
    name        VARCHAR(25),
    description VARCHAR(255) NULL
);

INSERT INTO m_role (id, name, description)
VALUES (1, 'member', 'desc');
INSERT INTO m_role (id, name, description)
VALUES (2, 'admin', 'admin');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS m_role;
-- +goose StatementEnd
