-- Обновить существующего пользователя до администратора
UPDATE users 
SET role = 'admin' 
WHERE email = 'your_email@example.com';

-- ИЛИ создать нового пользователя-администратора
-- Замените 'your_hashed_password' на реальный хешированный пароль
INSERT INTO users (email, password, name, role, created_at, updated_at) 
VALUES (
    'admin@example.com', 
    'your_hashed_password', -- Замените на реальный хешированный пароль
    'Администратор',
    'admin',
    NOW(),
    NOW()
);
