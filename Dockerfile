FROM golang:1.20-alpine

LABEL maintainer="araitkozha"
WORKDIR /
COPY . .
RUN go build -o groupie-tracker
EXPOSE 8080
CMD ["./groupie-tracker"]