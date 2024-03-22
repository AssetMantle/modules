FROM golang:1.17
COPY . .
RUN make test
