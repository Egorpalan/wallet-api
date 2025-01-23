Wallet API

Это приложение представляет собой REST API для управления кошельками. Оно позволяет выполнять операции пополнения и снятия средств с кошелька, а также получать текущий баланс.

## Стек технологий

- **Golang** — язык программирования для backend-разработки.
- **PostgreSQL** — база данных для хранения информации о кошельках.
- **Docker** — контейнеризация приложения и базы данных.
- **Docker Compose** — оркестрация контейнеров.

### Шаги для запуска

1. **Клонируйте репозиторий:**

   ```bash
   git clone https://github.com/Egorpalan/wallet-api.git
   cd wallet-api
   
2. Запустите приложение
    ```bash
   make build && make run


### Примеры запросов

1. Создание кошелька
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallets \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d479"
    }'

2. Пополнение кошелька (DEPOSIT)
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallet \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "operationType": "DEPOSIT",
    "amount": 1000
    }'
   
3. Снятие средств с кошелька (WITHDRAW)
    ```bash
   curl -X POST http://localhost:8080/api/v1/wallet \
    -H "Content-Type: application/json" \
    -d '{
    "walletId": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
    "operationType": "WITHDRAW",
    "amount": 500
    }'

4. Получение баланса кошелька
    ```bash
   curl -X GET http://localhost:8080/api/v1/wallets/f47ac10b-58cc-4372-a567-0e02b2c3d479
   

### Для запуска тестов выполните
```bash
   make test