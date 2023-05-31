FROM golang:1.19-buster
WORKDIR /src
COPY tfproviderdocs /usr/bin/tfproviderdocs
ENTRYPOINT ["/usr/bin/tfproviderdocs"]
