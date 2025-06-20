CREATE TABLE payrolls (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    attendance_period_id BIGINT NOT NULL,
    base_salary INT,
    attendance_days INT,
    attendance_amount int,
    overtime_hours int,
    overtime_amount INT,
    reimbursement_amount INT,
    total_take_home_pay INT,
    created_by BIGINT,
    updated_by BIGINT,
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (attendance_period_id) REFERENCES attendance_periods(id)
        ON DELETE CASCADE ON UPDATE CASCADE
);