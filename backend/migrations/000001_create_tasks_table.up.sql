-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    position VARCHAR(50) NOT NULL,
    department VARCHAR(50) NOT NULL,
    hire_date DATE NOT NULL,
    salary INT NOT NULL
);
-- +goose StatementEnd