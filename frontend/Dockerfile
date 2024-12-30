FROM node:18-alpine

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

COPY . .

# Set the environment variable for Vite
ENV NODE_ENV development

# Expose the development port (5173)
EXPOSE 5173

# Command to run the frontend in development mode
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]