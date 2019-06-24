FROM golang:1.11-alpine AS build-env

RUN apk update
RUN apk upgrade
RUN apk add --no-cache curl
RUN apk add --no-cache git
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
WORKDIR /go/src/github.com/ilhammhdd/kudaki-gateway-service/
COPY . /go/src/github.com/ilhammhdd/kudaki-gateway-service/
RUN dep ensure
RUN go build -o kudaki_gateway_service_app

FROM alpine

ARG KAFKA_VERSION
ARG ADDRESS
ARG GRPC_PORT
ARG KAFKA_BROKERS
ARG REST_PORT
ARG USER_AUTH_SERVICE_GRPC_ADDRESS
ARG REST_SUBDOMAIN
ARG GRPC_SUBDOMAIN
ARG DOCKER_NETWORK
ARG REDISEARCH_SERVER

ENV KAFKA_VERSION=$KAFKA_VERSION
ENV ADDRESS=$ADDRESS
ENV GRPC_PORT=$GRPC_PORT
ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV REST_PORT=$REST_PORT
ENV USER_AUTH_SERVICE_GRPC_ADDRESS=$USER_AUTH_SERVICE_GRPC_ADDRESS
ENV REST_SUBDOMAIN=$REST_SUBDOMAIN
ENV GRPC_SUBDOMAIN=$GRPC_SUBDOMAIN
ENV DOCKER_NETWORK=$DOCKER_NETWORK
ENV REDISEARCH_SERVER=$REDISEARCH_SERVER

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-gateway-service/kudaki_gateway_service_app .

LABEL traefik.docker.network=$DOCKER_NETWORK
LABEL traefik.rest.frontend.rule=Host:$REST_SUBDOMAIN
LABEL traefik.rest.port=$REST_PORT
LABEL traefik.grpc.frontend.rule=Host:$GRPC_SUBDOMAIN
LABEL traefik.grpc.port=$GRPC_PORT

EXPOSE $REST_PORT
EXPOSE $GRPC_PORT

ENTRYPOINT ./kudaki_gateway_service_app
