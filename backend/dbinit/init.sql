-- Drop tables if they exist
drop table if exists meetings;
drop table if exists users;

-- Create users table
create table users (
    id SERIAL primary key,
    email VARCHAR(254) unique not null,
    username VARCHAR(50) not null,
    first_name VARCHAR(255) not null,
    last_name VARCHAR(255) not null,
    password_hash VARCHAR(100) not null
);
-- Create meetings table
create table meetings (
	id SERIAL primary key,
	title VARCHAR(255) not null,
    description VARCHAR(255) not null,
	host_id INTEGER not null,
	creator_id INTEGER not null,
	start_time TIMESTAMP not null,
	end_time TIMESTAMP not null,
    foreign key (host_id) references users(id),
    foreign key (creator_id) references users(id)
)