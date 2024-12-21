FROM golang

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go", "run", "/app/main.go"]

EXPOSE 3001