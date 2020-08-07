FROM node:14.7-alpine

RUN mkdir -p /client/src

WORKDIR /client/src

COPY package.json .

RUN npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]