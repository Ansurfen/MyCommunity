FROM ubuntu:20.04

EXPOSE 9090

COPY /server/main .
COPY /server/application.yml .

RUN chmod +x ./main
RUN mkdir ./images/community -p
RUN mkdir ./images/user -p

CMD [ "./main" ]