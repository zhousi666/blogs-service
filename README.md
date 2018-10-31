# blogs-service


### What is the blogs service


blogs-service is the server of blogs，which stores and processes user data for the client.
Developed using golang, based on docker deployment.


### Dependencies


- install docker

- install docker-compose

- mysql ( You can see the config of db in [config/config.tol](https://github.com/zhousi666/blogs-service/blob/master/blogs-service/config/config.toml) ).


### Config 

 The config of server is made up of  [config/config.tol](https://github.com/zhousi666/blogs-service/blob/master/blogs-service/config/config.toml)  and [docker-compose.yml](https://github.com/zhousi666/blogs-service/blob/master/docker-compose.yml).


### Installation

    // get code
    git clone https://github.com/zhousi666/blogs-service.git

    // build blogs
    cd blogs-service
    sudo make

    // start blogs service
    sudo docker-compose -f docker-compose.yml up -d 
    or : sudo docker-compose -f docker-compose.yml up

