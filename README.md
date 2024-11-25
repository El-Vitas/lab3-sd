## Nombres y rol:
+ Sebastián Arrieta, 202173511-9
+ Jonathan Olivares, 202073096-2

## Sobre el uso:
- El servicio de los supervisores, tanto el 1 como el 2 requieren que se le entregue datos por consola, al igual que jayce.
- La distribución de los servicios es la siguiente:
  + En la maquina virtual 1 (MV1): Broker
  + En la maquina virtual 2 (MV2): Hex server 2 y Supervisor hex 2
  + En la maquina virtual 3 (MV3): Hex server 3 y Supervisor hex 1
  + En la maquina virtual 4 (MV4): Hex server 1 y Jayce
- Para los casos que requieren input seguir el mismo formato que se les pide en los print. En el caso de los supervisores las opciones se acceden escribiendo la opcion tal cuál es. Por ejemplo: 'AgregarProducto'
  
## Para la ejecución
- Los docker utilizan docker compose up
- Estos docker son ejecutados mediante un makefile.
- Para la ejecución del makefile se puede usar el comando `make run` en cada maquina virtual.
- **El orden de ejecución de las maquinas virtuales es:  MV1 -> MV2 -> MV3 -> MV4.**
- Antes de comenzar a tener interacción con el programa se debe esperar que estén todos los servicios corriendo de acuerdo al comando `make run`.
- MV2, MV3, MV4 poseen servicios que necesitan input. Al necesitar un input usando docker compose se hace un attach a ese servicio especifico por lo que no se muestra el log del otro servicio. 
- Al hacer attach no se muestra el primer print, de recomendación ingresar cualquier letra para que aparezca el resultado (Lo más seguro es que dirá algo como comando incorrecto o producto no encontrado. De igual forma se dejan los Print):
  - Jayce:
    + Ingresa una región y un producto:
    + \<Region> \<Producto>
  
  - Supervisores:
    + Supervisor numero: Ingrese un comando (AgregarProducto, RenombrarProducto, ActualizarValor, BorrarProducto):
  
- Para finalizar el programa se debe hacer manualmente, mediante ctrl+c de preferencia.

- **Para lograr ejecutar los servicios se debe usar sudo al usar los makefile**
