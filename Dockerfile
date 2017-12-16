FROM golang:1.9.2
MAINTAINER Peter Zhang

RUN echo "alias l='ls -CF'\nalias ll='ls -alF'\nalias ls='ls --color=auto'" >> /root/.bashrc

ENV APP_HOME /go/src/github.com/pipizhang/korkort
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
ADD . $APP_HOME

RUN apt-get update -qq
RUN apt-get install -y sqlite3
RUN go get github.com/Masterminds/glide
RUN glide install
RUN glide rebuild

CMD ["bash"]

