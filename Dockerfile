FROM golang:1.17.3

RUN apt update && \
    export PATH=$PATH:/usr/local/go/bin && \
    apt install -y ca-certificates libgnutls30 
    #RUN apt install sqlite3 build-essential && \
    RUN go get -u -v github.com/gin-gonic/gin
    RUN curl --silent --location https://deb.nodesource.com/setup_16.x | bash -
    RUN apt update
    RUN apt-get install --yes nodejs
    RUN apt-get install libc6-i386
    #RUN dpkg --add-architecture i386 && \
    #apt-get update
    RUN apt-get install -y gcc-multilib
    #RUN apt-get install libstdc++6:i386 libgcc1:i386 libcurl4-gnutls-dev:i386
    #RUN apt-get install libstdc++6:i386
    RUN apt-get update
    RUN apt-get install libstdc++6

CMD tail -f /dev/null
EXPOSE 8787