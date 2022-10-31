FROM node:lts-alpine AS builder

WORKDIR /usr/src/app

COPY ./package*.json ./
RUN npm install

COPY . .

RUN npm install -g @angular/cli
RUN ng build --configuration production


FROM nginx:alpine
COPY --from=builder /usr/src/app/dist/front_f4 /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80