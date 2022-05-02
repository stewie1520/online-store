FROM golang:1.18

RUN go get github.com/cespare/reflex

COPY . /app/

EXPOSE 8080

ENTRYPOINT ["reflex", "-c", "./reflex.conf"]
