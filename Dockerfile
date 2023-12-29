FROM golang:1.22-rc-bookworm
RUN go install github.com/cosmtrek/air@latest
WORKDIR /app
COPY . .
ENTRYPOINT ["air"]