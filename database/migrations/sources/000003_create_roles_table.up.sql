CREATE TABLE roles
(
    id                  SERIAL PRIMARY KEY,
    name                VARCHAR(255) NOT NULL UNIQUE
);

insert into roles (name) values ('admin');
insert into roles (name) values ('manager');
