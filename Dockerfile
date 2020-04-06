FROM golang:1.11.2-alpine
MAINTAINER weizhe.chang@gmail.com

COPY mailwithoutsideip /tmp

ADD crontab.txt /crontab.txt

WORKDIR /tmp


RUN chmod 777 ./mailwithoutsideip
RUN /usr/bin/crontab /crontab.txt

CMD ["/tmp/mailwithoutsideip"]