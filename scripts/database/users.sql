create table users (
    id serial,
    name varchar(80) not null,
    login varchar(100) not null unique,
    password varchar(200) not null,
    last_login timestamp default current_timestamp,
    created_at timestamp default current_timestamp,
    modified_at timestamp not null,
    deleted bool not null default false,
    primary key (id)
)