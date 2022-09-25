FROM debian:stretch-slim

WORKDIR /

ADD wanna-scheduler /usr/local/bin/wanna-scheduler

CMD ["wanna-scheduler"]