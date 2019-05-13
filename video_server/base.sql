create database video_server;

use video_server;
create table users (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    login_name VARCHAR(64),
    pwd TEXT NOT NULL,
    UNIQUE KEY unique_name (login_name)
);

create table video_info (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
    author_id INT UNSIGNED,
    name TEXT,
    display_ctime TEXT,
    create_time DATETIME
);

create table comments (
	id VARCHAR(64) NOT NULL PRIMARY KEY,
    video_id VARCHAR(64),
    author_id INT UNSIGNED,
    content TEXT,
    time DATETIME
);

create table sessions (
	session_id TINYTEXT NOT NULL PRIMARY KEY,
    login_name VARCHAR(64),
	TTL TINYTEXT
);