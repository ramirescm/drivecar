create table folders (
    id serial,
    parent_id int,
    name varchar(60) not null,
    created_at timestamp default current_timestamp,
    modified_at timestamp not null,
    deleted bool not null default false,
    primary key (id),
    constraint fk_folders foreign key(parent_id) references folders(id)
)