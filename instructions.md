# How to build the image
The Image can be build with the docker build command
```Bash
sudo docker build -t db .
```
# How to run the container
The Container can be started with 
```Bash
sudo docker run -d -it -p -5432:5432 -name db -e POSTGRES_PASSWORD=123 db
```
