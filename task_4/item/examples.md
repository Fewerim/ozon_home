# Примеры использования

#### Сохранить item
curl -X POST localhost:50052/v1/item -H 'Content-Type: application/json' -d '{"title": "Book", "description": "Very Great Book"}'

#### Получить item по id
curl localhost:50052/v1/item/1

#### Ошибка: несуществующий id
curl localhost:50052/v1/item/999

#### Ошибка: id <= 0
curl localhost:50052/v1/item/0

#### Список методов
grpcurl -plaintext localhost:50051 list

#### Сохранить item
grpcurl -plaintext -d '{"title": "Book", "description": "Very Great Book"}' localhost:50051 item.Item/PostItem

#### Получить item по id
grpcurl -plaintext -d '{"id": 1}' localhost:50051 item.Item/GetItem