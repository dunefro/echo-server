FROM golang:1.16-buster AS build
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go build -o /echo-server

## 
FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /echo-server /echo-server
EXPOSE 8080
ENTRYPOINT [ "/echo-server" ]