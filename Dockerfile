FROM golang:1.13
COPY main.go .
RUN go build -o /developer-productivity
CMD ["/developer-productivity"]

