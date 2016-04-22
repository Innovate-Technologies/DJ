FROM multiarch/ubuntu-core:armhf-xenial
#FROM multiarch/ubuntu-core:amd64-xenial # arch=amd64

RUN apt-get update && apt-get -y upgrade && apt-get -y install sudo curl unattended-upgrades && -sL https://deb.nodesource.com/setup_4.x | sudo -E bash - && apt-get -f -y install nodejs

COPY ./overlay-common /dj
RUN cd /dj && npm install && ./node_modules/.bin/babel ./ -d bin --minified --ignore 'node_modules/'
ENV username=""
ENV DEBUG="true"
ENV compiled="true"

EXPOSE 80

CMD cd /dj/bin && node app.js
