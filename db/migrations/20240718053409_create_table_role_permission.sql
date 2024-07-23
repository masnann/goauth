-- +goose Up
-- +goose StatementBegin

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE user_role (
    user_id INTEGER,
    role_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (role_id) REFERENCES roles(id)
);

CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    groups VARCHAR(50) NOT NULL,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE role_permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER,
    permission_id INTEGER,
    status BOOLEAN NOT NULL,
    FOREIGN KEY (role_id) REFERENCES roles(id),
    FOREIGN KEY (permission_id) REFERENCES permissions(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE role_permissions;
DROP TABLE permissions;
DROP TABLE user_role;
DROP TABLE roles;
-- +goose StatementEnd
