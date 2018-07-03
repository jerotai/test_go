#build stage
FROM golang:1.10.2-alpine3.7 AS build-env
ENV REPO_NAME=Stingray
RUN mkdir -p ${GOPATH}/src/${REPO_NAME}
ADD . ${GOPATH}/src/${REPO_NAME}
WORKDIR ${GOPATH}/src/${REPO_NAME}
RUN apk add --no-cache git  && \
    apk add --no-cache curl

RUN git config --global http.extraheader "PRIVATE-TOKEN: 9HA9sgfyzCcBqWECDpnX" && \
    git config --global url."https://ericcqcp:9HA9sgfyzCcBqWECDpnX@gitlab.com/".insteadOf "git@gitlab.com:" && \
    pwd && \
    ls

RUN go get -u github.com/alecthomas/log4go \
    github.com/gin-gonic/gin \
    github.com/go-redis/redis \
    github.com/spf13/viper \
    github.com/spf13/cobra \
    github.com/aviddiviner/gin-limit
    
RUN go build -o /stingray
RUN cp -r config /


#final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /usr/local/bin
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=build-env /stingray .
COPY --from=build-env /config config
RUN ls /
ENTRYPOINT ./stingray
