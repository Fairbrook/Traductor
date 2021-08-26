# Proyecto de la clase de Seminario de traductores de lenjuage

## Módulo 1 analizador léxico

### Entrega 2: Analizador léxico

Tipos admitidos

| Símbolo             | Descirpción                         |
| ------------------- | ----------------------------------- |
| Identificador       | letra(letra\|digito)\*              |
| Entero              | digito+                             |
| Real                | entero.entero                       |
| Adicion             | + -                                 |
| Multiplicacion      | \* /                                |
| Asignacion          | =                                   |
| Relacional          | < > <= >= != ==                     |
| And                 | &&                                  |
| Or                  | \|\|                                |
| Not                 | !                                   |
| Parentesis          | (,)                                 |
| Llave               | {,}                                 |
| Punto y coma        | ;                                   |
| Palabras reservadas | if, while, return, else, int, float |

Esta priemra entrega realiza el analizador léxico mediante una pseudo tabla de estados, se le llama seudo porque engloba los casos de múltiples caracteres en funciones como _isDigit_ y _isAlpha_ para determinar si un carcter es numérico o una letra

Esta tabla de estados es utilizada por un autómata que analiza la cadena de entrada hasta llegar a un caracter para el cual no se tiene un estado siguiente, una vez detenido regresa al estado 0 y reinicia el proceso

## Cómo correr el proyecto

Para el sistema operativo Windows 60bits puede ejecutar el archivo **.exe** que se encuentra en los release del repositorio, si está en otro sistema operativo o no funciona el ejecutable deberá compilarlo para la plataforma
Requerira un entorno de [golang](https://golang.org/)
instalar el paquete [go-asilectron](https://github.com/asticode/go-astilectron) y su respectivo [bundler](https://github.com/asticode/go-astilectron-bundler)

## Cómo utilizar el proyecto

Escriba el texto de entrada en el area de la izquierda y espera a que aparezca el resultado en la parte derecha
![pantalla](https://i.ibb.co/6RTWGNj/imagen.png)

## Tecnologías utilizadas

| Tecnologia                                                  | Utilizacion                                                                                                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------- |
| [golang](https://golang.org/)                               | Funcionalidad principal (backend)                                                                                                               |
| [electron](https://www.electronjs.org/)                     | Desarrollo de aplicaciones nativas con ui desarrollado en html, css y js                                                                        |
| [asilectron](https://github.com/asticode/astilectron)       | Proporciona sockets TCP para la comunicación entre Electron y cualquier lenguaje de programación con el fin de escuchar los eventos de electron |
| [go-asilectron](https://github.com/asticode/go-astilectron) | Api de asilectron desarrollada para go                                                                                                          |
| [quill](https://quilljs.com/)                               | Como editor de texto para el texto de entrada                                                                                                   |
| [AlpineJs](https://alpinejs.dev/)                           | Librería de JS para facilitar el comportamiento reactivo del UI                                                                                 |
