CREATE TABLE IF NOT EXISTS users_exam (
    id text PRIMARY KEY,
    firstname text NOT NULL,
    lastname text NOT NULL,
    email text unique NOT NULL,
    password text NOT NULL ,
    refresh_token TEXT,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp);