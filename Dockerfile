FROM golang:1.21.1-alpine3.18

ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV PACKAGES="ca-certificates git curl bash zsh make"
ENV ROOT /app

RUN apk update \
    && apk add --no-cache ${PACKAGES} \
    && update-ca-certificates

WORKDIR ${ROOT}

COPY go.mod go.sum ./

RUN go mod download

EXPOSE 8989

CMD ["go", "run", "."]
