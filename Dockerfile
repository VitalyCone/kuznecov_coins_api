FROM golang
WORKDIR /usr/app
COPY ./ ./
ENV DOCKER_ENV=true
RUN go build -o kzcv-coins cmd/apiserver/main.go
CMD ["./kzcv-coins"] 