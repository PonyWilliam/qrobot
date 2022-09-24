FROM golang:1.18
WORKDIR /src

ADD qrobot /src/
ADD config.yaml /src/


CMD [ "./qrobot" ]