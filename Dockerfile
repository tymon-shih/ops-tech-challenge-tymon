FROM golang:latest AS build

WORKDIR /src
COPY . .
RUN go build -o /out/ops-tech-challenge

FROM alpine:3.14 AS runtime
COPY --from=build /out/ops-tech-challenge /app/ops-tech-challenge
ENV SECRET="tymonsecret"

WORKDIR /app
EXPOSE 8080
CMD ["./ops-tech-challenge"]
