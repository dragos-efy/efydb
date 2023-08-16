FROM node:lts-alpine
ENV NODE_ENV=production
WORKDIR /app
COPY package*.json ./
RUN yarn

COPY . .
RUN yarn build

ENV HOST=0.0.0.0
EXPOSE 4173
CMD ["yarn", "preview", "--host", "0.0.0.0"]