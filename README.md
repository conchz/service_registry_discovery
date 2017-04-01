# service_registry_discovery

#### REFERENCE_DOCUMENT
1. [使用Docker、Registrator、Consul、Consul Template和Nginx实现高可扩展的Web框架 - DockOne.io](http://dockone.io/article/272)
2. [Consul + Consul-Template with Docker-Compose](http://blog.neilni.com/2015/09/14/consul-and-consul-template-with-docker-compose/)

#### Install Consul
``` docker run -d --name=consul1 --hostname=node1 -p 8400:8400 -p 8500:8500 -p 8600:53/udp progrium/consul -server -bootstrap -advertise $YOUR_IP -ui-dir /ui ```

#### Install Registrator
``` docker run -d --name=registrator --net=host --hostname=$YOUR_IP -v /var/run/docker.sock:/tmp/docker.sock gliderlabs/registrator consul://$YOUR_IP:8500 ```

#### Install DR CoN
```
    cd nginx/
    docker build -t drcon .
    docker run -d -e "CONSUL=$YOUR_IP:8500" -e "SERVICE=simple" -p 80:80 drcon
```
