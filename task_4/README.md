# Руководство по запуску БД для проверки задания

### **Команды**

**Подготовка окружения:**
```bash
cp .env.example .env
```

**Запуск БД:**
```bash
docker run --name hw4 -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=hw4 -p 5433:5432 -d postgres:18
```

Если контейнер `hw4` уже создан без базы `hw4`, создайте её отдельно:
```bash
docker exec hw4 psql -U postgres -d postgres -c 'CREATE DATABASE hw4;'
```

**Запуск миграций:**
```bash
make migrate-up
make migrate-status
```

**PSQL:**
```bash
docker exec -it hw4 psql -U postgres -d hw4
```

### **Проверка**

**Тестовый запрос:**
```sql
SELECT * 
  FROM users;
```

**Ожидаемый вывод:**
```sql
 id | name  | age |  city  
----+-------+-----+--------
  1 | Анна  |  25 | Москва
  2 | Борис |  40 | Москва
  3 | Вера  |  40 | Москва
  4 | Глеб  |  31 | Казань
  5 | Дина  |  19 | Казань
  6 | Егор  |  27 | СПб
(6 rows)
```


