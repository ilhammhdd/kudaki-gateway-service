ARG ARCH=

FROM ${ARCH}golang:1.15-alpine AS build-env

RUN apk update
RUN apk upgrade
RUN apk add build-base
WORKDIR /go/src/github.com/ilhammhdd/kudaki-gateway-service/
COPY . /go/src/github.com/ilhammhdd/kudaki-gateway-service/
RUN go mod tidy
RUN go build -o kudaki_gateway_service_app

FROM ${ARCH}alpine

ARG KAFKA_VERSION
ARG KAFKA_BROKERS
ARG REST_PORT
ARG USER_AUTH_SERVICE_GRPC_ADDRESS
ARG REST_SUBDOMAIN
ARG DOCKER_NETWORK
ARG REDISEARCH_SERVER
ARG FILE_SERVICE_REST_ADDRESS

ENV KAFKA_VERSION=$KAFKA_VERSION
ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV REST_PORT=$REST_PORT
ENV USER_AUTH_SERVICE_GRPC_ADDRESS=$USER_AUTH_SERVICE_GRPC_ADDRESS
ENV REST_SUBDOMAIN=$REST_SUBDOMAIN
ENV DOCKER_NETWORK=$DOCKER_NETWORK
ENV REDISEARCH_SERVER=$REDISEARCH_SERVER
ENV FILE_SERVICE_REST_ADDRESS=$FILE_SERVICE_REST_ADDRESS

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-gateway-service/kudaki_gateway_service_app .

LABEL traefik.docker.network=$DOCKER_NETWORK
LABEL traefik.rest.frontend.rule=Host:$REST_SUBDOMAIN
LABEL traefik.rest.port=$REST_PORT

EXPOSE $REST_PORT

ENTRYPOINT ./kudaki_gateway_service_app
