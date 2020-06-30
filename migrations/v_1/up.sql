CREATE DATABASE IF NOT EXISTS prot;

USE prot;

CREATE TABLE IF NOT EXISTS user
(
    id         varchar(40) PRIMARY KEY NOT NULL,
    username   varchar(20) UNIQUE      NOT NULL,
    nickname   varchar(20),
    password   varchar(40),
    salt       varchar(20),
    role       tinyint UNSIGNED,
    status     tinyint UNSIGNED,
    email      varchar(30),
    phone      varchar(20),
    created_by varchar(40),
    created_at datetime,
    updated_by varchar(40),
    updated_at datetime,
    deleted_by varchar(40),
    deleted_at datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS tag
(
    id         varchar(20) PRIMARY KEY NOT NULL,
    name       varchar(20)             NOT NULL,
    sort       int,
    created_by varchar(40),
    created_at datetime,
    updated_by varchar(40),
    updated_at datetime,
    deleted_by varchar(40),
    deleted_at datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

INSERT INTO tag (id, name)
VALUES ('wenzhang', '文章');

INSERT INTO tag (id, name)
VALUES ('shuoshuo', '说说');

INSERT INTO tag (id, name)
VAlUES ('riji', '日记');

INSERT INTO tag (id, name)
VAlUES ('shenghuo', '生活');

INSERT INTO tag (id, name)
VAlUES ('ganwu', '感悟');

INSERT INTO tag (id, name)
VAlUES ('it', 'IT');

INSERT INTO tag (id, name)
VAlUES ('go', 'Go');

INSERT INTO tag (id, name)
VAlUES ('javascript', 'JavaScript');

INSERT INTO tag (id, name)
VAlUES ('c', 'C');

INSERT INTO tag (id, name)
VAlUES ('rust', 'Rust');

INSERT INTO tag (id, name)
VAlUES ('java', 'Java');

INSERT INTO tag (id, name)
VAlUES ('python', 'Python');

INSERT INTO tag (id, name)
VAlUES ('typescript', 'Typescript');

INSERT INTO tag (id, name)
VAlUES ('frontend', '前端');

INSERT INTO tag (id, name)
VAlUES ('backend', '后端');

INSERT INTO tag (id, name)
VAlUES ('mysql', 'MySql');

INSERT INTO tag (id, name)
VAlUES ('nosql', 'NoSql');


CREATE TABLE IF NOT EXISTS article
(
    id         varchar(40) PRIMARY KEY NOT NULL,
    title      varchar(200)            NOT NULL,
    summary    varchar(200)            NOT NULL,
    content    text                    NOT NULL,
    public     tinyint UNSIGNED,
    status     tinyint UNSIGNED,
    view_count int UNSIGNED,
    created_by varchar(40),
    created_at datetime,
    updated_by varchar(40),
    updated_at datetime,
    deleted_by varchar(40),
    deleted_at datetime
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;


CREATE TABLE IF NOT EXISTS article_tag
(
    id         int PRIMARY KEY AUTO_INCREMENT NOT NULL,
    article_id varchar(40),
    tag_id     varchar(20)
#     CONSTRAINT article_tag_ibfk_1 FOREIGN KEY (article_id) REFERENCES article (id),
#     CONSTRAINT article_tag_ibfk_2 FOREIGN KEY (tag_id) REFERENCES tag (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

