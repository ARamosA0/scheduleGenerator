FROM golang:1.24

WORKDIR /app

RUN go install github.com/cosmtrek/air@v1.49.0

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

COPY . .
# RUN go build -v -o /usr/local/bin/app .

COPY migrate/migrate.go ./migrate/
# RUN go build -o /usr/local/bin/migrate ./migrate

CMD ["air"]
