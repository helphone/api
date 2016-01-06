FROM alpine:3.2
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

EXPOSE 3000

ADD api /bin/
ADD files /etc/files

CMD ["/bin/api"]
