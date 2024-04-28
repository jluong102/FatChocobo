FROM debian:stable

RUN mkdir /app # Have everything in its own dir
COPY config.json /app
COPY Makefile /app
COPY go.mod /app
COPY go.sum /app
COPY cmd/. /app/cmd
COPY pkg/. /app/pkg

RUN apt update
RUN apt install -y golang
RUN apt install -y make
RUN apt install -y curl # To avoid x509 errors

WORKDIR /app
RUN make

CMD ["./fatchocobo"]
