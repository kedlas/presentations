FROM node:7

COPY . /usr/src/app

WORKDIR /usr/src/app

RUN npm install

ARG rabbithost=localhost
ENV RABBIT_HOST ${rabbithost}

CMD ["sh", "-c", "tail -f /dev/null"]