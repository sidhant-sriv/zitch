FROM golang:alpine as builder

RUN mkdir /app
WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download && go mod tidy

COPY ./ ./

RUN go build -o main ./api

# Run stage
FROM alpine
WORKDIR /app
EXPOSE 4000
COPY --from=builder /app/main .
CMD [ "/app/main" ]