FROM scratch
MAINTAINER Gaël Gillard<gael@gaelgillard.com>

EXPOSE 3000

COPY api /
COPY files /files
ENTRYPOINT ["/api"]