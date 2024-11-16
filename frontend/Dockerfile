# Use the official Node.js 18 image as the base image
FROM node:18-alpine AS base

# Set the working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code to the container
COPY . .

# Build the Next.js application
RUN npm run build

# Use a lightweight production image to serve the application
FROM node:18-alpine AS production

# Set the working directory for production
WORKDIR /app

# Copy only necessary files from the build stage
COPY --from=base /app/package*.json ./
COPY --from=base /app/.next ./.next
COPY --from=base /app/public ./public
COPY --from=base /app/node_modules ./node_modules

# Set environment variables
ENV NODE_ENV=production

# Expose the application port
EXPOSE 3000

# Start the Next.js application
CMD ["npm", "run", "start"]