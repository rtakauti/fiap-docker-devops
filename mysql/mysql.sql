create database go;

create table tblContato if not exists
(
  Id int not null AUTO_INCREMENT,
  Name varchar(50),
  Email varchar(50),
  Phone varchar(50)
);
