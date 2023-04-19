# build
FROM node:lts-slim as build

WORKDIR /src

COPY package.json ./
RUN npm install

COPY . .
RUN npm run build

# run
FROM node:lts-slim as run

WORKDIR /app

COPY --from=build /src/package.json ./package.json
RUN npm install --omit=dev --ignore-scripts

COPY --from=build /src/build ./build

EXPOSE 3000
ENTRYPOINT [ "node", "build" ]
