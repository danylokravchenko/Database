docker build -t db .

docker run -e ADVERTISE_ADDR=172.17.0.2 -e CLUSTER_MEMBERS=3 -p 8080:8080 db
docker run -e ADVERTISE_ADDR=172.17.0.3 -e CLUSTER_ADDR=172.17.0.2 -e CLUSTER_MEMBERS=3 -p 8081:8080 db
docker run -e ADVERTISE_ADDR=172.17.0.4 -e CLUSTER_ADDR=172.17.0.3 -e CLUSTER_MEMBERS=3 -p 8082:8080 db
docker run -e ADVERTISE_ADDR=172.17.0.5 -e CLUSTER_ADDR=172.17.0.4 -p 8083:8080 db

or (but I haven't really found error in ports in docker-compose implementation)
docker-compose up -d --build
