FROM golang:1.25-alpine AS builder

WORKDIR /testing

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

FROM scratch

COPY --from=builder /testing/main ./main

CMD ["./main"]