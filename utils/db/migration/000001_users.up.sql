create table Users(
	ID int not null auto_increment primary key,
	Username varchar(255) not null,
	Fullname varchar(255),
	Email varchar(255),
	City varchar(255),
	Phonenumber varchar(20),
	Password varchar(255)not null,
	Points int,
	Balance int
);

create table Brands(
	ID int not null auto_increment primary key,
	Name varchar(255)
);

create table Vouchers(
	ID int not null auto_increment primary key,
	Name varchar(255),
	BrandID int,
	CiM int,
	CiP int,
	Code varchar(255),
	foreign key (BrandID) references Brands(ID)
);

create table Transactions(
	ID int not null auto_increment primary key,
	VoucherID int,
	UserID int,
	Status varchar(255),
	foreign key (VoucherID) references Vouchers(ID),
	foreign key (UserID) references Users(ID)
);