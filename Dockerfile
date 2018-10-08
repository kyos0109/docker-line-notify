FROM golang:1.11.1

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o line-notify

FROM plugins/base:multiarch

COPY --from=0 /app/line-notify /bin/

LABEL maintainer="kyos0109"

ENTRYPOINT [ "/bin/line-notify" ]
