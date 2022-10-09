# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye

WORKDIR /app

RUN apt-get update \
    && apt-get install -y curl git wget bash \
    && apt-get upgrade -y

# Install go-learn
RUN cd /tmp && \
    curl -s https://api.github.com/repos/franktore/go-learn/releases/tags/v0.1.0 \
         | grep "tarball_url" \
         | cut -d : -f 2,3 \
         | tr -d \" \
         | tr -d \, \
         | wget -qi - && \
    tar --strip-components=1 -xvf v0.1* && \
    cp releases/go-learn /app && \
    rm -rf /tmp/*

ENV PATH="/app:$PATH"

CMD [ "go-learn" ]