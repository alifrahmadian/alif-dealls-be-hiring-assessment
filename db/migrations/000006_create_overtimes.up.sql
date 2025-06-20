CREATE TABLE overtimes (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    payroll_id BIGINT,
    date DATE NOT NULL,
    overtime_hours INT NOT NULL,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (payroll_id) REFERENCES payrolls(id)
        ON DELETE SET NULL ON UPDATE CASCADE
);