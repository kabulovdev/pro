CREATE table custumer_base (
    id serial primary key,
    first_name text,
    last_name text,
    email text,
    password text,
    phonenumber text,
    refresh_token text NOT NULL,
    created_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP(0) WITH TIME zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP(0) WITH TIME zone NULL
);

create table custumer_bio (
    custumer_id int references custumer_base (id),
    bio text
);

create table custumer_address (
    id serial primary key,
    custumer_id int references custumer_base (id),
    street text,
    home_address text
);
