CREATE DATABASE IF NOT EXISTS prot;

USE prot;

CREATE TABLE IF NOT EXISTS user
(
    id        int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username  varchar(20)                    NOT NULL,
    nickname  varchar(20),
    password  varchar(20),
    role      tinyint,
    status    tinyint,
    email     varchar(20),
    phone     varchar(20),
    create_by int,
    create_at datetime,
    update_by int,
    update_at datetime,
    delete_by int,
    delete_at datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO user (username)
values ('admin');
INSERT INTO user (username)
values ('bugong');


CREATE TABLE IF NOT EXISTS tag
(
    id        smallint PRIMARY KEY AUTO_INCREMENT NOT NULL,
    code      varchar(20)                         NOT NULL,
    name      varchar(20)                         NOT NULL,
    sort      int,
    create_by int,
    create_at datetime,
    update_by int,
    update_at datetime,
    delete_by int,
    delete_at datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO tag (code, name)
values ('life', '生活');


CREATE TABLE IF NOT EXISTS article
(
    id         int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title      varchar(100)                   NOT NULL,
    summary    varchar(100)                   NOT NULL,
    content    text                           NOT NULL,
    tags       varchar(100)                   NOT NULL,
    public     tinyint,
    status     tinyint,
    view_count int,
    create_by  int,
    create_at  datetime,
    update_by  int,
    update_at  datetime,
    delete_by  int,
    delete_at  datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

