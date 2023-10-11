-- 1. Users Table:
CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Email VARCHAR(100) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    FirstName VARCHAR(50),
    LastName VARCHAR(50)
);

-- 2. Orders Table:
CREATE TABLE Orders (
    OrderID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT,
    OrderDate DATE NOT NULL,
    TotalAmount DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

-- 3. Products Table:
CREATE TABLE Products (
    ProductID INT AUTO_INCREMENT PRIMARY KEY,
    ProductName VARCHAR(255) NOT NULL,
    Description TEXT,
    Price DECIMAL(10, 2) NOT NULL,
    Material VARCHAR(255),
    Weight DECIMAL(8, 2),
    BrandID INT,
    SizeID INT,
    ColorID INT,
    FOREIGN KEY (BrandID) REFERENCES Brands(BrandID),
    FOREIGN KEY (SizeID) REFERENCES Sizes(SizeID),
    FOREIGN KEY (ColorID) REFERENCES Colors(ColorID)
);

-- 4. OrderItems Table:
CREATE TABLE OrderItems (
    OrderItemID INT AUTO_INCREMENT PRIMARY KEY,
    OrderID INT,
    ProductID INT,
    Quantity INT NOT NULL,
    PricePerUnit DECIMAL(10, 2) NOT NULL,
    TotalPrice DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

-- 5. Categories Table:
CREATE TABLE Categories (
    CategoryID INT AUTO_INCREMENT PRIMARY KEY,
    CategoryName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- 6. ProductCategories Table (Junction Table):
CREATE TABLE ProductCategories (
    ProductCategoryID INT AUTO_INCREMENT PRIMARY KEY,
    ProductID INT,
    CategoryID INT,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

-- 7. Brands Table:
CREATE TABLE Brands (
    BrandID INT AUTO_INCREMENT PRIMARY KEY,
    BrandName VARCHAR(255) NOT NULL,
    Description TEXT
);

-- 8. Sizes Table:
CREATE TABLE Sizes (
    SizeID INT AUTO_INCREMENT PRIMARY KEY,
    SizeName VARCHAR(20) NOT NULL,
    Description TEXT
);

-- 9. Colors Table:
CREATE TABLE Colors (
    ColorID INT AUTO_INCREMENT PRIMARY KEY,
    ColorName VARCHAR(50) NOT NULL
);

-- 10. Stock Table:
CREATE TABLE Stock (
    StockID INT AUTO_INCREMENT PRIMARY KEY,
    ProductID INT,
    Quantity INT NOT NULL,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

--------------------------------------------------------------
-- Insert sample brands for clothing
INSERT INTO Brands (BrandID, BrandName, Description)
VALUES
    (1, 'Adidas', 'Sportswear and athletic shoes'),
    (2, 'Zara', 'Fashion and apparel retailer'),
    (3, 'H&M', 'Global clothing brand');

-- Insert sample sizes for clothing
INSERT INTO Sizes (SizeID, SizeName, Description)
VALUES
    (1, 'Small', 'Small size for clothing'),
    (2, 'Medium', 'Medium size for clothing'),
    (3, 'Large', 'Large size for clothing');

-- Insert sample colors for clothing
INSERT INTO Colors (ColorID, ColorName)
VALUES
    (1, 'Red'),
    (2, 'Blue'),
    (3, 'Black');

-- Insert sample categories for clothing
INSERT INTO Categories (CategoryID, CategoryName, Description)
VALUES
    (1, 'T-Shirts', 'Casual and comfortable'),
    (2, 'Jeans', 'Denim pants'),
    (3, 'Dresses', 'Various dress styles');

-- Insert sample clothing products
INSERT INTO Products (ProductID, ProductName, Description, Price, Material, Weight, BrandID, SizeID, ColorID)
VALUES
    (1, 'Classic T-Shirt', 'A comfortable and classic t-shirt', 20.00, 'Cotton', 0.3, 1, 2, 1),
    (2, 'Slim Fit Jeans', 'Fitted jeans for a modern look', 45.00, 'Denim', 0.7, 2, 1, 2),
    (3, 'Floral Sundress', 'A floral summer dress', 35.00, 'Polyester', 0.4, 3, 3, 1);

-- Associate clothing products with categories
INSERT INTO ProductCategories (ProductCategoryID, ProductID, CategoryID)
VALUES
    (1, 1, 1),
    (2, 2, 2),
    (3, 3, 3);

-- Insert sample stock information for clothing
INSERT INTO Stock (StockID, ProductID, Quantity)
VALUES
    (1, 1, 50),
    (2, 2, 30),
    (3, 3, 20);
