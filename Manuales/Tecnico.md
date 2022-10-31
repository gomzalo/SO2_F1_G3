# **Manual técnico**

- [**Manual técnico**](#manual-técnico)
- [Fase 1](#fase-1)
  - [Requisitos del sistema](#requisitos-del-sistema)
  - [Instalar GO](#instalar-go)
    - [Descargar Go](#descargar-go)
    - [Configurar Variable de entorno](#configurar-variable-de-entorno)
    - [Configurar entorno Go](#configurar-entorno-go)
    - [Compilar archivo Go](#compilar-archivo-go)
  - [Instalar GCC](#instalar-gcc)
    - [Instalar cabeceras linux](#instalar-cabeceras-linux)
    - [Instalar dependencias](#instalar-dependencias)
  - [Kernel](#kernel)
- [Fase 2](#fase-2)
  - [Implementacion de SYSCALL - Strace](#implementacion-de-syscall---strace)
  - [Reportes](#reportes)
  - [Actualizacion de Consola](#actualizacion-de-consola)
- [Fase 3](#fase-3)
  - [Implementacion de MEMSIM](#implementacion-de-memsim)
  - [Funcion MEMSIM](#funcion-memsim)
  - [Tiempo e Impresion](#tiempo-e-impresion)
- [Fase 4](#fase-4)
  - [Arquitectura](#arquitectura)
  - [Base de Datos MongoDB - NOSQL](#base-de-datos-mongodb---nosql)
  - [Usuarios](#usuarios)
  - [memsim](#memsim)
  - [Despliegue](#despliegue)
  - [Despliegue](#despliegue-1)
  - [Dockerfile](#dockerfile)
  - [Comandos](#comandos)
    - [Docker](#docker)
    - [Kubernetes](#kubernetes)
  - [Pasos para realizar un rollout y un rollback](#pasos-para-realizar-un-rollout-y-un-rollback)

# Fase 1

## Requisitos del sistema

- Sistema GNU/Linux
- 1 Gb o más de RAM
- 10 Gb o más de Almacenamiento
- Compilador GCC instalado
- Compilador GoLang instalado
- Headers genericos del Kernel de Linux

## Instalar GO

### Descargar Go

Acceder al servidor

```sh
ssh usuario@server_ip
```

Moverse a la carpeta de usuario

```sh
cd ~
```

Descargar los archivos necesarios para compilar Go

```sh
curl -OL https://golang.org/dl/go1.19.linux-amd64.tar.gz
```

Validar el archivo descargado


```sh
sha256sum go1.16.7.linux-amd64.tar.gz
```

Descomprimir el archivo descargado y copiarlo a la carpeta `/usr/local`

```sh
sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
```

### Configurar Variable de entorno

Con el editor de preferencia, abrir el archivo `profile` y agregar la ruta de los binarios de go

```sh
sudo nano ~/.profile
```

```sh
export PATH=$PATH:/usr/local/go/bin
```

### Configurar entorno Go

Crear un archivo mod para importar paquetes instalados

```sh
go mod init proyecto
```

### Compilar archivo Go

```sh
go build .
```

## Instalar GCC

### Instalar cabeceras linux

```sh
sudo apt install linux-headers-$(uname -r)
```

### Instalar dependencias

```sh
sudo apt install build-essential
```

## Kernel

Compilar kernel de archivo `.c`

```sh
make
```

Instalar módulo Kernel

```sh
sudo insmod [nombreModulo].ko
```

Mostrar mensajes de consola al instalar el módulo Kernel

```sh
sudo dmesg
```

Desinstalar módulo Kernel

```sh
sudo rmmod [nombreModulo]
```

Revisar archivos escrito desde Kernel instalado

```sh
cat /proc/[nombreModulo]
```

# Fase 2

## Implementacion de SYSCALL - Strace

![Menu Principal](img/F2/codeStrace.png)

## Reportes 

![Menu Principal](img/F2/codeExportarReporte.png)

## Actualizacion de Consola
![Menu Principal](img/F2/codeActualizacionMensajes.png)

# Fase 3

## Implementacion de MEMSIM
![Menu Principal](img/F3/memsim.png)

## Funcion MEMSIM
![Menu Principal](img/F3/funcion.png)


## Tiempo e Impresion
![Menu Principal](img/F3/work.png)

# Fase 4

## Arquitectura
![Menu Principal](img/F4/arquitectura.png)

## Base de Datos MongoDB - NOSQL
![Menu Principal](img/F4/mongo.png)

## Usuarios
![Menu Principal](img/F4/usuariosdb.png)

## memsim
![Menu Principal](img/F4/usuariosdb.png)

## Despliegue
![Menu Principal](img/F4/deploymentk8s.png)

## Despliegue
![Menu Principal](img/F4/deploymentk8s.png)

## Dockerfile
![Menu Principal](img/F4/dockerfront.png)

[Documentación de APIRest](https://documenter.getpostman.com/view/5658161/2s84LGYaxu)

## Comandos

### Docker

- #### Crear imagen

  ```sh
  docker build -t [user_name]/[image_name]:[version] .
  ```
- #### Correr imagen

  ```sh
  docker run -p [host_port]:[container_port] [user_name]/[image_name]:[version]
  ```

- #### Ver imagenes

  ```sh
  docker images
  ```

- #### Ver contenedores

  ```sh
  docker ps
  ```

- #### Subir imagen a DockerHub

  ```sh
  docker push [user_name]/[image_name]:[version]
  ```

- #### Descargar imagen de DockerHub

  ```sh
  docker pull [user_name]/[image_name]:[version]
  ```

### Kubernetes

- #### Configurar kubeconfig utilizando ***doctl***

  - Conectar doctl a DigitalOcean
  
  ```sh
  doctl auth init
  ``` 

  - Obtener kubeconfig de DigitalOcean
  
  ```sh
  doctl kubernetes cluster kubeconfig save [nombre_cluster]
  ```

- #### Aplicar configuracion de Kubernetes
  <div id='apply'/>

  ```sh
  kubectl apply -f [archivo.yaml]
  ```

- #### Ver pods

  ```sh
  kubectl get pods
  ```
- #### Obtener servicios
  Agregar bandera `-w` para ver cambios en tiempo real.

  ```sh
  kubectl get svc
  ```

- #### Obtener deployments
  Agregar bandera `-w` para ver cambios en tiempo real.

  ```sh
  kubectl get deploy
  ```
- #### Crear servicio desde consola
  
  ```sh	
  kubectl expose deployment [deployment_name] --type=LoadBalancer --name=[service_name]  --port=[port]
  ```
- #### Eliminar un servicio

  ```sh
  kubectl delete service [service_name]
  ```
- #### Eliminar un deployment

  ```sh
  kubectl delete deployment [deployment_name]
  ```

- #### Obtener replicasets
  <div id='get-rs'/>
  Agregar bandera `-w` para ver cambios en tiempo real.
  
  ```sh
  kubectl get rs
  ```

- #### Realizar un rollout
  <div id='rollout-status'/>

  ```sh
  kubectl rollout undo [deployment_name]
  ```

- #### Verficar estado de rollout
  <div id='rollout-undo'/>

  ```sh
  kubectl rollout status deployment/[deployment_name]
  ```

- #### Actualizar imagen de docker en deployment.yaml
  <div id='update-tag'/>

  ```sh
  kubectl set image deployment/[deployment_name] [container_name]=[user_name]/[image_name]:[version]
  ```

## Pasos para realizar un rollout y un rollback

1. Actualizar el codigo y subir los cambios con un nuevo tag1 a DockerHub.
   
    ```sh
    docker build -t [user_name]/[image_name]:[version] .
    docker push [user_name]/[image_name]:[version]
    ```

2. Actualizar imagen de docker en el deployment deseado

    ```sh
    kubectl set image deployment/[deployment_name] [container_name]=[user_name]/[image_name]:[version]
    ```

3. Verificar estado del rollout

    ```sh
    kubectl rollout status deployment/[deployment_name]
    ```
  
4. Realizar un rollback.

    ```sh
    kubectl rollout undo deployment/[deployment_name]
    ```

    Se puede [verificar el estado del rollout](#rollout-status).


* [Abreviaciones de comandos en **Kubectl**](https://gist.github.com/rosskukulinski/640e34e335c505a260665e1dcce2bb46#file-k8s-resources-md)