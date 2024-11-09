-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS private_chats (
    id uuid primary key not null,
    name varchar(255),
    user_creator_id bigint not null references users(id),
    user_invited_id bigint not null references users(id),
    last_msg_id uuid,
    last_msg json,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

CREATE TABLE IF NOT EXISTS pvchat_messages (
    id uuid primary key not null,
    pvchat_id uuid not null references private_chats(id),
    from_id bigint not null references users(id),
    from_username varchar(255) not null,
    content text not null,
    is_delivered boolean default false,
    is_read boolean default false,
    created_at timestamptz not null default current_timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pvchat_messages;
DROP TABLE IF EXISTS private_chats;
-- +goose StatementEnd
