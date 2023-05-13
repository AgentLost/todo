create table users(
    user_id int generated always as identity ,
    username varchar(255),
    email varchar(255),
    password varchar(255),
    primary key (user_id),
    unique (username),
    unique (email)
);

create table task(
    task_id int generated always as identity ,
    user_id int,
    title text,
    description text,
    due timestamp,
    primary key (task_id),
    constraint fk_users
                 foreign key (user_id)
                    references users(user_id) on delete cascade
);