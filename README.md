*p.s.: please edit src/route/main.go to rewrite the path of unix clang first , since i am too lazy to change my code*
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
      env_file: 
      - ./BackEnd/.env

```
3. run the following command
 ```bash
$ docker-compose up -d # and wait for a year :P
$ docker-compose exec backend /bin/bash
$ cd app
$ sh run.sh
 ```
5. done , enjoy :P