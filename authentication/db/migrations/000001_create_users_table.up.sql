CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    first_name varchar(35) NOT NULL,
    last_name varchar(35) NOT NULL,
    email text UNIQUE,
    company_name varchar(72),
    activated bool NOT NULL DEFAULT false,
    is_blocked bool NOT NULL DEFAULT false,
    hashed_password bytea NOT NULL,
    password_changed_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);