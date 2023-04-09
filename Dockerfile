# --- Stage 1:
FROM golang:1.20-alpine as builder

# ENVs
ENV BUILD_PATH=/go/src/github.com/DiegoSantosWS/pismo
# Install go toolchain
RUN apk update && apk add --no-cache curl gcc git libc-dev
# COPY local files
WORKDIR ${BUILD_PATH}
COPY . .

# Get go dependencies
RUN go mod download

# gosec - Golang Security Checker | Commented due to 1.18 update issue.
RUN curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s -- -b ${GOPATH}/bin latest && \
   gosec ./...

# revive (Go lint successor)
RUN go install github.com/mgechev/revive@latest && \
    revive -config revive-config.toml ./...

# RUN UNIT TEST
RUN go test -test.short -v ./...

# Build dynamically linked Go binary
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
    go build -a -ldflags '-linkmode=external' -installsuffix cgo -o pismo
RUN ls -lah ${BUILD_PATH}/pismo
RUN cp ${BUILD_PATH}/pismo /bin/pismo

# --- Stage 2:
# Use a Docker multi-stage build to create a lean production image.
FROM alpine:3.11
# Install dependencies
RUN apk update && apk add --no-cache ca-certificates tzdata libc6-compat
# Copy binary from builder
COPY --from=builder /bin/pismo /pismo
# Run the application on container startup.
CMD ["/pismo"]
EXPOSE 8080