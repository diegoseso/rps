FROM debian

RUN apt-get update
RUN apt-get install -y nginx curl gnupg2 git

RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get install -y nodejs && nodejs -v && npm -v 
RUN npm install -g bower

RUN curl -O https://dl.google.com/go/go1.10.2.linux-amd64.tar.gz
RUN sha256sum go1.10*.tar.gz
# 4b677d698c65370afa33757b6954ade60347aaca310ea92a63ed717d7cb0c2ff
RUN tar xvf go1.10.2.linux-amd64.tar.gz && chown -R root:root ./go && mv go /usr/local

RUN echo $HOME
ENV GOPATH=$HOME/work
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

WORKDIR /var/www/html/sps

EXPOSE 80 8080