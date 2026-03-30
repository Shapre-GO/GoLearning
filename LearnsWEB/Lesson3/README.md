# 🚀 НАЧАЛО ПРОЕКТА
---
### **📦 Установка базы данных:** 
```bash
sudo apt install postgresql postgresql-contrib    
```     
|
### **⚙️ Запуск службы:**
```bash
sudo systemctl start postgresql
```
> ⚡ Автоматический запуск службы ⬇️  
```bash
sudo systemctl enable postgresql
```  
|
### **🗄️ Создание базы данных и пользователя:**
> 🔑 Выполняем вход ⬇️
```bash
sudo -u postgres psql
```

> 🎲 Создаём БД "authdb" ⬇️
```bash
CREATE DATABASE authdb;
```

> 👤 Создаём пользователя "appuser" с паролем "secret" ⬇️
```bash
CREATE USER appuser WITH PASSWORD 'secret';
```

> ✅ Добавляем пользователя "appuser" с админкой в БД "authdb" ⬇️
```bash
GRANT ALL PRIVILEGES ON DATABASE authdb TO appuser;
```
|
### **🗄️ Создание таблицы users:**
```bash
-- Подключение к базе authdb
\c authdb

-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```
> В данном случае вынес в отдельный ".sql" файл и запустил с помощью команды ⬇️
```bash
sudo -u postgres psql -f migrations.sql
```
> ✅ Получаем необходимый драйвер для PostgreSql в ZED ⬇️
```bash
go get github.com/jackc/pgx/v5
```
