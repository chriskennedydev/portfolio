FROM golang:1.16 as builder
WORKDIR /usr/local/go/src/portfolio
ENV GOBIN /usr/local/go/bin
COPY . .
RUN go get 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o site .

FROM scratch 
COPY --from=builder /usr/local/go/src/portfolio /app/
WORKDIR /app
ENV DATABASE_URL "postgres://postgres:devpass@172.16.0.126:5432/nufor"
EXPOSE 5000

CMD [ "./site" ]

