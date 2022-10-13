# syntax=docker/dockerfile:1

FROM golang:alpine as build

ARG GIT_USERNAME
ENV GITUSERNAME=$GIT_USERNAME

WORKDIR /app
ENV PATH="/app:$PATH"

# RUN apt-get update \
#     && apt-get install -y curl git wget bash openssh-server supervisor \
#     && apt-get upgrade -y

RUN apk update \
    && apk add --no-cache curl git wget bash openssh-server supervisor vim

# create and set user (student) permissions
# RUN adduser -G www-data student --system --no-create-home && \
#     chgrp -R www-data /app /go /etc /var /home /bin && \
#     chmod -R 777 /app /go /etc /var /home /bin
# USER student:www-data

# Fetch go-learn
# RUN cd /tmp && \
#     curl -s https://api.github.com/repos/{$GITUSERNAME}/go-learn/tags \
#          | grep "tarball_url" \
#          | cut -d : -f 2,3 \
#          | tr -d \" \
#          | tr -d \, \
#          | wget -qi - && \
#     tar --strip-components=1 -xvf v0.1.0* && \
#     cp releases/go-learn /app && \
#     cp -r deploy/ /app/config && \
#     rm -rf /tmp/*
RUN git clone https://github.com/franktore/go-learn.git
RUN cd go-learn && \
    go build -o /go/bin/go-learn
COPY entry_point.sh .
COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY sshd_config /etc/ssh/sshd_config

RUN cd /etc/ssh/ && \
    ssh-keygen -A

EXPOSE 80 2222

RUN chmod +x entry_point.sh

CMD ["entry_point.sh"]
# CMD [ "/usr/bin/supervisord -c /etc/supervisor/supervisord.conf" ]

## Deploy final stage, get rid of build stage
# FROM gcr.io/distroless/base-debian10

# WORKDIR /

# COPY --from=build /bin/sh /bin/sh
# COPY --from=build /app/go-learn /go-learn
# COPY --from=build /app/config /config

# EXPOSE 80 2222

# USER nonroot:nonroot

# ENV PATH="/:$PATH"

# CMD [ "go-learn" ]