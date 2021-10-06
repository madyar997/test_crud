DROP TABLE IF EXISTS comments;
DROP TABLE IF EXISTS articles;
DROP TABLE IF EXISTS authors;

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
);

insert into authors VALUES (1, 'Author 1 FirstName', 'Author 1 LastName', 'author1@gmail.com');
insert into authors VALUES (2, 'Author 2 FirstName', 'Author 2 LastName', 'author2@gmail.com');
insert into authors VALUES (3, 'Author 3 FirstName', 'Author 2 LastName', 'author3@gmail.com');

insert into articles values (1, 1, 'Article 1 Title', 'article 1 Body', '06.10.2021');
insert into articles values (2, 2, 'Article 2 Title', 'article 2 Body', '06.10.2021');
insert into articles values (3, 2, 'Article 2 Title', 'article 2 Body', '06.10.2021');

insert into comments values (1, 'comment 1 content', 1);
insert into comments values (2, 'comment 2 content', 2);
insert into comments values (3, 'comment 3 content', 3);