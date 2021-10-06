DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS comments;

create table authors
(
    Id         serial PRIMARY KEY,
    first_name varchar(50) not null,
    last_name  varchar(50) not null,
    email      varchar(50) not null
);

CREATE TABLE articles
(
    id         serial PRIMARY KEY,
    author_id  bigint        not null,
    title      varchar(100)  not null,
    body       varchar(1000) not null,
    created_on Date,
    foreign key (author_id) references Authors (id)
);

CREATE TABLE comments
(
    id         serial primary key,
    content    varchar(200) not null,
    article_id bigint       not null,
    foreign key (article_id)
        references articles (id)
)