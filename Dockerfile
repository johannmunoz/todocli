FROM golang:1.18-alpine as dev

WORKDIR /work

FROM golang:1.18-alpine as build

WORKDIR /todocli
COPY ./todocli/* /todocli/
RUN go build -o todocli

FROM alpine as runtime 
COPY --from=build /todocli/todocli /usr/local/bin/todocli
COPY ./todocli/todos.json /
COPY run.sh /
RUN chmod +x /run.sh

ENTRYPOINT ["./run.sh"]