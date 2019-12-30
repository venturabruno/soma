FROM golang:latest as builder
COPY ./src/soma .
RUN go build -ldflags "-s -w" -o /soma

FROM scratch
COPY --from=builder /soma /soma
ENTRYPOINT ["/soma"]