FROM golang:1.15.5-alpine3.12 as debug

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev    
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /go/src/work

COPY . /go/src/work/

RUN go build -o main .

FROM alpine:3.12 as prod
COPY --from=debug /go/src/work .

CMD "./main"