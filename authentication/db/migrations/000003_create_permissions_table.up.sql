CREATE TABLE IF NOT EXISTS permissions(
    id bigserial PRIMARY KEY,
    code text NOT NULL
);


CREATE TABLE IF NOT EXISTS users_permissions(
    user_id bigint NOT NULL REFERENCES users on Delete Cascade,
    permission_id bigint NOT NULL REFERENCES permissions on Delete Cascade,
    primary key (user_id, permission_id)
);

-- Add the two permissions to the table.
INSERT INTO permissions (code)
VALUES 
    ('logs:read'),
    ('logs:write');