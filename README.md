# miprimerservidor
servidor tcp
Para ver la funcionalidad de este software se deve ejecutar la carpeta servidor (go run .) y el cliente, client.go
Se puede ejecutar cuantos clientes se requieran para hacer para la prueba.
Primero se inicia el servidor automaticamente cuando se ejecuta la carpeta servidor
Luego se envia desde el cliente los comandos necesarios para enviar los archivos al respectivo canal de la siguiente manera:
  ./holamundo.txt.fin 1.fin
  ./ (es la ruta del archivo)
  holamundo.txt (es el nombre del archivo)
  .fin (es para aclarar que finaliza el nombre del archivo y la ruta)
  " " un espacio para separar el canal a enviar del nombre del archivo a enviar
  1 (el numero de canal a enviar)
  .fin para aclarar que termina el numero del canal
  
  Los nombres de los canales se asignan automaticamente desde el 0 hasta el numero de cliente que se ejecute 
