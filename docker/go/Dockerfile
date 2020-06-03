FROM golang:1.14-alpine

ENV PACKAGES="bash \
    bash-completion \
    git \
    nano \
    vim \
    curl \
    net-tools \
    coreutils \
    sudo \
    build-base \
    "

RUN apk add ${PACKAGES}

ARG USER_GID=1000
ARG USER_UID=1000
RUN addgroup -g ${USER_GID} app \
    && adduser -D -s /bin/bash -G app -u ${USER_UID} app \
    && chown -R app:app /opt \
    && echo 'app ALL=NOPASSWD: ALL' > /etc/sudoers.d/app \
    && chmod 0440 /etc/sudoers.d/app \
    && visudo --check

RUN chown -R app:app /go

USER app
WORKDIR /go

RUN go get -v golang.org/x/tools/cmd/godoc \
    && go get -v golang.org/x/lint/golint \
    && go get -v golang.org/x/tools/cmd/gotype \
    && go get -v github.com/mdempsky/gocode \
    && go get -v github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get -v github.com/ramya-rao-a/go-outline \
    && go get -v github.com/stamblerre/gocode \
    && go get -v github.com/rogpeppe/godef \
    && go get -v github.com/sqs/goreturns \
    && go get -v golang.org/x/tools/cmd/goimports \
    && go get -v github.com/go-delve/delve/cmd/dlv

ARG MODULE_PATH
ENV MODULE_ABS_PATH="/go/src/${MODULE_PATH}"

WORKDIR ${MODULE_ABS_PATH}

ENTRYPOINT [ "godoc", "-http=:6060" ]