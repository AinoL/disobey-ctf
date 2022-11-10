FROM nginx:alpine
RUN apk add --no-cache bash
RUN apk add --no-cache vim
COPY static-html-directory /usr/share/nginx/html
COPY /custom.conf /etc/nginx/conf.d/default.conf
COPY /admin/index.html /usr/share/nginx/html/admin/index.html