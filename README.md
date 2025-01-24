Wallet API

Это приложение представляет собой REST API для управления кошельками. Оно позволяет выполнять операции пополнения и
снятия средств с кошелька, а также получать текущий баланс.

## Стек технологий

- **Golang** 
- **PostgreSQL** 
- **Docker** 


### Шаги для запуска

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/Egorpalan/wallet-api.git
   cd wallet-api

2. Настройка переменных окружения
   Перед запуском проекта создайте файл `config.env` в корне проекта и заполните его своими значениями:

   ```bash
   SERVER_PORT=8080
   DB_HOST=postgres
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=postgres
   DB_NAME=wallet
   DB_SSLMODE=disable

3. Запустите приложение
    ```bash
   make run

### Примеры запросов

1. Создание кошелька
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallets \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d379"
    }'

2. Пополнение кошелька (DEPOSIT)
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallet \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d379",
    "operationType": "DEPOSIT",
    "amount": 1000
    }'

3. Снятие средств с кошелька (WITHDRAW)
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallet \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d379",
    "operationType": "WITHDRAW",
    "amount": 500
    }'

4. Получение баланса кошелька
    ```bash
   curl -X GET http://localhost:8080/api/v1/wallets/f47ac10b-58cc-4372-a567-0e02b2c3d379

### Для запуска тестов выполните

```bash
   make test