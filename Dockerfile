FROM golang:1.18 AS builder

WORKDIR /

COPY go.mod main.go /

RUN go mod download

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

#Using multistage builds to lower image size
FROM scratch

WORKDIR /

COPY --from=builder /main ./

EXPOSE 8080

CMD ["./main"]