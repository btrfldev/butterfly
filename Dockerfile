FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /build

#Dependencies
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

#Build
COPY . .
RUN go build -o ./bin/dust ./dust/*.go

#WorkSpace
FROM alpine:3.19 as runner

WORKDIR /dust
COPY --from=builder /build/bin/dust /dust/dust

ENV PORT=1615

#Run
EXPOSE 1615-1615
ENTRYPOINT [ "./dust" ]