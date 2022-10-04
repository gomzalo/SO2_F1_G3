# **Manual de usuario**

- [**Manual de usuario**](#manual-de-usuario)
- [Fase 1](#fase-1)
  - [Menu Principal Meta OS System](#menu-principal-meta-os-system)
  - [Ingresar Datos](#ingresar-datos)
  - [Menu Opciones](#menu-opciones)
  - [Ingreso de Datos](#ingreso-de-datos)
  - [Resultados IOTOP](#resultados-iotop)
  - [Resultados TOP](#resultados-top)
- [Fase 2](#fase-2)
  - [SYSCALL](#syscall)
  - [Reporte JSON](#reporte-json)
- [Fase 3](#fase-3)

# Fase 1

## Menu Principal Meta OS System

En este apartado puedes seleccionar cada una de las caracteristicas necesarias a utilizar durante esta fase.

![Menu Principal](img/F1/menu.png)
## Ingresar Datos
![Menu Principal](img/F1/menu2.png)
## Menu Opciones
![Menu Principal](img/F1/top.png)
## Ingreso de Datos
![Menu Principal](img/F1/nombre.png)
## Resultados IOTOP
![Menu Principal](img/F1/iotop.png)
## Resultados TOP
![Menu Principal](img/F1/rsultTop.png)

# Fase 2

## SYSCALL

1. Ingresar a la aplicación y elegir la opción *1. Nueva ejecución*, luego ingresar a la opción *3. STRACE* y seleccionar la opción *1. Ingresar nombre*.
![Menu Strace](img/F2/MenuStrace.png)

2. Ingresar un comando, como 
   ```sh
   echo hello
   ```
   y el sistema debera de mostrar los procesos que se ejecutan en el sistema.
   ![Strace](img/F2/Strace.png)

## Reporte JSON

1. Ingresar a la aplicación y elegir la opción *2. Reporte*.
   ![Menu Reporte](img/F2/MenuReportes.png)
2. Seleccionar la opción *1. Bitácora*; el sistema debera de indicar que se genero el reporte.
   ![Reporte Generado](img/F2/ReporteGenerado.png)
3. Para visualizar el reporte, nos salimos de la aplicación y nos dirigimos a la carpeta del proyecto, se ingresa el comando:
   ```sh
   cat bitacora.json
   ```
   siendo *bitacora.json* el nombre del archivo que se genero.
   ![Reporte JSON](img/F2/ReporteJSON.png)


# Fase 3

## Funcion MEMSIM

1. Despues de colocar el nombre se debe seleccionar la opcion *4. MEMSIM*.
   [Menu Reporte](img/F3/opMemsim.png)
2. Se solicitaran dos parametros, el primero hace referencia a la cantidad de ciclos que se ejecutara el proceso        mientras que el segundo parametro es una lista en la que cada unidad esta delimitada por comas.
   ![Reporte Generado](img/F3/secuencia.png)
3. Se presentara el detalle de como fueron operados los procesos, su orden de inicio y en que momento terminaron
   ![Reporte Generado](img/F3/final.png)
