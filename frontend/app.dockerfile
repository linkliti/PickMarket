FROM node:20-alpine
# Dependencies
WORKDIR /application
COPY ./frontend/package.json .
RUN npm install

# Build
COPY ./frontend .
RUN npm run build
CMD [ "npm", "run", "preview" ]