sudo apt update

// Docker installation
for pkg in docker.io docker-doc docker-compose docker-compose-v2 podman-docker containerd runc; do sudo apt-get remove $pkg; done

// getting images 
sudo docker pull mongo
sudo docker pull golang

// running container

sudo docker run -d -p 27017:27017 --name mongoDB mongo
sudo docker run -d -p 8080:8080 --name golang-app docker-golang-simple-crud
