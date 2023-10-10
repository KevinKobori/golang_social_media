CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE IF EXISTS publications;
DROP TABLE IF EXISTS followers;
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    createdAt timestamp default current_timestamp()
) ENGINE=INNODB;

CREATE TABLE followers(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    primary key(user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE publications(
    id int auto_increment primary key,
    title varchar(50) not null,
    conteudo varchar(300) not null,

    autor_id int not null,
    FOREIGN KEY (autor_id)
    REFERENCES users(id)
    ON DELETE CASCADE,

    curtidas int default 0,
    criadaEm timestamp default current_timestamp
) ENGINE=INNODB;
