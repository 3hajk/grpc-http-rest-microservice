FROM golang:1.17-alpine as builder
RUN apk add --no-cache ca-certificates git openssh openssl make gcc
ADD . /src
WORKDIR /src
RUN make

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN addgroup -g 1001 -S worker \
    && adduser -S -D -H -u 1001 -s /sbin/nologin -G worker -g worker worker
RUN mkdir -p /app
COPY --from=builder /src/bin/service /app
WORKDIR /app

# Metadata params
ARG VERSION
ARG BUILD_DATE
ARG VCS_URL
ARG VCS_REF
ARG NAME
ARG VENDOR

# Metadata
LABEL org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.name=$NAME \
      org.label-schema.vcs-url=https://github.com/3hajk/grpc-http-rest-microservice/$VCS_URL \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vendor=$VENDOR \
      org.label-schema.version=$VERSION \
      org.label-schema.docker.schema-version="0.0.1"

USER worker
ARG CI_COMMIT_SHORT_SHA=00000000
ENV SENTRY_RELEASE=${CI_COMMIT_SHORT_SHA}
CMD ["./service"]
