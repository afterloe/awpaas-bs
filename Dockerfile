FROM alpine:3.8
MAINTAINER afterloe <lm6289511@gmail.com>

ENV \
    PROJECT_DIR="/app"
WORKDIR ${PROJECT_DIR}
COPY app ${PROJECT_DIR}
COPY package.json ${PROJECT_DIR}
EXPOSE 8080
CMD ./app