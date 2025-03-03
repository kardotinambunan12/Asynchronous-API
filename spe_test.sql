show tables ;

RENAME TABLE addresses TO customer;
create database spe_test;
use spe_test;


show databases;

CREATE TABLE customer
(
    customer_id          VARCHAR(100) NOT NULL,
    customer_pan  VARCHAR(100) NOT NULL,
    customer_name      VARCHAR(200),
    tgl_rekam        VARCHAR(100),
    petugas_rekam    VARCHAR(100),
    PRIMARY KEY (customer_id)
) ENGINE InnoDB;

select * from customer;


select * from merchant;

CREATE TABLE merchant
(
    merchant_id          VARCHAR(36) NOT NULL,
    merchant_name      VARCHAR(200),
    merchant_city  VARCHAR(100) NOT NULL,
    tgl_rekam        VARCHAR(50),
    petugas_rekam    VARCHAR(100),
    PRIMARY KEY (merchant_id)
) ENGINE InnoDB;




CREATE TABLE transaction
(
    request_id          VARCHAR(32) NOT NULL,
    customer_pan      VARCHAR(200),
    amount  VARCHAR(100) NOT NULL,
    transaction_datetime        VARCHAR(50),
    rrn    VARCHAR(100),
    bill_number    VARCHAR(100),
    customer_name    VARCHAR(100),
    merchant_id    VARCHAR(100),
    merchant_name    VARCHAR(100),
    merchant_city    VARCHAR(100),
    currency_code    VARCHAR(100),
    payment_status    VARCHAR(100),
    payment_description    VARCHAR(100),
    tgl_rekam        VARCHAR(50),
    petugas_rekam    VARCHAR(100),
    PRIMARY KEY (request_id),
        FOREIGN KEY fk_merchant (merchant_id) REFERENCES merchant (merchant_id),
        FOREIGN KEY fk_customer (customer_pan) REFERENCES customer (customer_pan)
) ENGINE InnoDB;

# RENAME TABLE transaction TO transactions;

select * from transactions;

delete from transactions where request_id='aefa7b68dc02443681a2123da02405b3';


SELECT bill_number
FROM spe_test.transactions
ORDER BY bill_number DESC
LIMIT 1;

DESC transactions;

delete from transactions where request_id='468eab07695b4d0f9096e1f86307814b';

# INSERT INTO spe_test.transactions
# (request_id, customer_pan, amount, transaction_datetime, rrn, bill_number, customer_name,
#  merchant_id, merchant_name, merchant_city, currency_code, payment_status, payment_description)
# VALUES ('aefa7b68dc02443681a2123da02405b4', '8327732737474787324', '150000.75',
#         '2024-03-01T14:30:00Z', '987654321012', 'INV123456', 'John Doe', 'cc655cf4-f98b-44a6-be4c-bd258c087551', 'Toko Elektronik ABC',
#         'Jakarta', 'IDR', 'Completed', 'Pembayaran berhasil');


# ALTER TABLE customer DROP PRIMARY KEY;
#
# ALTER TABLE customer ADD PRIMARY KEY (customer_pan);

select request_id, customer_pan, amount,
					   transaction_datetime, rrn, bill_number,
					   customer_name, merchant_id, merchant_name,
					   merchant_city, currency_code, payment_status,
					   payment_description, tgl_rekam, petugas_rekam
				from spe_test.transactions where request_id = 'a76b70b0ceb043e583330f509f6236e6'
				and bill_number = 'INV123456';






DESCRIBE merchant;
DESCRIBE customer;


# drop table merchant;
