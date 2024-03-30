FROM ubuntu:jammy

RUN apt-get update && \
    apt-get install -y wget

# setup go 1.22
# see: https://github.com/docker-library/golang/blob/d5ba02dca99c1a2d221c4a20e45eedc0f0380f31/1.22/bookworm/Dockerfile
RUN set -eux; \
    sha256="aab8e15785c997ae20f9c88422ee35d962c4562212bb0f879d052a35c8307c7f"; \
    wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz -O go.tgz; \
    echo "$sha256 *go.tgz" | sha256sum -c -; \
    tar -C /usr/local -xzf go.tgz; \
    rm go.tgz;
ENV PATH /usr/local/go/bin:$PATH

# setup node 16
RUN curl --silent --location https://deb.nodesource.com/setup_16.x | bash - && \
    apt-get update && \
    apt-get install -y nodejs

RUN apt-get install -y ca-certificates libgnutls30 libc6-i386 gcc-multilib libstdc++6

RUN apt-get install -y python3.9 python3-pip


EXPOSE 8787
WORKDIR /app 

COPY requirements.txt ./
RUN pip install -r requirements.txt

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src
COPY main.go ./

RUN go build -o main main.go
COPY . .

CMD tail -f /dev/null
# CMD [ "./main" ]