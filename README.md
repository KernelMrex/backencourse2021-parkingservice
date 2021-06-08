# Parking service

## Описание
Сервис отдает информацию о парковках, агрегируя её и дополняя своей.

## Зависимости
* Reservation Service

## Взаимодействие
* gRPC. Подробнее в `api/parkinglistservice/parkinglistservice.proto`

## Управление 
* Запуск `make`
* Остановка `make stop`

## Развертывание
1. Запустить скрипт `bin/generate-grpc.sh`
2. Запустить `make`
