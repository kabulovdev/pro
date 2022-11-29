create table posts (
    poster_id int,
    id serial primary key,
    description_post text,
    created_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP(0) WITH TIME zone NULL
);

create table medias (
    id serial primary key,
    post_id int references posts (id),
    name text,
    link text,
    type text
);