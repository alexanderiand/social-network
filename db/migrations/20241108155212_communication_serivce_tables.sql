-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS groups (
    id bigserial primary key not null,
    name varchar(255) not null,
    info text,
    group_avatar text,
    admin_id bigint not null references users(id),
    moderators_ids bigint[],
    subscribers_ids bigint[],
    inviteds_ids bigint[],
    last_post_id bigint,
    last_post json,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS group_events (
    id uuid primary key not null,
    title varchar(255) not null,
    description text not null,
    day_time timestamptz not null,
    going_users_ids bigint[],
    not_goring_users_ids bigint[],
    group_id bigint not null references groups(id),
    author_id bigint not null references users(id),
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS group_chats (
    id uuid primary key not null,
    name varchar(255) not null,
    chat_avatar text,
    group_id bigint not null references groups(id),
    admin_id bigint not null references users(id),
    moderators_ids bigint[],
    members_ids bigint[],
    last_msg json,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS grchat_messages (
    id uuid primary key not null,
    from_id bigint not null references users(id),
    from_username varchar(255) not null,
    group_id bigint not null references groups(id),
    content text not null,
    is_publish boolean default true,
    is_delivered boolean default false,
    is_read boolean default false,
    created_at timestamptz not null default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS grchat_messages;
DROP TABLE IF EXISTS group_chats;
DROP TABLE IF EXISTS group_events;
DROP TABLE IF EXISTS groups;
-- +goose StatementEnd
