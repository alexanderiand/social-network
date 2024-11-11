-- +goose Up
-- +goose StatementBegin
    CREATE TABLE IF NOT EXISTS users (
        id bigserial primary key not null, -- integer autoincrement 
        first_name varchar(100) not null,
        last_name varchar(100) not null,
        email varchar(60) not null unique,
        date_of_birth timestamp not null,
        avatar text, -- optional
        nickname varchar(64), -- optional
        about_me text varchar(255),
        password_hash text not null,
        groups_ids bigint[], 
        groupchat_ids bigint[],
        pvchats_ids bigint[], 
        created_at timestamptz not null default current_timestamp, 
        updated_at timestamptz not null default current_timestamp 
    );

    -- sessions
    CREATE TABLE IF NOT EXISTS sessions (
        id bigserial primary key not null,
        user_id bigint not null references users(id),
        expires_at timestamptz not null,
        created_at timestamptz not null default current_timestamp
    );

    -- profiles
    CREATE TABLE IF NOT EXISTS profiles (
        id bigserial primary key not null, 
        user_id bigint not null references users(id),
        user_info text, 
        is_private_profile boolean default false,
        following bigint[],
        followers bigint[] 
    );
    
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS profiles; -- Сначала удаляем таблицы которые имею рефернсы, а потом таблицу на которую ссылаются  
DROP TABLE IF EXISTS sessions; 
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
