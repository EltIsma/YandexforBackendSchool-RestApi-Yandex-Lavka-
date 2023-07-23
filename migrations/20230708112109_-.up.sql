CREATE TABLE couriers (
    id serial PRIMARY KEY,
    type varchar(50) NOT NULL,
    regions integer[] NOT NULL,
    working_hours text[] NOT NULL
);

CREATE TABLE "orders"(
    id serial PRIMARY KEY,
    courier_id integer REFERENCES couriers,
    weight float NOT NULL,
    region integer NOT NULL,
    delivery_hours varchar(255),
    complete_time timestamp,
    cost integer NOT NULL
);