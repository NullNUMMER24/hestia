FROM ubuntu:latest

# Install MariaDB server and client
RUN apt-get update && \
    apt-get install -y mariadb-server mariadb-client && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# Copy Hestia database dump into image
COPY hestia.sql /tmp/

# Initialize database with Hestia dump
RUN /usr/bin/mysqld_safe --datadir='/var/lib/mysql' --skip-networking & \
    sleep 10s && \
    mysql -u root < /tmp/hestia.sql && \
    killall mysqld

# Start MariaDB service
CMD ["mysqld_safe"]

# Expose default MariaDB port
EXPOSE 3306

