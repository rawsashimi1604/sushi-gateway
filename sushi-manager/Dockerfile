## Dockerfile
################################
## BUILD ENVIRONMENT ###########
################################

FROM node:20-alpine3.20 As build
WORKDIR /usr/src/app
COPY package*.json package-lock.json ./
RUN npm install
COPY ./ ./
RUN npm run build

################################
#### PRODUCTION ENVIRONMENT ####
################################

FROM nginx:stable-alpine as production
COPY --from=build /usr/src/app/nginx /etc/nginx/conf.d
COPY --from=build /usr/src/app/dist /usr/share/nginx/html
EXPOSE 5173
COPY env.sh /docker-entrypoint.d/env.sh
RUN chmod +x /docker-entrypoint.d/env.sh

# ENTRYPOINT ["nginx", "-g", "daemon off;"] 