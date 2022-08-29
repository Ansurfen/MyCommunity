FROM ubuntu:20.04

EXPOSE 9090

COPY /server/main .
COPY /server/application.yml .

CMD [ "./main" ]