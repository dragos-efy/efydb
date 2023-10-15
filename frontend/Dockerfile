FROM node:lts-alpine AS build

ARG PUBLIC_API_URL
ENV PUBLIC_API_URL $PUBLIC_API_URL

WORKDIR /app

COPY package.json yarn.lock ./

RUN yarn install --frozen-lockfile

COPY . .

RUN yarn build


FROM node:lts-alpine AS runner

WORKDIR /app

ENV NODE_ENV=production

COPY --from=build /app/package.json .

COPY --from=build /app/build .

RUN yarn install --frozen-lockfile --production

EXPOSE 3000

CMD ["node", "index.js"]


