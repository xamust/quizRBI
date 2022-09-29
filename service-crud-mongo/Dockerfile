FROM golang:latest
MAINTAINER Stepan K. <xamust@gmail.com>
WORKDIR /service-crud-mongo/
VOLUME ["/opt/service-crud-mongo"]
COPY . ./
RUN make build
CMD [ "build/service-crud-mongo" ]