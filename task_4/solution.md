# **Отчет по выполнению ДЗ-4 "SQL и базы данных"**

### **Задача 1. "Долгожитель"**

**Запрос: (простой вариант, если нам нужен только один старожитель)**
```sql
SELECT name, age 
  FROM users 
 WHERE city='Москва' 
 ORDER BY age DESC 
 LIMIT 1;
```

**Вывод:**
```sql
 name  | age 
-------+-----
 Борис |  40
(1 row)
```

**Пояснения:** в данном случае, postgreSQL не гарантирует, что всегда будет выводиться *Борис*. Результат запроса *Борис* или *Вера* - чистая случайность в данном случае, потому что в каких-то ситуациях может выпасть *Вера*, а в каких-то *Борис*. 
  
Поэтому более правильным решением будет следующий запрос:

**Запрос: (более сложный вариант с помощью подзапроса, но более точный)**
```sql
SELECT name, age
  FROM users
 WHERE age = (SELECT MAX(age) FROM users)
   AND city = 'Москва';
```

**Вывод:**
```sql
 name  | age 
-------+-----
 Борис |  40
 Вера  |  40
(2 rows)
```

**Пояснения:** в данном случае, мы получаем **всех** старожителей, у которых совпадает возраст, и не теряем тех, у кого может совпадать возраст, но postgreSQL их не вывел.

### **Задача 2. "Перепись населения"**

**Запрос:**
```sql
SELECT city, COUNT(name) as count_people 
  FROM users 
 GROUP BY city 
HAVING COUNT(name)> 1 
 ORDER BY count_people DESC, city ASC;
```

**Вывод:**
```sql
  city  | count_people 
--------+--------------
 Москва |            3
 Казань |            2
(2 rows)
```

### **Задача 3. "Учёт трат"**

**Запрос:**
```sql
SELECT u.name, COALESCE(SUM(o.price), 0) as price_order
  FROM users as u 
  LEFT JOIN orders as o ON u.id = o.user_id 
 GROUP BY u.id 
 ORDER BY price_order DESC;
```

**Вывод:**
```sql
 name  | price_order 
-------+-------------
 Борис |      120000
 Глеб  |       57000
 Анна  |       51500
 Егор  |       38000
 Вера  |           0
 Дина  |           0
(6 rows)
```

### **Задача 4. "Любители музыки"**

**Запрос 1: C JOIN**
```sql
SELECT DISTINCT u.name
  FROM users as u
  JOIN orders as o ON u.id = o.user_id
 WHERE item = 'наушники'
 ORDER BY name; 
```

**Запрос 2: С подзапросом**
```sql
SELECT DISTINCT name
  FROM users
 WHERE id IN 
    (
      SELECT user_id 
        FROM orders 
       WHERE item = 'наушники'
    ) 
 ORDER BY name;
```

**Вывод:**
```sql
 name 
------
 Глеб
 Егор
(2 rows)
```

### **Задача 5. "Бюджетный выбор"**

**Запрос:**
```sql
SELECT item, price
 FROM orders
WHERE price = 
    (
      SELECT MIN(price) 
        FROM orders 
      WHERE price > 10_000
    );
```

**Вывод:**
```sql
   item   | price 
----------+-------
 наушники | 12000
(1 row)
```
  