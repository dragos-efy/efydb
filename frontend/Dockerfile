FROM node:19 as build

ENV NODE_ENV=production 

WORKDIR /app

COPY package.json ./
RUN yarn install
COPY . ./
RUN yarn build

FROM node:19-alpine3.17

WORKDIR /app
COPY --from=build /app .

ENV HOST=0.0.0.0
EXPOSE 4173
CMD ["yarn", "preview", "--host", "0.0.0.0"]