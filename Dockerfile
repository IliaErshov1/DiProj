FROM golang:alpine as builder
ADD . /app/
WORKDIR /app 

COPY . .

RUN go build -o main .

#Run stage
FROM alpine
#ADD . /app/
WORKDIR /app

COPY --from=builder /app/main .
COPY . .
EXPOSE 8080

CMD [ "/app/main" ]
