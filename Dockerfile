FROM golang:1.21
LABEL "com.datadoghq.ad.logs"='[{"source": "go", "service": "instrumentedhttpserver"}]'
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o server .
CMD ["/app/server"]