# Copyright 2019 Core Services Team.

FROM golang:1.18-alpine as builder

WORKDIR /medicplus

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go install ./cmd/medicplus

FROM alpine as production
COPY --from=builder /go/bin /bin
USER nobody:nobody
EXPOSE 8001
ENTRYPOINT ["/bin/medicplus"]
