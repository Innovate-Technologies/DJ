#FROM meyskens/ubuntu-multiarch:armhf-xenial # arch=armhf
#FROM meyskens/ubuntu-multiarch:amd64-xenial # arch=amd64

RUN apt-get update && apt-get -y upgrade && apt-get -y install git sudo curl unattended-upgrades python build-essential && curl -sL https://deb.nodesource.com/setup_6.x | sudo -E bash - && apt-get -f -y install nodejs

COPY ./overlay-common /dj
RUN cd /dj && npm install
ENV username=""
ENV DEBUG="true"
#ENV compiled="true"

EXPOSE 80

CMD cd /dj/bin && node app.js
