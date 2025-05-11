# Пет-проект: gRPC, React, Golang, Redis, Docker

Этот проект демонстрирует взаимодействие фронтенда на **React** и бэкенда на **Golang** через **gRPC**, с использованием **Redis** для кэширования и **Docker** для развертывания. Проект находится в этом репозитории, но в ветке master.

## Технологии
- **Frontend**: React (TypeScript)
- **Backend**: Golang (gRPC)
- **Кэширование**: Redis
- **Развертывание**: Docker, Docker Compose

## Архитектура
- Связь между фронтендом и бэкендом осуществляется через **gRPC** с генерацией `proto`-файлов.
- Бэкенд разворачивается в Docker (отдельные образы для сервера, API и Redis).
- Для корректной работы настроен `docker-compose.yml`.

## Запуск проекта

### Предварительные требования
- Установленные [Docker](https://docs.docker.com/get-docker/) и [Docker Compose](https://docs.docker.com/compose/install/).
- Node.js (для фронтенда).

### Инструкция по запуску
1. **Запустите бэкенд и Redis**:
   ```bash
   docker-compose up -d
   npm install
   npm start