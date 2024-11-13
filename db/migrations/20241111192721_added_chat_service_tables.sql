-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS private_chats (
    id uuid primary key not null,
    name varchar(32),
    user_creator_id bigint not null references users(id),
    user_invited_id bigint not null references users(id),
    last_msg_id  uuid references pvchat_messages(id), 
    created_at not null timestamptz default current_timestamp,
    updated_at not null timestamptz default current_timestamp
);

CREATE TABLE IF NOT EXISTS pvchat_messages (
    id uuid primary key not null, 
    pvchat_id uuid not null references private_chats(id),
    from_id bigint not null references users(id),
    from_username varchar(255) not null,
    content text not null,
    is_dilivered boolean default false,
    is_read boolean default false, 
    created_at not null timestamptz default current_timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS pvchat_messages;
DROP TABLE IF EXISTS private_chats;
-- +goose StatementEnd
