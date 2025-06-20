CREATE TABLE reimbursements (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    payroll_id BIGINT,
    reimbursement_amount INT NOT NULL,
    date DATE NOT NULL,
    description TEXT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (payroll_id) REFERENCES payrolls(id)
        ON DELETE SET NULL ON UPDATE CASCADE
);