FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/cluster-autoscaler-status-exporter cmd/cluster-autoscaler-status-exporter.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /app
COPY --from=builder /app/cluster-autoscaler-status-exporter /app/
COPY docs/samples/config.yaml /app/config.yaml
ENTRYPOINT [ "./cluster-autoscaler-status-exporter" ]
CMD ["--config", "/app/config.yaml"]