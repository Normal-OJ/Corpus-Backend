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
        - ./BackEnd/unix-clan/lib:/var/lib/unixclan/lib
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
$ cd /app
$ sh start.sh
```
5. done , enjoy :P

## How to restart backend
```bash
$ docker-compose restart backend
$ docker-compose exec backend /bin/bash
$ cd /app
$ sh start.sh
```

## How to remove existed database data
This may need the root permission
```bash
$ sh remove_db.sh
```
After that, restart the backend (see above)

## How to upload data
This can be done in your computer.
```bash
$ python3 corpus_upload.py <source folder>
```

After that, go to cha_store folder, then go to every folders under it, add description.json for each of them.
description.json looks like this:
```json
{
  "provider": "XXX教授",
  "introduction": [
        "情境:玩玩具敘說",
        "語料筆數:24",
        "幼兒年齡:3-6 歲",
        "橫斷式語料"
  ],
  "quoteInfo":"Chang, C. (1998). The development of autonomy in preschool Mandarin\nChinese-speaking children’s play narratives. Narrative Inquiry, 8(1), 77-111."
}
```
