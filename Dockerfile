FROM golang:1.21.0 AS build

WORKDIR /src
COPY . .
ARG COMMIT_HASH
RUN CGO_ENABLED=0 go build -ldflags "-X main.SECRET=${COMMIT_HASH}" -o /out/ops-tech-challenge

FROM alpine:3.14 AS runtime
# FROM golang:latest AS runtime
# ENV SECRET="tymonsecret"

WORKDIR /app
COPY --from=build /out/ops-tech-challenge ops-tech-challenge
EXPOSE 8080

CMD ["./ops-tech-challenge"]
