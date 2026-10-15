# Корзина и каталог товаров

Проект состоит из двух сервисов:

- **product-service** — gRPC-сервис на порту `50051`. При запуске загружает товары из `product-service/data.json` в память. Предоставляет методы `GetProduct` (поиск по SKU) и `ListProducts` (список товаров по страницам, отсортированный по SKU). Для списка значения по умолчанию: `limit=20`, `offset=0`.
- **cart** — HTTP-сервис на порту `8080`. Хранит корзины в памяти, проверяет наличие товара через gRPC перед добавлением, суммирует количество повторно добавленного товара и возвращает корзину с общей стоимостью и сортировкой по SKU. HTTP-запросы логируются middleware.

Данные корзин теряются после перезапуска. Swagger UI не реализован.

## Запуск

Из каталога `task_2`:

```bash
make run-all
```

Для повторной генерации Go-кода из `.proto` нужен `protoc` и плагины `protoc-gen-go`, `protoc-gen-go-grpc`:

```bash
make generate
```

## Примеры запросов

```bash
# Получить товар
grpcurl -plaintext -d '{"sku":1076963}' localhost:50051 productService.ProductService/GetProduct

# Получить первую страницу каталога; следующая — offset=20
grpcurl -plaintext -d '{"offset":0,"limit":20}' localhost:50051 productService.ProductService/ListProducts

# Добавить две единицы товара в корзину пользователя 1007
curl -i -X POST localhost:8080/user/1007/cart/1076963 \
  -H 'Content-Type: application/json' -d '{"count":2}'

# Получить корзину
curl -i localhost:8080/user/1007/cart
```

Некорректные параметры дают HTTP `400` (для gRPC — `InvalidArgument`), отсутствующая или пустая корзина — `404`, а добавление несуществующего товара — `412`. HTTP-ошибки возвращаются текстом.
