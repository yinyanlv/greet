CREATE DATABASE IF NOT EXISTS prot;

USE prot;

CREATE TABLE IF NOT EXISTS user
(
    id        int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    username  varchar(20) UNIQUE             NOT NULL,
    nickname  varchar(20),
    password  varchar(40),
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
VALUES ('admin');
INSERT INTO user (username)
VALUES ('bugong');


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
VALUES ('wenzhang', '文章');
INSERT INTO tag (code, name)
VALUES ('shuoshuo', '说说');
INSERT INTO tag (code, name)
VAlUES ('riji', '日记');


CREATE TABLE IF NOT EXISTS article
(
    id         int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    title      varchar(200)                   NOT NULL,
    summary    varchar(200)                   NOT NULL,
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

