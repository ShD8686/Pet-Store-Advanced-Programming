-- Таблица питомцев (Приют / Продажа)
CREATE TABLE IF NOT EXISTS pets (
                                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                                    name TEXT NOT NULL,
                                    category TEXT, -- Собака, кошка, попугай
                                    price REAL DEFAULT 0,
                                    status TEXT DEFAULT 'available', -- 'available', 'adopted', 'sold'
                                    description TEXT,
                                    is_for_adoption INTEGER DEFAULT 0 -- 1 если в приюте, 0 если на продажу
);

-- Таблица товаров (Корма, игрушки)
CREATE TABLE IF NOT EXISTS products (
                                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        name TEXT NOT NULL,
                                        category TEXT, -- food, toys, pharmacy
                                        price REAL NOT NULL,
                                        stock INTEGER NOT NULL DEFAULT 0,
                                        description TEXT
);

-- Таблица записей (Ветеринар / Няня)
CREATE TABLE IF NOT EXISTS appointments (
                                            id INTEGER PRIMARY KEY AUTOINCREMENT,
                                            service_type TEXT, -- vaccination, grooming, pet-sitting (няня)
                                            pet_name TEXT,
                                            owner_name TEXT,
                                            appointment_date TEXT NOT NULL, -- В SQLite даты храним как текст или числа
                                            status TEXT DEFAULT 'pending'
);

CREATE TABLE IF NOT EXISTS orders (
                                      id INTEGER PRIMARY KEY AUTOINCREMENT,
                                      pet_id INTEGER,
                                      user_id INTEGER,
                                      total_price REAL
);

-- Наполним начальными данными для теста
-- Пример для одного товара
INSERT OR IGNORE INTO products (id, name, category, price, stock, description) VALUES
    (1, 'Premium Dog Food', 'food', 5500.0, 20, 'High protein food for adult dogs'),
    ('Cat Laser Pointer', 'toys', 1200.5, 15, 'Fun interactive toy for cats');

INSERT INTO pets (name, category, price, status, is_for_adoption) VALUES
                                                                      ('Buddy', 'Dog', 0, 'available', 1), -- Собака в приюте
                                                                      ('Misty', 'Cat', 15000, 'available', 0); -- Кошка на продажу