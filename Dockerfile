# Build in app in a Go container
FROM docker.io/golang:1.22-alpine as builder
RUN apk update && apk upgrade && apk --no-cache add git
COPY go.mod /app/go.mod
COPY go.sum /app/go.sum
COPY cmd /app/cmd
COPY ent /app/ent
COPY pkg /app/pkg
COPY templates /app/templates
COPY web /app/web
WORKDIR /app
RUN go env -w GOPROXY=direct && go env -w GOSUMDB=off
RUN go mod download
RUN CGO_ENABLED=0 go build -o main cmd/server/main.go
# Move artifact to smaller container with no Go tools installed
FROM docker.io/alpine:3.19.1
RUN apk update && apk upgrade && apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main app
ENTRYPOINT ["/app/app"]