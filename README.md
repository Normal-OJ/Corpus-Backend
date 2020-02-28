# How To Use This Sh*t
1. download it
2. create a docker-compose file
```yaml
# docker compose file for lazy bones  :P
version: '3'

services:
   backend:
      build: ./BackEnd
      volumes:
        - ./BackEnd:/app
        - ./somefolder:/cha_store # some folder to store cha files
      env_file:
      - ./BackEnd/.env
      ports:
      - 8787:8787

```
3. run the following command
 ```bash
$ docker-compose up -d # and wait for a year :P
$ docker-compose exec backend /bin/bash
$ cd app
$ sh run.sh
 ```
5. done , enjoy :P