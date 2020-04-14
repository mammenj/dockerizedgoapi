-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS books
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    title          VARCHAR(255) NOT NULL,
    author         VARCHAR(255) NOT NULL,
    published_date DATE         NOT NULL,
    image_url      VARCHAR(255) NULL,
    description    TEXT         NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS users
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    userid         VARCHAR(255) NOT NULL,
    name           VARCHAR(255),
    password       VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS roles
(
    id             INT UNSIGNED NOT NULL AUTO_INCREMENT,
    roleid         VARCHAR(255) NOT NULL,
    userid         INT UNSIGNED NOT NULL,
    created_at     TIMESTAMP    NOT NULL,
    updated_at     TIMESTAMP    NULL,
    deleted_at     TIMESTAMP    NULL,
    PRIMARY KEY (id),
    CONSTRAINT `fk_user_roles` FOREIGN KEY (userid) references users(id)
);

INSERT INTO books (title, author, description, published_date) VALUES ("Not to corona or not", "john", "Corona book for survivors","1999-12-13"); 
INSERT INTO books (title, author, description, published_date) VALUES ("From hello to hell", "john", "This is ride of my life","1999-12-13"); 
INSERT INTO books (title, author, description, published_date) VALUES ("How to do this and that", "mary", "No good book for no good people","1999-12-13"); 
INSERT INTO books (title, author, description, published_date) VALUES ("War and Peace", "paul", "Without war and no peace:))","1999-12-13"); 
INSERT INTO users (userid, name, password) VALUES ("john", "john mammen", "$2a$04$lvsBcAVRZ4cSIzaKzv5B7OKdpS3ucy9OhGoc3A6EVsOGCSfzhDdQK"); 
INSERT INTO users (userid, name, password) VALUES ("mary", "mary first", "$2a$04$lvsBcAVRZ4cSIzaKzv5B7OKdpS3ucy9OhGoc3A6EVsOGCSfzhDdQK"); 
INSERT INTO users (userid, name, password) VALUES ("paul", "paul second", "$2a$04$lvsBcAVRZ4cSIzaKzv5B7OKdpS3ucy9OhGoc3A6EVsOGCSfzhDdQK"); 
INSERT INTO roles (roleid,userid) VALUES ("admin",1);
INSERT INTO roles (roleid,userid) VALUES ("editor",2);
INSERT INTO roles (roleid,userid) VALUES ("viewer",3); 
INSERT INTO roles (roleid,userid) VALUES ("viewer",1); 


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;

DROP TABLE IF EXISTS books;