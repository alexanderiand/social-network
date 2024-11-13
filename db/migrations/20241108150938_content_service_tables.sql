-- +goose Up
-- +goose StatementBegin

-- CONTENT SERVICE TABLES
CREATE TABLE IF NOT EXISTS categories (
    id bigserial primary key not null,
    title varchar(255) not null,
    created_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS posts (
    id bigserial primary key not null,
    title varchar(255) not null,
    content text not null,
    image text,
    tags varchar(64)[],
    user_id bigint not null references users(id),
    category_id bigint not null references categories(id),
    is_public boolean default true,
    is_private boolean default false,
    is_almost_private boolean default false,
    almost_private_users_ids bigint[],
    likes bigint,
    dislikes bigint,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS comments (
    id bigserial primary key not null,
    content text not null,
    image text,
    tags varchar(64)[],
    user_id bigint not null references users(id),
    post_id bigint not null references posts(id),
    likes bigint,
    dislikes bigint,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
