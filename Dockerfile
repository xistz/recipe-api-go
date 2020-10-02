FROM golang:1.15-alpine as base

RUN apk --update add build-base

WORKDIR /api

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

FROM base as development

RUN go build -o /out/recipe-api

FROM base as builder

RUN CGO_ENABLED=0 go build -o /out/recipe-api

FROM scratch as deployment

COPY --from=builder /out/recipe-api /recipe-api

CMD ["/recipe-api"]
