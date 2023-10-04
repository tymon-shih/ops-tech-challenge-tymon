FROM golang:latest AS build

WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 go build -o /out/ops-tech-challenge

FROM alpine:3.14 AS runtime
# FROM golang:latest AS runtime
# ENV SECRET="tymonsecret"

WORKDIR /app
COPY --from=build /out/ops-tech-challenge ops-tech-challenge
EXPOSE 8080

CMD ["./ops-tech-challenge"]
