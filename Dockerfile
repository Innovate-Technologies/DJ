FROM multiarch/ubuntu-core:armhf-xenial
#FROM multiarch/ubuntu-core:amd64-xenial # arch=amd64

RUN curl -sL https://deb.nodesource.com/setup_4.x | sudo -E bash - && apt-get -y upgrade && apt-get -f -y install unattended-upgrades nodejs

COPY ./overlay-common /dj
RUN cd /dj && npm install && ./node_modules/.bin/babel ./ -d bin --minified --ignore 'node_modules/'
ENV username=""
ENV DEBUG="true"
ENV compiled="true"

EXPOSE 80

CMD cd /dj/bin && node app.js
