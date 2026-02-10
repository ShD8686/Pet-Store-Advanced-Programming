package sql

-- Таблица товаров (Marketplace: корма, игрушки, питомцы)
CREATE TABLE IF NOT EXISTS products (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
category VARCHAR(100), -- food, toys, pharmacy, pets
price DECIMAL(10, 2) NOT NULL,
stock INT NOT NULL DEFAULT 0,
description TEXT
);

-- Таблица записей (Vet/Nanny services)
CREATE TABLE IF NOT EXISTS appointments (
id SERIAL PRIMARY KEY,
service_type VARCHAR(100), -- vaccination, grooming, sitting
pet_name VARCHAR(100),
appointment_date TIMESTAMP NOT NULL,
status VARCHAR(50) DEFAULT 'pending'
);

-- Наполним начальными данными для теста
INSERT INTO products (name, category, price, stock, description) VALUES
('Premium Dog Food', 'food', 5500.00, 20, 'High protein food for adult dogs'),
('Cat Laser Pointer', 'toys', 1200.50, 15, 'Fun interactive toy for cats'),
('Winter Pet Jacket', 'accessories', 8000.00, 5, 'Warm jacket for small dogs');
