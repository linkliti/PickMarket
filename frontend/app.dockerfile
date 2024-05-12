# Dependencies
FROM node:20-alpine AS build
WORKDIR /application
COPY ./package.json .
RUN npm install --verbose

# Build
COPY . .
RUN npm run build --verbose

# Server
# FROM node:20-alpine
# RUN npm install -g vite --verbose
# RUN npm install -g @vitejs/plugin-react-swc --verbose
# WORKDIR /application
# COPY ./vite.config.ts .
# COPY --from=build /application/dist ./dist
CMD [ "npx" , "vite", "preview"]