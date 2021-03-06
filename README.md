# Galera Cluster docker based 
support mariadb10.1

#### Requirement 
* docker 
* go1.10
* `go get -u github.com/golang/dep/cmd/dep` and `dep ensure -v`
* build => `go build`
* Please refer to docker-entrypoint.sh for further detail
* Note that probably it takes a long time when run for the first time . because docker will pull image from the cloud

#### Short explanation 
* cmd - contain command line arguments
* var - contain environment variable. You should modify this env based on your needs 

#### Command 
##### build docker image (optional)
example : 
```
./galeras build-docker-image "pathto/galeras/docker_mariadb_10.1/" galera-mariadb:10.1
```

##### run node and start cluster
use `-command mysqld,init` to start new cluster. otherwise `-command mysqld,join`
Node 1 
```
./galeras node run --name=galera-node-1 --host 172.25.0.2 --env-file var/env_1.env --net galeranet --ip 172.25.0.2 --add-host galera-node-2:172.25.0.3,galera-node-3:172.25.0.4 --port 3306,4444,4567,4568 --image yogyrahmawan/galera-mariadb:10.1 --additional-command mysqld,init
```
Node 2
```
./galeras node run --name=galera-node-2 --host 172.25.0.3 --env-file var/env_2.env --net galeranet --ip 172.25.0.3 --add-host galera-node-1:172.25.0.2,galera-node-3:172.25.0.4 --port 3306,4444,4567,4568 --image yogyrahmawan/galera-mariadb:10.1 --additional-command mysqld,join
```
Node 3
```
./galeras node run --name=galera-node-3 --host 172.25.0.4 --env-file var/env_3.env --net galeranet --ip 172.25.0.4 --add-host galera-node-1:172.25.0.2,galera-node-2:172.25.0.3 --port 3306,4444,4567,4568 --image yogyrahmawan/galera-mariadb:10.1 --additional-command mysqld,join
```

#### monitor nodes 
```
./galeras monitor --username root --password root --node galera-node-1
```

#### remove nodes 
```
./galeras node rm --name galera-node-1 --name galera-node-2
```

#### runtest nodes
It will start 3 nodes . By Default, this command will used image from `yogyrahmawan/galera-mariadb:10.1`. You can just run this below command or alternatively you can pull the image first by executing `docker pull yogyrahmawan/galera-mariadb:10.1`
Command :
```
./galeras runtest
```

#### run websocket server
Requirements : run galera node using command `./galeras runtest` Then you can run below command. The `path` is path to public html
```
./galeras runwsserver --path path_to_public_html_file
```
example :
```
./galeras runwsserver --path "/Volumes/hdd storage/Users/yogydev/Development/projects/golang/src/github.com/yogyrahmawan/galeras/public"
```
By default it will check using `galera-node-1` so make sure galera-node-1 is up

#### open client 
you can open initial web client by open `localhost:3000` on your browser. Note : you need to connect to the internet to run this client since the css and other packages are downloaded from the cloud.

