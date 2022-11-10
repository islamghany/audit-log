CREATE TABLE IF NOT EXISTS logs(
    id bigserial primary key,
    event_name text NOT NULL,
    description text NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT now(),

    -- customer_id not required because the error can occur before 
    -- querying the current user data in the consumer service.
    customer_id bigint
);