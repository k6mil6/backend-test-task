Для проверки потребуется создать файл dev.env

```
HTTP_ADDRESS=:8080
DB_ADDRESS=host=db user=postgres password=postgres dbname=db port=5432 sslmode=disable
ENV=dev
GRPC_PORT=50051
CLIENT_ADDRESS=grpc:50051
CLIENT_TIMEOUT=5m
CLIENT_RETRIES=3
```

Запускается в докере:
docker-compose up

Пример:

```
curl --location 'http://localhost:8080/klines?symbol=BTCUSDT&interval=30m&limit=3'
```

Ответ

```
[
    {
        "openTime": 1710883800000,
        "open_price": 63779.99,
        "high_price": 64109.98,
        "low_price": 63581.1,
        "close_price": 63952.01,
        "volume": 487.18344,
        "kline_close_time": 1710885599999,
        "quote_asset_volume": 31119439.9411381,
        "number_of_trades": 34283,
        "taker_buy_base_asset_volume": 263.87625,
        "taker_buy_quote_asset_volume": 16858199.0162209
    },
    {
        "openTime": 1710885600000,
        "open_price": 63952.01,
        "high_price": 64036.5,
        "low_price": 63008.53,
        "close_price": 63013.19,
        "volume": 1479.41806,
        "kline_close_time": 1710887399999,
        "quote_asset_volume": 93814397.1364283,
        "number_of_trades": 60376,
        "taker_buy_base_asset_volume": 629.29065,
        "taker_buy_quote_asset_volume": 39895289.0903326
    },
    {
        "openTime": 1710887400000,
        "open_price": 63013.2,
        "high_price": 63277.2,
        "low_price": 62788.89,
        "close_price": 62923.53,
        "volume": 708.01255,
        "kline_close_time": 1710889199999,
        "quote_asset_volume": 44600193.087898,
        "number_of_trades": 28547,
        "taker_buy_base_asset_volume": 302.50022,
        "taker_buy_quote_asset_volume": 19057474.6720919
    }
]
```