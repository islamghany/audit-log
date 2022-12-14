CREATE TABLE IF NOT EXISTS keys (
    id uuid PRIMARY KEY,
    user_id bigint NOT NULL,
    token text not NULL,
    user_agent text not null,
    user_ip text not null,
    expires_at timestamptz not null,
    scope text NOT NULL ,
    created_at timestamptz not null default (now())
);

ALTER TABLE users ADD  api_key uuid REFERENCES keys(id) NOT null;

ALTER TABLE "keys"
ADD CONSTRAINT "keys_user_fk"
FOREIGN KEY ("user_id")
REFERENCES "users" ("id")
ON DELETE CASCADE;