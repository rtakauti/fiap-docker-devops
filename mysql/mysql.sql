create database if not exists go;

create table if not exists tblContato
(
  Id int not null AUTO_INCREMENT,
  Name varchar(50),
  Email varchar(50),
  Phone varchar(50),
  PRIMARY KEY (Id)
);
