-- name: CreateKey :one
INSERT INTO keys(
    id,
    user_id,
    token,
    user_agent,
    user_ip,
    expires_at,
    scope
)
VALUES ($1,$2,$3,$4,$5,$6,$7)
RETURNING *;


-- name: GetKey :one
select * from keys
where id= $1 
LIMIT 1;


-- name: DeleteSession :exec
delete from keys
where id = $1;

-- name: DeleteAllSessionForUser :exec
delete from keys
where user_id = $1 AND scope in (sqlc.arg(scopes)::string[]);