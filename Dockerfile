# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

# Install go-learn
RUN cd /tmp && \
    curl -s https://api.github.com/repos/franktore/go-learn/releases/tags/v0.1.0 \
         | grep "browser_download_url.*linux-amd64.*tar.gz" \
         | cut -d : -f 2,3 \
         | tr -d \" \
         | wget -qi - && \
    tar --strip-components=1 -xvf go-learn-v* && \ 
    cp go-learn /opt && \
    rm -rf /tmp/*

CMD /opt/go-learn