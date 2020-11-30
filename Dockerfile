FROM golang:1.15.0-alpine3.9 as debug

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git \
        dpkg \
        gcc \
        git \
        musl-dev

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH


RUN go get  github.com/awalterschulze/gographviz v0.0.0-20190522210029-fa59802746ab
RUN go get	github.com/gophercloud/gophercloud v0.9.0 // indirect
RUN go get 	github.com/imdario/mergo v0.3.8 // indirect
RUN go get 	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
RUN go get 	golang.org/x/time v0.0.0-20191024005414-555d28b269f0 // indirect
RUN go get 	k8s.io/api v0.17.4
RUN go get 	k8s.io/apimachinery v0.17.4
RUN go get 	k8s.io/client-go v0.17.0
RUN go get 	k8s.io/utils v0.0.0-20200318093247-d1ab8797c558 // indirect

WORKDIR /go/src/work

COPY . /go/src/work/

RUN go build -o k8mapper .

FROM alpine:3.9 as prod
COPY --from=debug /go/src/work/app /
CMD ./app
