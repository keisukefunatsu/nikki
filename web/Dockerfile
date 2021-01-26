FROM golang:1.15.7-buster

RUN apt update \
    && apt install -y -q build-essential sqlite3 libsqlite3-dev postgresql libpq-dev vim

# Installing Node 14
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash 
RUN apt update && apt install nodejs

# Installing yarn
RUN npm install -g --no-progress yarn \
    && yarn config set yarn-offline-mirror /npm-packages-offline-cache \
    && yarn config set yarn-offline-mirror-pruning true

EXPOSE 3000
ENV HOST=0.0.0.0

RUN mkdir -p /web/src
WORKDIR /web/src
