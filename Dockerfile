from golang:1.15-alpine

ENV SOURCE_DIR=/usr/go/src/littlecat
ENV BUILD_DIR=/usr/go/build
WORKDIR ${SOURCE_DIR}

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o /littlecat main.go

RUN mkdir -p ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN cp /littlecat ${BUILD_DIR}

CMD ["/littlecat"]