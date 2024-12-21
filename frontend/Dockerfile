FROM node:22

WORKDIR /app

COPY frontend/ ./

RUN npm install -g serve

EXPOSE 14000

CMD ["serve", "-s", ".", "-l", "14000"]
