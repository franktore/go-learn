# syntax=docker/dockerfile:1

FROM golang:1.19-bullseye as build

ARG GIT_USERNAME
ENV GITUSERNAME=$GIT_USERNAME

WORKDIR /app

RUN apt-get update \
    && apt-get install -y curl git wget openssh-server supervisor \
    && apt-get upgrade -y

# Fetch go-learn
RUN cd /tmp && \
    curl -s https://api.github.com/repos/{$GITUSERNAME}/go-learn/tags \
         | grep "tarball_url" \
         | cut -d : -f 2,3 \
         | tr -d \" \
         | tr -d \, \
         | wget -qi - && \
    tar --strip-components=1 -xvf v0.1.0* && \
    cp releases/go-learn /app && \
    cp -r deploy/ /app/config && \
    rm -rf /tmp/*

## Deploy final stage, get rid of build stage
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app/go-learn /go-learn
COPY --from=build /app/config /config

EXPOSE 80 2222

USER nonroot:nonroot

ENV PATH="/:$PATH"

CMD [ "go-learn" ]