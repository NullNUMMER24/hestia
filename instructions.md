# Backend
### How to build the image
The Image can be build with the docker build command
```Bash
sudo docker build -t hestia-backend .
```
### How to run the container
The Container can be started with 
```Bash
sudo docker run -d -it -p 5432:5432 --name hestia-backend -e POSTGRES_PASSWORD=123 hestia-backend
```
### Connect to DB from localhost
```Bash
psql -h localhost -U postgres
```
# Frontend
### Build the image
```Bash
sudo docker build -t hestia-frontend .
```
# Docker Network
### Create a hestia-network
```Bash
sudo docker network create hestia-network
```
### Start Backend with the hestia-network
```Bash
sudo docker run --name hestia-backend --hostname hestia-backend --network hestia-network -p 5432:5432 -e POSTGRES_PASSWORD=123 -d hestia-backend
```
### Start Frontend with the hestia-network
```Bash
sudo docker run --network hestia-network -p 3000:3000 hestia-frontend
```