# Stage 1
FROM node:22.12 as builder
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build

# Stage 2
FROM builder
RUN apt-get update && apt-get install -y curl
COPY --from=builder /app .
CMD ["npm", "run", "preview"]