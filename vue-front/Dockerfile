FROM node:18 as build
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM alpine AS export
WORKDIR /app
COPY --from=build /app/dist ./dist
