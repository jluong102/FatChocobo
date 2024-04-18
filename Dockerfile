FROM debian:stable

RUN mkdir /pkg # Have everything in its own dir
COPY config.json /pkg
COPY Makefile /pkg
COPY go.mod /pkg
COPY go.sum /pkg
COPY cmd/. /pkg/cmd

RUN apt update
RUN apt install -y golang
RUN apt install -y make
RUN apt install -y curl # To avoid x509 errors

WORKDIR /pkg
RUN make

CMD ["./fatchocobo"]
