
# Social Network NorthVoid Amogys-Network

// ## Database design

// ### User (Auth) or SSO Service

// Users
Table users as U {
  id bigint [pk, not null]
  email varchar [not null, unique]
  first_name varchar [not null]
  last_name varchar [not null]
  date_of_birth datetime [default: 'now', not null]
  avatar  text  [not null, note: "optional"]
  nickname varchar  [note: "optional"]
  about_me text [note: "optional"]
  password_hash text [not null]
  groups_ids text [note: ""]
  pvchats_ids text [note: ""]
  created_at datetime [default: 'now', note: "optional"]
  updated_at datetime   [default: 'now',  note: "optional"]
}

// additional fileds: groups_ids, pvchats_ids, age, gender

// Sessions
Table sessions as S {
  id text [pk, not null, note: "google uuid v4"]
  user_id bigint [ref: > U.id, not null]
  expires_at datetime [not null, default: 'now']
  created_at datatime [not null, default: 'now']
}






// ### Profile Service

// Users Profiles
Table  profiles as P {
  id bigint [pk, not null]
  user_id bigint [ref: > U.id, not null, note: "foreign key to users(id) - the profile owner"]
  user_info text [note: "all the profile owner information, without password, user_id" ]
  is_private_profile bool [default: false, not null]
  following text [note: "this user.id > users_ids as []string -> json -> text"]
  followers text [note: "this user.id > users_ids as []string -> json -> text"]
}

// ### Content Service

// Categories
Table categories as C {
  id bigint [pk, not null]
  title varchar [not null]
  created_at datetime [not null, default: 'now']
}

// Posts
Table posts as Po {
  id bigint [pk, not null]
  title varchar [not null]
  content text [not null]
  image text [not null, note: "path to the image file dirs"]
  tags text [note: "save into db as raw text, but text is json ( []string{}) as text"]
  user_id bingint [ref: > U.id, not null]
  category_id bigint [ref: > C.id, not null]
  is_public bool [not null, default: true]
  is_private bool [not null, default: false]
  is_almost_private bool [not null, default: false]
  almost_private_users_access text [note: "users.ids"]
  likes bigint [note: "count of all post's likes, optional"]
  dislikes bigint [note: "count of all post's dislikes, optional"]
  created_at datetime [not null, default: 'now']
  updated_at datetime [not null, default: 'now']
}

// Comments
Table comments as Co {
  id bigint [pk, not null]
  content text [not null]
  image text [not null]
  tags text [note: "optional"]
  likes bigint [note: "optional"]
  dislikes bigint [note: "optional"]
  user_id bigint [ref: > U.id, not null]
  post_id bigint [ref: > Po.id, not null]
  created_at datetime [not null, default: 'now']
  updated_at datetime [not null, default: 'now']
}

// ### Communication or Chat Service

// Groups
Table groups as G {
  id varchar [pk, not null, note: "google uuid v4"]
  name varchar [not null]
  info text [not null, note: "information about this group"]
  group_avatar text [note: "optional"]
  admin_id bigint [ref: > U.id, not null, note: "user creator id, the admin's of the group"]
  moderators_ids text [note: "moderators of the group, only admin can add a new moderator, moders > []user.id"]
  subscribers_ids text [note: "subcribers > []user.id"]
  inviteds_users text [note: "invited users []user.id, only this group subcribers can invite a new user to the group"]
  last_post text [note: "the last post of this group"]
  created_at datatime [not null, default: 'now']
  updated_at datatime [not null, default: 'now']
}

// Group post > Posts

// Group comments of the gr posts > Comments

// Group Events
Table group_events as Ge {
  id varchar [pk, not null, note: "google uuid v4"]
  title varchar [not null]
  description text [not null]
  day_time datetime [not null, default: 'now']
  options text [not null, note: "two options (going, not going), save into database as json>text"]
  group_id bigint [ref: > G.id, not null]
  author_id bigint [ref: > U.id, not null]
  created_at datetime [not null, default: 'now']
  updated_at datetime [not null, default: 'now']
}

// Groups chats
Table group_chats as Gc {
  id varchar [pk, not null, note: "google uuid v4"]
  name varchar [not null]
  chat_avatar text [note: "optional"]
  group_id bigint [ref: > G.id, not null]
  admin_id bigint [ref: > U.id, not null]
  moderators_ids text [note: "optional, []users.id"]
  members_ids text [note: "all group subscribers"]
  last_msg text [note: "the last message of the group chat"]
  created_at datetime [not null, default: 'now']
  updated_at datetime [not null, default: 'now']
}

// Group Chat Messages
Table grchat_messages as Gm {
  id varchar [pk, not null, note: "google uuid v4"]
  from_id bigint [ref: > U.id, not null]
  from_username varchar [not null]
  group_id bigint [ref: >G.id, not null]
  content text [not null]
  is_publish bool
  is_delivered bool
  is_read bool
  created_at datetime [not null, default: 'now']
}

// Private Chats
Table private_chats as Pc {
  id varchar [pk, not null, note: "google uuid v4"]
  name varchar [not null]
  user_creator_id bigint [ref: > U.id, not null]
  user_invited_id bigint [ref: > U.id, not null]
  last_msg text [note: "the last message of the group chat"]
  created_at datetime [not null, default: 'now']
  updated_at datetime [not null, default: 'now']
}

// Private Chat Messages
Table pvchat_messages as Pm {
  id varchar [pk, not null, note: "google uuid v4"]
  pvchat_id varchar [ref: > Pc.id, not null]
  from_id bigint [ref: > U.id, not null]
  from_username varchar [not null]
  content text [not null]
  is_smile bool [not null, default: false]
  smile_code text
  is_delivered bool
  is_read bool
  created_at datetime [not null, default: 'now']
}

// ### Notification Service

// Events
Table events as Ev {
  id varchar [pk, not null, note: "google uuid v4"]
  event_maker_name varchar [not null]
  event_maker_id text [not null]
  event_dest_user_id bigint [ref: > U.id, not null]
  is_request_to_following bool [not null, default: false]
  following_user_id bigin
  is_group_join_invite_request bool [not null, default: false]
  igroup_id varchar
  is_join_group_request bool [not null, default: false]
  group_id varchar
  is_group_event_for_mems bool [not null, default: false]
  event_id varchar
}

---

## Services

### SSO (Auth) Service:

#### Requirments
  1. Requirments: (Это то что сервис должен уметь вообще)
    - Registration (Identification) - Sign-up
    - Authentification - Sign-in - Token - JWT Token
    - Logout, Sign-out
    - Authorization - roles

    - Provide user information, 
    - Also provide information about use profile
    
  2. Technical realization: (Это то что нужно заказчику, технические требования)
    - Sessions
    - Cookies - Full State Service

    -- State less - don't storage any data about user or other entities into itself
    - JWT Bearer Authentification 


### Content Service
  
#### Requirments
  1. Requirments:



--- Project strucutes

- SSO Service
- Content Service
- Chat Service
- Comunication Service
- Event Service | Notify Service


- Monolita <- Microservices

Image backend - 1



Image frontend - 1



### Back-end
- docs
- cmd
  - social network
    - main.go
  
- pkg
  - config
    config.go

  - sso_service
    config
      config.go
      local.yaml
    main.go
  - content_service
    main.go
  - ...

////////////////////////////////////// Our choice 
  clean architecture 
  internal
    - app
    - trasport_common (controllers) -> request (message, command, event)
      - http
        - rest
      - amqp
      - websocket
      - grpc
      - cli 

    - sso_service
      - usecase
      - entity
    - content_service
      - usecase
      - entity
    - chat_service
      - usecase
      - entity

    - infrastructure (gateway, repository) - data
      - repo
        - storage
          - postgresql
          - mongodb
          - redis
      - webapi
        - services apis
        - content_service
  pkg
    - config
    - logger
    ...
//////////////////////////////////////

  clean arch + DDD
  internal
    - sso_service
      - app
      - transport
        - net interfaces
      - usecase
      - infras
        - repo
        - webapi
        -.....
    - content_serice
      - app
      - transport
        - net interfaces
      - usecase
      - infras
        - repo
        - webapi
        -.....



