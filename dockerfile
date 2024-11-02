FROM golang:latest

#Set Working Directory
WORKDIR /usr/src/app

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY .env ./

COPY . .

RUN make install \
  && make build

# Expose port
EXPOSE 8080

# Run application
CMD ["make", "start"]
