# Этапы работы.
* [x] Упростить стек.
* [x] Добавить базу данных.
* [x] Добавить поддержку sha256.
* [ ] #7 Убрать системные метрики Prometheus.
* [x] Залечить баги.
    * [x] #1 Не происходит отключения при некорректной авторизации.
    * [x] #2 Гонка при отключении пула если одновременно с отключением приходит команда.
    * [x] #3 Ошибка декодирования bool при mining.authorize.
    * [x] #4 Invalid miner subscribe response.
    * [x] #5 Не находится уже зарегистрированный пользователь.
    * [x] #6 При синхронизации отдавать правильный Extranonce2Size.

# Команды тестирования mining.configure.
Майнер без поддержки `mining.configure`.
```bash
echo -e '{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-multi/1.3.6"]}\n{"id": 2, "method": "mining.authorize", "params": ["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]}\n' | nc 127.0.0.1 9332
```

Майнер без поддержки `mining.configure`, но с поддержкой `mining.extranonce.subscribe`.
```bash
echo -e '{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-multi/1.3.6"]}\n{"id": 2, "method": "mining.authorize", "params": ["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]}\n{"id": 3, "method": "mining.extranonce.subscribe", "params": []}\n' | nc 127.0.0.1 9332
```

Майнер с поддержкой только расширения `subscribe-extranonce` в `mining.configure`.
```bash
echo -e '{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-multi/1.3.6"]}\n{"id": 2, "method": "mining.authorize", "params": ["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]}\n{"id": 3, "method": "mining.configure", "params": [["subscribe-extranonce"],{}]}\n' | nc 127.0.0.1 9332
```

Майнер с поддержкой только расширения `version-rolling` в `mining.configure`.
```bash
echo -e '{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-multi/1.3.6"]}\n{"id": 2, "method": "mining.authorize", "params": ["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]}\n{"id": 3, "method": "mining.configure", "params": [["version-rolling"],{"version-rolling.mask":"ffffffff", "version-rolling.min-bit-count":2}]}\n' | nc 127.0.0.1 9332
```

Майнер с поддержкой обоих расширений в `mining.configure`.
```bash
echo -e '{"id": 1, "method": "mining.subscribe", "params": ["cpuminer-multi/1.3.6"]}\n{"id": 2, "method": "mining.authorize", "params": ["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]}\n{"id": 3, "method": "mining.configure", "params": [["subscribe-extranonce", "version-rolling"],{"version-rolling.mask":"ffffffff", "version-rolling.min-bit-count":2}]}\n' | nc 127.0.0.1 9332
```
Майнер, выполняющий первой командой `mining.configure`.
```bash
echo -e '{"id": 1, "method": "mining.configure", "params": [["version-rolling"],{"version-rolling.mask":"ffffffff", "version-rolling.min-bit-count":2}]}\n' | nc 127.0.0.1 9332
```
