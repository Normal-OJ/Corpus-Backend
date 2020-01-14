FROM ubuntu
RUN apt update && \
    apt -y install g++ && \
    apt -y install golang && \
    apt -y install nano && \
    apt -y install python3 && \
    apt -y install git && \
    apt -y install sudo && \
    go get -u github.com/gin-gonic/gin
 
CMD tail -f /dev/null
EXPOSE 8787

