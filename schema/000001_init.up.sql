DROP TABLE user_lists;
DROP TABLE books_in_list;
DROP TABLE lists;
DROP TABLE book;
DROP TABLE users;



CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);
INSERT INTO users (username, password_hash) values ('Koocks', '617364313233667364666c7065773132333371777378356a192b7913b04c54574d18c28d46e6395428ab'); -- pass 1
INSERT INTO users (username, password_hash) values ('Ali', '617364313233667364666c7065773132333371777378da4b9237bacccdf19c0760cab7aec4a8359010b0'); -- pass 2
INSERT INTO users (username, password_hash) values ('Abdul', '617364313233667364666c706577313233337177737877de68daecd823babbb58edb1c8e14d7106e83bb'); -- pass 3
INSERT INTO users (username, password_hash) values ('Jopik', '617364313233667364666c70657731323333717773781b6453892473a467d07372d45eb05abc2031647a'); -- pass 4

CREATE TABLE book
(
    id serial not null unique,
	name   varchar(255) not null,
	author varchar(255) not null,
	pages  int    not null
);

INSERT INTO book (name, author, pages) values ('Jopik v ogne', 'Kool Pop', 30); -- pass 4
INSERT INTO book (name, author, pages) values ('Jopik v ogne 2', 'Kool Pop', 30); -- pass 4
INSERT INTO book (name, author, pages) values ('Jopik v ogne 3', 'Kool Pop', 30); -- pass 4
INSERT INTO book (name, author, pages) values ('Jora i Kola', 'Tol Jop', 20); -- pass 4
INSERT INTO book (name, author, pages) values ('ogneLisa', 'Irel Mn', 10); -- pass 4
INSERT INTO book (name, author, pages) values ('Xi', 'Puh', 60); -- pass 4
INSERT INTO book (name, author, pages) values ('Nefrit', 'Jora Krisha', 40); -- pass 4

CREATE TABLE lists
(
    id serial not null unique,
    title varchar(255) not null,
    description varchar(255)
);

INSERT INTO lists (title, description) values ('На лето', 'Соси жопу'); 
INSERT INTO lists (title, description) values ('По кайфу', 'Есть жи');

CREATE TABLE books_in_list
(
    id serial not null unique,
    book_id int references book (id) on delete cascade not null,
    list_id int references lists (id) on delete cascade not null    
);

INSERT INTO books_in_list (book_id, list_id) values (1, 1); 
INSERT INTO books_in_list (book_id, list_id) values (2, 1);
INSERT INTO books_in_list (book_id, list_id) values (3, 1);
INSERT INTO books_in_list (book_id, list_id) values (4, 1);
INSERT INTO books_in_list (book_id, list_id) values (5, 1);
INSERT INTO books_in_list (book_id, list_id) values (2, 2); 
INSERT INTO books_in_list (book_id, list_id) values (4, 2); 


CREATE TABLE user_lists
(
    id serial not null unique,
    users_id int references users (id) on delete cascade not null,
    list_id int references lists (id) on delete cascade not null
);

INSERT INTO user_lists (users_id, list_id) values (1, 1);
INSERT INTO user_lists (users_id, list_id) values (2, 2);





