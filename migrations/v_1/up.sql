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

INSERT INTO tag (id, name, sort)
VALUES ('wenzhang', '文章', 1);

INSERT INTO tag (id, name, sort)
VALUES ('shuoshuo', '说说', 2);

INSERT INTO tag (id, name, sort)
VAlUES ('shenghuo', '生活', 3);

INSERT INTO tag (id, name, sort)
VAlUES ('ganwu', '感悟', 4);

INSERT INTO tag (id, name, sort)
VAlUES ('riji', '日记', 5);

INSERT INTO tag (id, name, sort)
VAlUES ('it', 'IT', 6);

INSERT INTO tag (id, name, sort)
VAlUES ('frontend', '前端', 7);

INSERT INTO tag (id, name, sort)
VAlUES ('backend', '后端', 8);

INSERT INTO tag (id, name, sort)
VAlUES ('javascript', 'JavaScript', 9);

INSERT INTO tag (id, name, sort)
VAlUES ('go', 'Go', 10);

INSERT INTO tag (id, name, sort)
VAlUES ('c', 'C', 11);

INSERT INTO tag (id, name, sort)
VAlUES ('rust', 'Rust', 12);

INSERT INTO tag (id, name, sort)
VAlUES ('python', 'Python', 13);

INSERT INTO tag (id, name, sort)
VAlUES ('java', 'Java', 14);

INSERT INTO tag (id, name, sort)
VAlUES ('typescript', 'TypeScript', 15);

INSERT INTO tag (id, name, sort)
VAlUES ('database', '数据库', 16);

INSERT INTO tag (id, name, sort)
VAlUES ('mysql', 'MySql', 17);


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

