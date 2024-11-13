-- +goose Up
-- +goose StatementBegin
-- SSO SERVICE TABLES
CREATE TABLE IF NOT EXISTS users (
    id bigserial primary key not null,
    first_name varchar(64) not null,
    last_name varchar(64) not null,
    date_of_birth timestamp not null,
    email varchar(50) not null, 
    avatar text,
    nickname text,
    about_me text,
    password_hash text not null,
    groups_ids bigint[],
    groupchats_ids bigint[],
    pvchats_ids bigint[],
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS sessions (
    id uuid primary key not null,
    user_id bigint not null references users(id),
    expires_at timestamptz not null,
    created_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS jwt_tokens (
    id uuid primary key not null,
    user_id bigint not null references users(id),
    refresh_token text not null,
    expires_at timestamptz not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS profiles (
    id bigserial primary key not null,
    user_id bigint not null references users(id),
    user_info text,
    is_private boolean default false,
    followings_ids bigint[],
    followers_ids bigint[],
    profile_posts_ids bigint[],
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp

);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions;
DROP TABLE IF EXISTS jwt_tokens;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
