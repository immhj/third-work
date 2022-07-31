create database infomation;
use infomation;
create table warehouse(
  num varchar(20) primary key,
  space int
);
create table cloth(
  num varchar(20) primary key,
  price int,
  size varchar(20),
  tp varchar(20)
);
create table supplier(
  num varchar(20) primary key,
  name varchar(20)
);
create table supplier_env(
  snum varchar(20),
  cnum int,
  qulify varchar(20)
);
alter table supplier_env add primary key(cnum ,snum);
