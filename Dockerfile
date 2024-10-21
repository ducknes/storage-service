FROM golang:1.23 as build
WORKDIR /build
COPY . .

RUN mkdir out && \
    mkdir out/database && \
    mv database/mocks/ out/database/mocks/ && \
    mv .config/ out/

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -o out/app

FROM alpine
EXPOSE 8080

RUN apk update && apk add --no-cache tzdata

WORKDIR /app

COPY --from=build /build/out ./
RUN ls -la
RUN chmod +x ./app

ENTRYPOINT ./app