FROM node
WORKDIR /juegos
COPY package*.json ./
RUN npm install
COPY . .
EXPOSE 4505
CMD [ "node", "server.js" ]