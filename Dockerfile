# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

RUN apk update \
    && apk add curl git wget

# Install go-learn
RUN cd /tmp && \
    curl -s https://api.github.com/repos/franktore/go-learn/releases/tags/v0.1.0 \
         | grep "tarball_url" \
         | cut -d : -f 2,3 \
         | tr -d \" \
         | tr -d \, \
         | wget -qi - && \
    tar --strip-components=1 -xvf v0.1* && \
    cp releases/go-learn /opt && \
    cp entry_point.sh /app && \
    rm -rf /tmp/*

ENV PATH="/opt:$PATH"
RUN echo $PATH
CMD ["/app/entry_point.sh"]