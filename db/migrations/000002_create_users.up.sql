CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username varchar (255) UNIQUE NOT NULL,
    password bytea NOT NULL,
    role_id bigint,
    FOREIGN KEY (role_id) references roles(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE,
    base_salary INT,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP
);