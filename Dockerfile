FROM golang:1.14.2
MAINTAINER Happier233 "happier233@qq.com"

ENV OJ_HOME $GOPATH/src
ENV DATA_PATH $OJ_HOME/Data
ENV LOG_PATH $OJ_HOME/log
ENV RUN_PATH $OJ_HOME/run

WORKDIR $GOPATH/src/

RUN mkdir -p $OJ_HOME/log

RUN mkdir -p $GOPATH/src/github.com/ZJGSU-ACM
ADD . $GOPATH/src/github.com/ZJGSU-ACM/GoOnlineJudge

# Get dependence
RUN git clone --depth=1 https://github.com/ZJGSU-ACM/restweb.git $GOPATH/src/github.com/ZJGSU-ACM/restweb

# Enable Go Module
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn

# Build OJ
RUN cd $GOPATH/src/github.com/ZJGSU-ACM/restweb/restweb && go install
RUN cd $GOPATH/src && restweb build github.com/ZJGSU-ACM/GoOnlineJudge

EXPOSE 8080

CMD cd github.com/ZJGSU-ACM/GoOnlineJudge; $GOPATH/bin/GoOnlineJudge
