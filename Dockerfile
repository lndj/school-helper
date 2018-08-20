FROM golang:1.9-alpine

RUN mkdir -p $(go env GOPATH)/src/github.com/lndj/school_helper \
    && apk add --no-cache git mercurial \
    && go get -d -u github.com/golang/dep \
    && cd $(go env GOPATH)/src/github.com/golang/dep \
    && DEP_LATEST=$(git describe --abbrev=0 --tags) \
    && git checkout $DEP_LATEST \
    && go install -ldflags="-X main.version=$DEP_LATEST" ./cmd/dep \
    && git checkout master

COPY . /go/src/github.com/lndj/school_helper
WORKDIR /go/src/github.com/lndj/school_helper

RUN dep ensure -v \
    && go build -o school_helper main.go \
    && apk del git mercurial

ENV APP_ENV production
ENV APP_PORT 5050

EXPOSE 5050

CMD ["school_helper"]
