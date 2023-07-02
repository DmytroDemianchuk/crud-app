CREATE TABLE musics
(
    id serial not null unique,
    name varchar(255) not null,
    artist varchar(255) not null,
    album varchar(255) not null,
    genre varchar(255) not null,
    released_year varchar(255) not null
);

CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255)not null,
    registered_at timestamptz not null
);

CREATE TABLE refresh_tokens
(
    id serial not null unique,
    user_id varchar(255) not null,
    token varchar(255) not null,
    expires_at timestamptz not null
);