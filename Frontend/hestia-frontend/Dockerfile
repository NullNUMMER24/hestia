# Use an official Node.js runtime as a parent image
FROM node:14

# Set the working directory to /app
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

RUN npm install

# Copy the app source code to the container
COPY app.js /app
COPY views /app/views
COPY public /app/public

# Make port 3000 available to the world outside this container
EXPOSE 3000

# Run the app when the container launches
CMD node app.js
