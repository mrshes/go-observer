CREATE TABLE users (
    id serial primary key,
    login varchar(50) not null,
    email varchar(255) unique not null,
    password varchar(255) not null,
    created_at timestamp not null
)