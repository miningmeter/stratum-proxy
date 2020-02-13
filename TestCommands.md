# Subscribe
## Invalid mining.subscribe.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Already subscribed worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Subscribe worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```

# Authorize.
## Unsubscribed worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Already authorized worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.configure\", \"params\": [[\"subscribe-extranonce\"], {}]}\n{\"id\":3, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n{\"id\":4, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Authorize worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```

# Submit.
## Ignore share from unsubscribed worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.submit\", \"params\":[\"miner.miner\", \"6e1f\", \"9a279248\", \"5d317e62\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Ignore share from unauthorized worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.submit\", \"params\":[\"miner.miner\", \"6e1f\", \"9a279248\", \"5d317e62\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Job not found.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n{\"id\":3, \"method\":\"mining.submit\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"6e1f\", \"9a279248\", \"5d317e62\", \"5d317e62\"]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```

# ExtranonceSubscribe
## Extranonce.subscribe unsubscribed worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.extranonce.subscribe\", \"params\":[]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Extranonce.subscribe unauthorized worker.
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.extranonce.subscribe\", \"params\":[]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Extranonce.subscribe
```bash
echo -ne "{\"id\":1, \"method\":\"mining.subscribe\", \"params\":[\"cpuminer/2.5.0\"]}\n{\"id\":2, \"method\":\"mining.authorize\", \"params\":[\"1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG\", \"X\"]}\n{\"id\":3, \"method\":\"mining.extranonce.subscribe\", \"params\":[]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```

# Configure
## Invalid mining.configure
```bash
echo -ne "{\"id\":1, \"method\":\"mining.configure\", \"params\":[]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
## Mining.configure
```bash
echo -ne "{\"id\":1, \"method\":\"mining.configure\", \"params\": [[\"version-rolling\"], {\"version-rolling.mask\": \"1fffe000\", \"version-rolling.min-bit-count\": 2}]}\n" | while IFS= read -r line; do echo $line; sleep .5; done | nc 127.0.0.1 9332
```
