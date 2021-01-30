# coffeebeans-people-backend

#Installation steps

Install golang
mkdir -p $HOME/go/{bin,src}
Set following in .bash_profile export GOPATH=$HOME/go export PATH=$PATH:$GOPATH/bin
Clone coffeebeans-people-backend
Install mongo (or use the docker-compose file) 
In project root run "go mod tidy"
In project root run "go mod vendor"
In project root run "go run main/*.go -file=dev_local.json"


Deployment:

For setting up mongo run :
docker run --name mongodb -v /tmp/data/:/data/db -d mongo 

For setting up go and building it run :
docker build -t go .

For running your go bin run : 
docker run --name go -p 9000:9000 -d go

