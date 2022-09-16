create table Users(
	ID int not null auto_increment primary key,
	Username varchar(255) not null,
	Fullname varchar(255),
	Email varchar(255),
	City varchar(255),
	Phonenumber varchar(20),
	Password varchar(255)not null,
	Points int,
	Balance int,
	Role varchar(20) not null,
);

create table Brands(
	ID int not null auto_increment primary key,
	Name varchar(255)
);

create table Vouchers(
	ID int not null auto_increment primary key,
	Name varchar(255),
	Brand_ID int,
	CIM int,
	CIP int,
	foreign key (Brand_ID) references Brands(ID)
);

create table Transactions(
	ID int not null auto_increment primary key,
	Voucher_ID int,
	User_ID int,
	Items int,
	Status varchar(255),
	Code varchar(255),
	Total int,
	foreign key (Voucher_ID) references Vouchers(ID),
	foreign key (User_ID) references Users(ID)
);