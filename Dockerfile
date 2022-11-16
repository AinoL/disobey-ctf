FROM golang:1.19-alpine AS build
WORKDIR /app
COPY ./app/go.mod ./app/go.sum /app/
RUN go mod download
COPY ./app/*.go /app/
RUN go build -o /ctf

FROM nginx:alpine

ENV NGINX_HOSTNAME "localhost"

COPY ./custom.conf /etc/nginx/templates/default.conf.template
COPY ./admin /usr/share/nginx/html/admin
COPY --from=build /ctf /ctf
COPY ./images /images
COPY ./templates /templates
COPY ./static /static
COPY ./40-run-ctf.sh /docker-entrypoint.d/
