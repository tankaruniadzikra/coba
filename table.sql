-- Colors Table
CREATE TABLE Colors (
    ColorID INT PRIMARY KEY,
    ColorName VARCHAR(50) NOT NULL
);

-- Sizes Table
CREATE TABLE Sizes (
    SizeID INT PRIMARY KEY,
    SizeName VARCHAR(20) NOT NULL,
    Description TEXT
);

-- Categories Table
CREATE TABLE Categories (
    CategoryID INT PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- Brands Table
CREATE TABLE Brands (
    BrandID INT PRIMARY KEY,
    BrandName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- Products Table
CREATE TABLE Products (
    ProductID INT PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    CategoryID INT,
    BrandID INT,
    SizeID INT,
    ColorID INT,
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID),
    FOREIGN KEY (BrandID) REFERENCES Brands(BrandID),
    FOREIGN KEY (SizeID) REFERENCES Sizes(SizeID),
    FOREIGN KEY (ColorID) REFERENCES Colors(ColorID)
);

-- ProductDetails Table (One-to-One relationship with Products)
CREATE TABLE ProductDetails (
    DetailID INT PRIMARY KEY,
    ProductID INT UNIQUE,
    Material VARCHAR(255),
    Weight DECIMAL(8, 2),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- ProductCategories (Many-to-Many relationship with Products)
CREATE TABLE ProductCategories (
    ProductCategoryID INT PRIMARY KEY,
    ProductID INT,
    CategoryID INT,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- Stock Table (To manage stock for each product)
CREATE TABLE Stock (
    StockID INT PRIMARY KEY,
    ProductID INT,
    Quantity INT NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- Users Table (To manage logged-in users)
CREATE TABLE Users (
    UserID INT PRIMARY KEY,
    Email VARCHAR(100) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    FirstName VARCHAR(50),
    LastName VARCHAR(50)
);

-- Orders Table
CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    UserID INT,
    OrderDate DATE NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- OrderItems Table (To track the items in each order)
CREATE TABLE OrderItems (
    OrderItemID INT PRIMARY KEY,
    OrderID INT,
    ProductID INT,
    Quantity INT NOT NULL,
    PricePerUnit DECIMAL(10, 2) NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

--------------------------------------------------------------
-- Insert data into Colors Table
INSERT INTO Colors (ColorName) VALUES 
    ('Red'),
    ('Green'),
    ('Blue');

-- Insert data into Sizes Table
INSERT INTO Sizes (SizeName, Description) VALUES
    ('S', 'Small'),
    ('M', 'Medium'),
    ('L', 'Large');

-- Insert data into Categories Table
INSERT INTO Categories (CategoryName, Description) VALUES
    ('Shirts', 'Various types of shirts'),
    ('Pants', 'Different styles of pants'),
    ('Shoes', 'Footwear of all kinds');

-- Insert data into Brands Table
INSERT INTO Brands (BrandName, Description) VALUES
    ('Brand A', 'Description for Brand A'),
    ('Brand B', 'Description for Brand B'),
    ('Brand C', 'Description for Brand C');

-- Insert data into Products Table
INSERT INTO Products (ProductName, Description, Price, CategoryID, BrandID, SizeID, ColorID) VALUES
    ('Shirt 1', 'Description for Shirt 1', 25.00, 1, 1, 1, 1),
    ('Pants 1', 'Description for Pants 1', 35.00, 2, 2, 2, 2),
    ('Shoes 1', 'Description for Shoes 1', 50.00, 3, 3, 3, 3);

-- Insert data into ProductDetails Table
INSERT INTO ProductDetails (ProductID, Material, Weight) VALUES
    (1, 'Cotton', 0.5),
    (2, 'Denim', 0.7),
    (3, 'Leather', 0.9);

-- Insert data into ProductCategories Table
INSERT INTO ProductCategories (ProductCategoryID, ProductID, CategoryID) VALUES
    (1, 1, 1),
    (2, 2, 1),
    (3, 2, 2),
    (4, 3, 2),
    (5, 3, 3);

-- Insert data into Stock Table
INSERT INTO Stock (ProductID, Quantity) VALUES
    (1, 100),
    (2, 150),
    (3, 200);

-- Insert data into Users Table
INSERT INTO Users (Email, Password, FirstName, LastName) VALUES
    ('john@doe.com', 'john', 'John', 'Doe'),
    ('jane@doe.com', 'jane', 'Jane', 'Doe'),
    ('bob@smith.com', 'bob', 'Bob', 'Smith');

-- Insert data into Orders Table
INSERT INTO Orders (UserID, OrderDate, TotalAmount) VALUES
    (1, '2023-10-15', 150.00),
    (2, '2023-10-16', 200.00),
    (3, '2023-10-17', 250.00);

-- Insert data into OrderItems Table
INSERT INTO OrderItems (OrderID, ProductID, Quantity, PricePerUnit, TotalPrice) VALUES
    (1, 1, 2, 25.00, 50.00),
    (1, 2, 1, 35.00, 35.00),
    (2, 3, 3, 50.00, 150.00);
