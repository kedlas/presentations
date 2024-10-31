# "Asynchronní integrace aplikací - RabbitMQ"
## Tomáš Sedláček

### Symfony sraz Hradec Králové, březen 2017
Prezentace dostupná na adrese: https://prezi.com/p0gzzfnolbr4/asynchronni-integrace-aplikaci-rabbitmq/
Všechny zdrojové kódy jsou v tomto repozitáři. 

### RabbitMQ prostředí:
```
Spuštění docker kontejneru:
docker rm my-rabbit
docker run -d --hostname my-rabbit -p 15672:15672 -p 5672:5672 --name my-rabbit rabbitmq:3-management-alpine

Zkontrolujeme, zda instance RabbitMQ beží:
docker exec -ti my-rabbit rabbitmqctl status

Zjistíme jakou má kontejner IP adresu:
docker inspect my-rabbit 
```

RabbitMQ management frontend by měl být dostupný na adrese:
http://host:15672/ (e.g.: http://localhost:15672/#/)

### Demo Sorter/Resequencer

Na localhost:
```
$ git clone https://github.com/kedlas/2017-rabbitmq-prednaska.git
$ cd 2017-rabbitmq-prednaska
$ npm install
$ cd src
$ node consumer.js simple|sorter|resequencer

$ node publisher.js
```

Přes docker:

Nahraďte "localhost" za IP adresu rabbitmq kontejneru (např --build-arg rabbithost=172.17.0.2) 
```
# docker build --build-arg rabbithost=localhost -t my_app .
```

Nastartujeme instance aplikací. Právě 1 pro publishera a alespoň 1 pro consumery (níže např simple consumer)
```
$ docker rm my_consumer_app
$ docker run -ti --name=my_simple_consumer_app my_app
$ docker exec -ti my_simple_consumer_app bash
  $ node consumer.js simple

$ docker rm my_publisher_app
$ docker run -ti --name=my_publisher_app my_app
$ docker exec -ti my_publisher_app bash
  $ node publisher.js
```
