
FROM node:18-alpine as build


WORKDIR /app


COPY package.json package-lock.json ./
RUN npm install


COPY . .
RUN npm run build


FROM nginx:stable

# Copy the Vite build output to Nginx's default HTML location
COPY --from=build /app/dist /usr/share/nginx/html


COPY ./nginx.conf /etc/nginx/nginx.conf


EXPOSE 5173

CMD ["nginx", "-g", "daemon off;"]
