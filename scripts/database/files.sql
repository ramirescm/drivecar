create table files (
    id serial,
    folder_id int,
    owner_id int not null,
    name varchar(100) not null,
    type varchar(50) not null,
    path varchar(250) not null,
    created_at timestamp default current_timestamp,
    modified_at timestamp not null,
    deleted bool not null default false,
    primary key (id),
    constraint fk_folder_file foreign key(folder_id) references folders(id),
    constraint fk_owner foreign key(owner_id) references users(id)
)