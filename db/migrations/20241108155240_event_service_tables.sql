-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id uuid primary key not null,
    event_type varchar(64) not null check(event_type in ('req_to_flw', 'group_join_invite', 'group_join_req', 'group_event')),
    payload json not null,
    status varchar(30) not null check(status in ('new', 'processing', 'done')),
    reserved_at timestamptz not null default current_timestamp,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
