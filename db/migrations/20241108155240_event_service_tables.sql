-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS events (
    id uuid primary key not null,
    event_maker_username varchar(255),
    event_maker_id bigint,
    event_dest_user_id bigint,
    is_req_to_flw boolean,
    flw_user_id bigint,
    is_group_join_invite boolean,
    inv_group_id bigint,
    is_group_join_request boolean,
    group_id bigint,
    is_group_event boolean,
    event_id uuid,
    ev_group_id bigint
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS events;
-- +goose StatementEnd
