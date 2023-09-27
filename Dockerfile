FROM golang:1.21.0 AS build

WORKDIR /src
COPY . .
RUN go build -o /out/ops-tech-challenge

FROM alpine:3.14 AS runtime
COPY --from=build /out/ops-tech-challenge /app/ops-tech-challenge
ENV SECRET="tymonsecret"

WORKDIR /app
CMD ["./ops-tech-challenge"]
