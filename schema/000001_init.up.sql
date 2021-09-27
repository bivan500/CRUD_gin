CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE book
(
    id serial not null unique,
	name   varchar(255) not null,
	author varchar(255) not null,
	pages  int    not null
);

CREATE TABLE books_in_read_list
(
    id serial not null unique,
    book int references book (id) on delete cascade not null
);

CREATE TABLE read_lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255),
    books_to_read int references books_in_read_list (id) on delete cascade not null
);

CREATE TABLE users_read_list
(
    id serial not null unique,
    users_id int references users (id) on delete cascade not null,
    read_list_id int references read_lists (id) on delete cascade not null
);






