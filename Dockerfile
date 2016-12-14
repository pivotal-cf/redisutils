FROM ubuntu:trusty

RUN apt-get update
RUN apt-get install -y wget gcc make flex bison git

ENV GOPATH /root/go
ENV GOROOT /root/goroot
RUN mkdir $GOROOT
RUN wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
RUN tar -xvf go1.7.4.linux-amd64.tar.gz --strip 1 -C $GOROOT
ENV PATH "$PATH:$GOROOT/bin:$GOPATH/bin"
RUN rm go1.7.4.linux-amd64.tar.gz
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
RUN go get github.com/tools/godep

RUN wget https://mmonit.com/monit/dist/monit-5.2.5.tar.gz
RUN tar -xvf monit-5.2.5.tar.gz
RUN cd monit-5.2.5 && ./configure --without-ssl
RUN cd monit-5.2.5 && make && make install
RUN rm -r monit-5.2.5 monit-5.2.5.tar.gz

RUN mkdir -p /root/go/src/github.com/pivotal-cf
ENV PACKAGE_DIR /root/go/src/github.com/pivotal-cf/redisutils
COPY . $PACKAGE_DIR
RUN cp $PACKAGE_DIR/docker/monitrc /etc/monitrc
RUN chmod 700 /etc/monitrc
RUN chmod +rx $PACKAGE_DIR/docker/jobs/foo
RUN chmod +rx $PACKAGE_DIR/docker/test.sh

CMD monit && /bin/bash
