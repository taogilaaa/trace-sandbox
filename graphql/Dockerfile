# Dependencies Image
FROM node:12-alpine AS dependencies
WORKDIR /usr/src/app

COPY package.json yarn.lock ./

# install dependencies
RUN yarn install --frozen-lockfile

# Builder Image
FROM dependencies AS builder
WORKDIR /usr/src/app

COPY . .

# run build
RUN yarn build

# prune uneeded dependencies inside node_modules
RUN yarn install --production --ignore-scripts --prefer-offline

# Main Image
FROM node:12-alpine
WORKDIR /usr/src/app

ENV NODE_ENV=production

COPY --from=builder /usr/src/app/node_modules node_modules
COPY --from=builder /usr/src/app/dist ./dist
COPY --from=builder /usr/src/app/package.json ./

EXPOSE 50040

CMD ["node", "dist/main.js"]
