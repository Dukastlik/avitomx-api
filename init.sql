CREATE TABLE products(
id serial primary key, 
merchant_id int not null,
offer_id int not null,
name varchar not null,
price int,
quantity int not null,
avaliable boolean);