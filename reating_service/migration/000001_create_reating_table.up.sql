create table reating (
    id serial primary key,
    post_id int,
    custumer_id int,
    rating int,
    description text,
    created_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP(0) WITH TIME zone NULL
);