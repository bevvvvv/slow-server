#!/bin/bash

apt-get update
apt-get install -y wget

wget https://go.dev/dl/go1.18.3.linux-amd64.tar.gz
tar -zxvf go1.18.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.18.3.linux-amd64.tar.gz

# GOPATH
echo "export PATH=\$PATH:/usr/local/go/bin" >> /etc/profile
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile

go version
