CREATE TABLE musics
(
    id serial not null unique,
    name varchar(255) not null,
    artist varchar(255) not null,
    album varchar(255) not null,
    genre varchar(255) not null,
    released_year varchar(255) not null
);