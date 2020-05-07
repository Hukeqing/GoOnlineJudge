# GoOnlineJudge

GoOnlineJudge is an ACM/ICPC online judge platform.

[**Demo**](http://acm.zjgsu.edu.cn)

## Contents
+ [Installation](https://github.com/ZJGSU-ACM/GoOnlineJudge#installation)
	+ [Prerequisites](https://github.com/ZJGSU-ACM/GoOnlineJudge#prerequisites)
	+ [Docker](https://github.com/ZJGSU-ACM/GoOnlineJudge#docker)
	+ [Quick Start](https://github.com/ZJGSU-ACM/GoOnlineJudge#quick-start)
	+ [Manual Installation](https://github.com/ZJGSU-ACM/GoOnlineJudge#manual-installation)
	+ [Tips](https://github.com/ZJGSU-ACM/GoOnlineJudge#tips)
+ [Maintainers](https://github.com/ZJGSU-ACM/GoOnlineJudge#maintainers)
+ [Contributions](https://github.com/ZJGSU-ACM/GoOnlineJudge#contributions)
+ [Roadmap](https://github.com/ZJGSU-ACM/GoOnlineJudge#roadmap)
+ [License](https://github.com/ZJGSU-ACM/GoOnlineJudge#license)

## Installation
### Prerequisites

**Disclaimer**:

GoOnlineJudge works best on GNU/Linux and has been tested on Ubuntu 14.04+. Windows and Mac OS X are **not** recommended because [**RunServer**](https://github.com/ZJGSU-ACM/RunServer) cannot be built on both of them. 

### Docker

If you are Windows or Mac OS X user, you can try out [docker-oj](https://github.com/ZJGSU-ACM/docker-oj), based on docker image and works out of the box.

### Quick Start

Be careful! This section might be out-of-date. **Always** check the Manual Installation guide for your safety.

GoOnlineJudge is installed by running the following commands in your terminal. You can install it via the command-line with  `curl`.

#### via curl
```bash
curl -sSL https://raw.githubusercontent.com/ZJGSU-ACM/GoOnlineJudge/master/install.sh | sh
```

### Manual Installation
#### Dependences
+ Go
  + GoOnlineJudge is mainly written in Go. 
  + Get Go from [golang.org](http://golang.org)

+ MongoDB
  + MongoDB is a cross-platform document-oriented databases.
  + Get MongoDB from [MongoDB.org](https://www.mongodb.org/)

+ mgo.v2
  + mgo.v2 offers a rich MongoDB driver for Go.
  + Get mgo.v2 via
  ```
  go get gopkg.in/mgo.v2
  ```
  + API documentation is available on [godoc](http://godoc.org/gopkg.in/mgo.v2)

+ [flex](http://flex.sourceforge.net/)
  + flex is the lexical analyzer used in [**RunServer**](https://github.com/ZJGSU-ACM/RunServer).
  + Get flex using following command if you are running Ubuntu.
  ```bash
  sudo apt-get install flex
  ```

+ [SIM](http://www.dickgrune.com/Programs/similarity_tester/)
  + SIM is a software and text similarity tester. It's used in RunServer.
  + SIM is shipped along with RunServer.

+ [GCC](https://gcc.gnu.org/)
  + The GNU compiler Collection.
  + Get GCC from [GNU](https://gcc.gnu.org) or using following command if you are running Ubuntu
  ```bash
  sudo apt-get gcc-8 g++-8 gdb make dpkg-dev
  ``` 

+ [AdoptOpenJDK](https://adoptopenjdk.net/)
  + AdoptOpenJDK provides rock-solid OpenJDK binaries for the Java ecosystem and also provides infrastructure as code, and a Build farm for builders of OpenJDK, on any platform.
  + AdoptOpenJDK is used to judge Java code.
  + hotspot has a better cpu performance than openj9.
  ```bash
  sudo apt-get install adoptopenjdk-8-hotspot
  ```

+ [Python](https://www.python.org/)
  + Python2 and Python3 is used to judge Python code.
  + Use virtualenv to create a clean environment.
  ```bash
  sudo apt-get install python python-virtualenv python3 python3-virtualenv
  ```

+ iconv-go
  + iconv-go provides iconv support for Go.
  + Get iconv-go via
  ```
  go get github.com/djimenez/iconv-go
  ```

#### Install

Create Python2 and Python3 Virtual Environment
```bash
mkdir -p /usr/local/cjudger/venv/
virtualenv --no-setuptools --no-pip --no-wheel -p /usr/bin/python2 /usr/local/cjudger/venv/py2
virtualenv --no-setuptools --no-pip --no-wheel -p /usr/bin/python3 /usr/local/cjudger/venv/py3
ln -s /usr/local/cjudger/venv/py2/bin/python2 /usr/local/cjudger/py2
ln -s /usr/local/cjudger/venv/py3/bin/python3 /usr/local/cjudger/py3
```

Enable Golang Module and set Module Proxy
```bash
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
```

Obtain latest version via `git` and `go get`, source codes will be in your $GOPATH/src. 
```bash
mkdir -p $GOPATH/src/github.com/ZJGSU-ACM
cd github.com/ZJGSU-ACM
git clone https://github.com/ZJGSU-ACM/restweb.git
git clone https://github.com/ZJGSU-ACM/vjudger.git
git clone https://github.com/ZJGSU-ACM/GoOnlineJudge.git
git clone https://github.com/ZJGSU-ACM/RunServer.git
```

```bash
# Set $OJ_HOME variable
export OJ_HOME="yourself oj home" #e.g. export OJ_HOME=/$GOPATH/src
export DATA_PATH=$OJ_HOME/Data
export LOG_PATH=$OJ_HOME/log
export RUN_PATH=$OJ_HOME/run
export JUDGE_HOST=your_judge_host #e.g. export JUDGE_HOST="http://127.0.0.1:8888"
export MONGODB_PORT_27017_TCP_ADDR=127.0.0.1
export PATH=$PATH:$GOPATH/bin

# directory for MongoDB Data
mkdir $OJ_HOME/Data

# directory for problem set
mkdir $OJ_HOME/ProblemData

# directory for running user's code
mkdir $OJ_HOME/run

# directory for log
mkdir $OJ_HOME/log
```

Make sure you have these directories in your $GOPATH/src:

```
github.com/ZJGSU-ACM/restweb/
github.com/ZJGSU-ACM/GoOnlineJudge/
github.com/ZJGSU-ACM/RunServer/
github.com/ZJGSU-ACM/vjudge/   
gopkg.in/  
restweb/  
```

And these directories in your $OJ_HOME:

```
ProblemData/  
run/  
log/  
```

Now, it's time for compilation.
```bash
cd $GOPATH/src/github.com/ZJGSU-ACM/restweb
go install ./...
cd $GOPATH/src
restweb build github.com/ZJGSU-ACM/GoOnlineJudge
cd $GOPATH/src/github.com/ZJGSU-ACM/RunServer
./make.sh
```

Start OJ Web
```bash
cd $GOPATH/src
restweb run GoOnlineJudge
```
Start Judge
```bash
RunServer
```
Now,you can visit OJ on [http://127.0.0.1:8080](http://127.0.0.1:8080).

#### Tips

+ You should always run MongoDB first then followed by OJ.

+ Running web server at 80 port requires administrator privileges. For security reasons, do **not** run our OJ at 80 port.

+ If you want to visit OJ at 80 port, [nginx](http://nginx.org), the HTTP and reverse proxy server is recommended.

## Maintainers

+ memelee

+ sakeven

+ clarkzjw

+ rex-zed

+ happier233

## Roadmap
+ Binary packaging on mainstream distributions
+ Maybe a built-in simple blog
+ Rebase RESTful API
+ Modern design front-end

## Contributions
+ We are open for all kinds of pull requests!

+ Just please follow the [Golang style guide](./docs/Golang_Style_Guide.md).

## License
See [LICENSE](LICENSE)
