FROM golang:1.20 as builder
ENV CGO_ENABLED=1
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN make build

FROM gcr.io/distroless/static-debian11:nonroot
COPY --from=builder --chown=nonroot:nonroot /app/bin/trial-api /app/trial-api
EXPOSE 8000
CMD [ "/app/trial-api" ]