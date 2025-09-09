FROM golang:1.25-alpine3.22 AS build

WORKDIR /app
COPY . .

RUN go mod tidy && go build -o server cmd/main.go

# ----------------------------------

FROM alpine:3.22

LABEL com.github.rvolykh.authors="rvolykh@gmail.com"

WORKDIR /app
COPY --from=build /app/server .

EXPOSE 8080
USER nobody

CMD ["./server", "-host", "0.0.0.0", "-port", "8080"]
