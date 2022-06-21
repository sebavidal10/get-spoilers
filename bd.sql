CREATE DATABASE spoilers;
\c spoilers;

create table spoilers (
  id serial primary key,
  content text,
  movie text
);

insert into spoilers (content, movie) values ('Cypher es un traidor', 'The Matrix');
insert into spoilers (content, movie) values ('El personaje de Bruce Willis esta muerto', 'The six sence');
