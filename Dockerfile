FROM golang:1.19-alpine AS build
WORKDIR /app
COPY ./app/go.mod ./app/go.sum /app/
RUN go mod download
COPY ./app/*.go /app/
RUN go build -o /ctf

FROM nginx:alpine
RUN apk add --no-cache bash
RUN apk add --no-cache vim
COPY static-html-directory /usr/share/nginx/html
COPY /custom.conf /etc/nginx/conf.d/default.conf
COPY /admin /usr/share/nginx/html/admin
COPY --from=build /ctf /ctf
COPY ./40-run-ctf.sh /docker-entrypoint.d/
