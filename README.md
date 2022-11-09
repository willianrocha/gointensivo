# gointensivo


Create DB
```
$ sqlite3 orders.db
sqlite> CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));
```