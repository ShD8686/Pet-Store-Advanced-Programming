CREATE TABLE IF NOT EXISTS users (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     username TEXT NOT NULL UNIQUE,
                                     password TEXT NOT NULL,
                                     role TEXT DEFAULT 'user' -- 'admin' или 'user'
);

-- Таблица питомцев (Приют / Продажа)

CREATE TABLE IF NOT EXISTS pets (
                                    id INTEGER PRIMARY KEY AUTOINCREMENT,
                                    chip_number TEXT UNIQUE,
                                    name TEXT,
                                    type TEXT,
                                    gender TEXT,
                                    breed TEXT,
                                    status TEXT,
                                    image_url TEXT,
                                    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS listings (
                                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                                        type TEXT NOT NULL,
                                        pet_type TEXT NOT NULL,
                                        breed TEXT,
                                        photo_url TEXT NOT NULL,
                                        reward REAL DEFAULT 0,
                                        price REAL DEFAULT 0,
                                        has_insurance BOOLEAN DEFAULT 0,
                                        description TEXT,
                                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
                                     id INTEGER PRIMARY KEY AUTOINCREMENT,
                                     email TEXT UNIQUE,
                                     password TEXT,
                                     role TEXT DEFAULT 'user',
                                     created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
