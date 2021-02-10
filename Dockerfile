FROM golang:1.13.6

RUN apt update && \
    export PATH=$PATH:/usr/local/go/bin && \
    go get -u -v github.com/gin-gonic/gin && \
    apt install -y sqlite3 gcc-multilib

CMD tail -f /dev/null
EXPOSE 8787