CREATE TABLE users (
    id int primary key AUTO_INCREMENT,
    name varchar(255) not null,
    phone varchar (255) not null ,
    password varchar (255) not null ,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

)