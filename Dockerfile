FROM golang:1.13.6

RUN apt update && \
    export PATH=$PATH:/usr/local/go/bin && \
    go get -u -v github.com/gin-gonic/gin && \
    apt install sqlite3

CMD tail -f /dev/null
EXPOSE 8787