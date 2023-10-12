CREATE TABLE books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    isbn VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    genre VARCHAR(255),
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL
)
;

INSERT INTO books (isbn, title, author, genre, price, quantity)
VALUES
('9780061120084', 'To Kill a Mockingbird', 'Harper Lee', 'Fiction', 12.99, 50),
('9780142407332', 'The Catcher in the Rye', 'J.D. Salinger', 'Fiction', 9.99, 30),
('9781451673319', 'The Great Gatsby', 'F. Scott Fitzgerald', 'Fiction', 11.49, 40),
('9780451524935', '1984', 'George Orwell', 'Science Fiction', 10.99, 25),
('9780062315007', 'The Hobbit', 'J.R.R. Tolkien', 'Fantasy', 15.99, 20),
('9780141187761', 'Pride and Prejudice', 'Jane Austen', 'Romance', 9.95, 35),
('9780060558128', 'The Alchemist', 'Paulo Coelho', 'Fiction', 11.99, 45),
('9780141439587', 'Moby-Dick', 'Herman Melville', 'Adventure', 13.49, 15),
('9780451524935', 'Brave New World', 'Aldous Huxley', 'Science Fiction', 12.99, 28),
('9780060929879', 'The Lord of the Rings', 'J.R.R. Tolkien', 'Fantasy', 18.99, 15)
;
