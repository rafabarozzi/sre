## Go + Docker + k8s

1) Desenvolver app web em Go Simples.
2) Build com Docker
3) Push para o docker hub
4) Deploy no k8s
5) Exposição da app usando ingress.

# Criar a aplicação go 

``` 
1) Criar o arquivo main.go

2) Ir até o diretório

#Iniciar o modulo
go mod init myproject

```

## Testar a aplicação

```
#BUILD
go build

#Executar
go run main.go
```

### Fazer build e Push

``` 
#Build 
docker build -t teste:latest .

#Login
docker login

#Push
docker push rafabarozzi/teste:latest
```

### Deploy no kubernetes

``` 
kubectl apply -f kube-manifests/
```

### Obter o IP publico da VM

```
curl ifconfig.me
```

### Comandos Docker

```
#Listar containers executando
docker ps

#Parar container 
docker stop container

#Listar containers parados
docker ps -a 

#Remover os containers
docker rm container

#Listar as imagens
docker images

#Remover uma imagem
docker rmi image
```

# Comandos kubernetes

``` 
#Listar os pods
kubectl get pods -n namespace

#Listar os services
kubectl get svc -n namespace

#Listar o ingress
kubectl get ingress -n namespace

#Listar os secrets
kubectl get secret -n namespace

#Remover o namespace
kubectl delete namespace

#Remover os pods (para outros objetos só mudar o nome)
kubectl delete pod <podname> -n namespace

#Ver o log dos pods
kubectl logs -f <podname> -n namespace

#Ver a descrição dos pods (para outros objetos só mudar o nome)
kubectl describe pod <podname> -n namespace
```