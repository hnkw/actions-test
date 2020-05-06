FROM golang:1.11

ADD ./actions-test /usr/local/bin/

ENTRYPOINT [ "actions-test" ]
