FROM golang:1.19
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod tidy
COPY . ./
WORKDIR /app/apps/api
RUN go build -o /server
CMD ["sh", "./run.sh"]