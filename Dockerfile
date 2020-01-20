FROM ubuntu

# install misc and never use vi
RUN apt update && \
    apt -y install g++ && \
    apt -y install nano && \
    apt -y install python3 && \
    apt -y install git && \
    apt -y install sudo

# install go v 1.13
RUN apt install wget && \
    wget https://dl.google.com/go/go1.13.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.13.6.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin && \
    go get -u -v github.com/gin-gonic/gin

CMD tail -f /dev/null
EXPOSE 8787