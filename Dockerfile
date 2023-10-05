FROM ubuntu:trusty

RUN \
  apt-get update && \
  apt-get install -y \
    wget gcc make flex bison git vim openssh-server

# allow SSH connections to container
EXPOSE 22

# redis installation
RUN wget http://download.redis.io/releases/redis-3.2.6.tar.gz
RUN tar -xvf redis-3.2.6.tar.gz
RUN cd redis-3.2.6 && make && make install
RUN rm -r redis-3.2.6 redis-3.2.6.tar.gz

# vcap user creation
RUN useradd -ms /bin/bash vcap
RUN echo vcap:vcap | chpasswd
ENV VCAP_HOME /home/vcap
WORKDIR $VCAP_HOME
RUN mkdir $VCAP_HOME/bin
RUN chown vcap:vcap $VCAP_HOME/bin

# copy monit config
ADD docker/jobs $VCAP_HOME/jobs
ADD docker/monitrc $VCAP_HOME/monitrc
RUN chmod 700 $VCAP_HOME/monitrc
RUN chmod +rx $VCAP_HOME/jobs/foo $VCAP_HOME/jobs/bar $VCAP_HOME/jobs/baz
RUN chown vcap:vcap $VCAP_HOME/monitrc $VCAP_HOME/jobs $VCAP_HOME/jobs/*

# copy test script
ADD docker/test.sh $VCAP_HOME/test.sh
RUN chmod +rx $VCAP_HOME/test.sh
RUN chown vcap:vcap $VCAP_HOME/test.sh

# switch to vcap
USER vcap
ENV HOME /home/vcap
ENV PATH "$PATH:$HOME/bin"
WORKDIR $HOME

# monit installation
RUN wget https://mmonit.com/monit/dist/monit-5.2.5.tar.gz
RUN tar -xvf monit-5.2.5.tar.gz
RUN cd monit-5.2.5 && ./configure --without-ssl --prefix=$HOME
RUN cd monit-5.2.5 && make && make install
RUN rm -r monit-5.2.5 monit-5.2.5.tar.gz

# go installation
ENV GOPATH $HOME/go
ENV GOROOT $HOME/goroot
RUN mkdir $GOROOT $GOPATH
RUN wget https://storage.googleapis.com/golang/go1.10.linux-amd64.tar.gz
RUN tar -xvf go1.10.linux-amd64.tar.gz --strip 1 -C $GOROOT
ENV PATH "$PATH:$GOROOT/bin:$GOPATH/bin"
RUN rm go1.10.linux-amd64.tar.gz

# go packages installation
RUN go get github.com/onsi/ginkgo/v2/ginkgo
RUN cd $GOPATH/src/github.com/onsi/ginkgo/v2/ginkgo && git checkout 1d2fb67b14d3a770782be056751836928af84c5d
RUN go install github.com/onsi/ginkgo/v2/ginkgo
RUN go get github.com/tools/godep
