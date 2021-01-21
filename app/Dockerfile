FROM gobuffalo/buffalo:latest

ARG CODECOV_TOKEN

ENV GOPROXY         https://proxy.golang.org
ENV GO111MODULE     on

RUN wget https://github.com/gobuffalo/buffalo/releases/download/v0.16.21/buffalo_0.16.21_Linux_x86_64.tar.gz
RUN tar -xvzf buffalo_0.16.21_Linux_x86_64.tar.gz 
RUN mv buffalo /usr/local/bin/buffalo 


RUN mkdir -p /app/src
COPY ./app/src /app/src
ENV ADDR=0.0.0.0
EXPOSE 3000
WORKDIR /app/src

RUN go get -u -v -tags sqlite github.com/gobuffalo/pop/...
RUN go install -tags sqlite github.com/gobuffalo/pop/soda
# RUN go mod tidy
# RUN go test -tags "sqlite integration_test" -cover -race -v ./...