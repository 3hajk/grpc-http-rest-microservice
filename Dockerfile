FROM golang:1.17-alpine as builder
ARG SSH_PRIVATE_KEY
ARG GITLAB_USER_EMAIL
ARG GITLAB_USER_NAME
RUN apk add --no-cache ca-certificates git openssh openssl make gcc
RUN git config --global user.email "$GITLAB_USER_EMAIL" && \
    git config --global user.name "$GITLAB_USER_NAME"
RUN mkdir -p ~/.ssh && \
    echo "$SSH_PRIVATE_KEY" > ~/.ssh/id_rsa && \
    chmod -R 600 ~/.ssh && \
    ssh-keyscan -t rsa ssh.charly.guru >> ~/.ssh/known_hosts
ADD . /src
WORKDIR /src
RUN make

FROM alpine:latest
RUN apk add --no-cache ca-certificates
RUN addgroup -g 1001 -S worker \
    && adduser -S -D -H -u 1001 -s /sbin/nologin -G worker -g worker worker
RUN mkdir -p /app
COPY --from=builder /src/alp-matcher /app
WORKDIR /app
USER worker
ARG CI_COMMIT_SHORT_SHA=00000000
ENV SENTRY_RELEASE=${CI_COMMIT_SHORT_SHA}
CMD ["./service"]
