FROM node:lts-alpine

ENV NODE_ENV=production

WORKDIR /app

COPY package.json yarn.lock ./

RUN --mount=type=cache,target=/root/.cache/yarn \
    yarn install --frozen-lockfile --production

COPY . .

RUN yarn build

EXPOSE 3000
CMD ["node", "build"]
