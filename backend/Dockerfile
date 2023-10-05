# Use a PostgreSQL base image
FROM postgres:latest

# Set the database name
ENV DB_NAME Hestia

# Copy the SQL script into the image
COPY init.sql /docker-entrypoint-initdb.d/

# Expose PostgreSQL port
EXPOSE 5432

# Start the PostgreSQL server
CMD ["postgres"]


