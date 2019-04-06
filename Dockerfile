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

# ARG ADDRESS
# ARG GRPC_PORT
ARG KAFKA_BROKERS
ARG REST_PORT
ARG USER_SERVICE_GRPC_ADDRESS
ARG REST_FRONTEND_RULE
ARG GRPC_FRONTEND_RULE

# ENV ADDRESS=$ADDRESS
# ENV GRPC_PORT=$GRPC_PORT
ENV KAFKA_BROKERS=$KAFKA_BROKERS
ENV REST_PORT=$REST_PORT
ENV USER_SERVICE_GRPC_ADDRESS=$USER_SERVICE_GRPC_ADDRESS
ENV REST_FRONTEND_RULE=$REST_FRONTEND_RULE
ENV GRPC_FRONTEND_RULE=$GRPC_FRONTEND_RULE

COPY --from=build-env /go/src/github.com/ilhammhdd/kudaki-gateway-service/kudaki_gateway_service_app .

LABEL traefik.frontend.rule=Host:$REST_FRONTEND_RULE
LABEL traefik.rest.port=$REST_PORT
LABEL traefik.grpc.frontend.rule=Host:$GRPC_FRONTEND_RULE
LABEL traefik.grpc.port=$GRPC_PORT

EXPOSE $REST_PORT
EXPOSE $GRPC_PORT

ENTRYPOINT ./kudaki_gateway_service_app
