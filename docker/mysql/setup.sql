
USE movie_archive;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;

CREATE TABLE IF NOT EXISTS users (
    user_id INT(10) NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(150) NOT NULL UNIQUE,
    user_password VARCHAR(150) NOT NULL,
    PRIMARY KEY (user_id)
    ) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

CREATE TABLE IF NOT EXISTS movie (
    movie_id INT(10) NOT NULL AUTO_INCREMENT,
    movie_name VARCHAR(150) NOT NULL UNIQUE,
    movie_description VARCHAR(150) NOT NULL,
    movie_type VARCHAR(150) NOT NULL,
    PRIMARY KEY (movie_id)
    ) ENGINE=InnoDB DEFAULT CHARSET=UTF8MB4;

INSERT INTO users (user_name, user_password) VALUES ('deniz' , '123');

INSERT INTO movie (movie_name, movie_description, movie_type) VALUES ('Harry Potter' , 'bla bla', 'Fantastic');
