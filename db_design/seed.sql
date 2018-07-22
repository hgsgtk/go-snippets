CREATE TABLE books (
    isbn varchar(255) NOT NULL,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    price int NOT NULL
);

INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 1000),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 600),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 700);

ALTER TABLE books ADD PRIMARY KEY (isbn);
