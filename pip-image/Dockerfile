FROM alpine:latest

WORKDIR /workspace

COPY requirements.txt ./
COPY init_local_pip.sh ./

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --update --no-cache curl jq py3-configobj py3-pip py3-setuptools python3 python3-dev

RUN pip config set global.index-url https://mirrors.aliyun.com/pypi/simple/ & pip config set install.trusted-host mirrors.aliyun.com

RUN pip install --upgrade setuptools & pip download -d /python-packages -r requirements.txt

RUN init_local_pip.sh /python-packages

RUN pip install pypiserver

CMD ["pypi-server", "/python-packages", "&>/var/log/pypi-server.log & exit 0"]

