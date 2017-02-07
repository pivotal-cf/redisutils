FROM ubuntu:trusty

RUN \
  apt-get update && \
  apt-get install -y wget gcc make

RUN wget http://download.redis.io/releases/redis-3.2.6.tar.gz
RUN tar -xvf redis-3.2.6.tar.gz
RUN cd redis-3.2.6 && make && make install
RUN rm -r redis-3.2.6 redis-3.2.6.tar.gz

EXPOSE 6379

ENTRYPOINT  ["/usr/local/bin/redis-server", "--protected-mode", "no"]
