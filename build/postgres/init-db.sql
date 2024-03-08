create table Users(
    user_id integer NOT NULL PRIMARY KEY,
    lastname varchar,
    firstname varchar,
    middlename varchar,
    email varchar,
    pswd varchar,
    passport varchar,
    inn varchar, 
    snils varchar,
    birthday date
);

insert into users values ('adminov', 'admin', 'adminovich', 'admin@mail.ru', 'admin', '1234 567890', '12345678', '12345678', '2024-03-08');
