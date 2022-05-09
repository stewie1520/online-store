FROM golang:1.18

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/google/wire/cmd/wire@v0.5.0
RUN go install github.com/go-delve/delve/cmd/dlv@v1.8.3
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@v1.13.0

COPY . /app/

EXPOSE 8080 8181

ENTRYPOINT ["sh", "./start.sh"]
