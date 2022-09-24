FROM golang:1.18
WORKDIR /src

ADD qrobot /src/
ADD config.yaml /src/

RUN chmod +x ./qrobot

ENTRYPOINT [ "./qrobot" ]

EXPOSE 9000