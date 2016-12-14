FROM ubuntu:trusty

RUN apt-get update
RUN apt-get install -y wget gcc make flex bison git vim

# vcap user creation
RUN useradd -ms /bin/bash vcap
ENV VCAP_HOME /home/vcap
WORKDIR $VCAP_HOME
RUN mkdir $VCAP_HOME/bin
RUN chown vcap:vcap $VCAP_HOME/bin

# copy monit config
ADD docker/jobs $VCAP_HOME/jobs
ADD docker/monitrc $VCAP_HOME/monitrc
RUN chmod 700 $VCAP_HOME/monitrc
RUN chmod +rx $VCAP_HOME/jobs/foo
RUN chown vcap:vcap $VCAP_HOME/monitrc $VCAP_HOME/jobs $VCAP_HOME/jobs/*

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
RUN wget https://storage.googleapis.com/golang/go1.7.4.linux-amd64.tar.gz
RUN tar -xvf go1.7.4.linux-amd64.tar.gz --strip 1 -C $GOROOT
ENV PATH "$PATH:$GOROOT/bin:$GOPATH/bin"
RUN rm go1.7.4.linux-amd64.tar.gz

# go packages installation
RUN go get github.com/onsi/ginkgo/ginkgo
RUN go get github.com/onsi/gomega
RUN go get github.com/tools/godep
