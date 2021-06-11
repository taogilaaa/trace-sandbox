CREATE TABLE sale_order (
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL,
  payment_method varchar(15) NOT NULL,
  order_date timestamp NOT NULL
);

CREATE TABLE sale_order_product (
  id serial PRIMARY KEY,
  sale_order_id int NOT NULL,
  FOREIGN KEY (sale_order_id) REFERENCES sale_order (id),
  name varchar(255) NOT NULL,
  quantity int NOT NULL
);

INSERT INTO sale_order(email, payment_method, order_date) VALUES
 ('michaelsuyama@northwind.com', 'cash', '2021-06-10'),
 ('nancydavolio@northwind.com', 'cashless', '2021-06-10'),
 ('davidbuchanan@northwind.com', 'cash', '2021-06-10');

INSERT INTO sale_order_product(sale_order_id, name, quantity) VALUES
 (1, 'Chicken', 1),
 (1, 'Pepsi', 3),
 (1, 'Momogi', 10),
 (2, 'Coca Cola', 5),
 (2, 'Mint Candy', 200),
 (3, 'Canned Coffee', 10);