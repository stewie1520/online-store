FROM golang:1.18

RUN go install github.com/cespare/reflex@latest
RUN go install github.com/google/wire/cmd/wire@latest

COPY . /app/

EXPOSE 8080

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]
