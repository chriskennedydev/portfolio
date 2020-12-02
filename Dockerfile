FROM golang:1.14 as builder
WORKDIR /go/src/portfolio
ENV GOBIN /go/bin
RUN go get github.com/gorilla/handlers
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o site .

FROM scratch 
COPY --from=builder /go/src/portfolio /app/
WORKDIR /app

EXPOSE 5000

CMD [ "./site" ]

