FROM node:20
WORKDIR /usr/src/app
COPY package*.json ./
RUN npm install
COPY . .
ENV NODE_ENV=production
ENV PORT=5173
EXPOSE 5173
CMD ["npm", "run", "build"]
