# short-url

#### Pasos para ejecutar el proyecto

1. Clonar repositorio
2. Ejecutar el siguiente comando **`make file`**
3. Conectarse a la DB y ejecutar el sql **`db/migration/create_short_url_table.sql`**


#### Curls

##### Crear URL corta
```
curl --location --request POST 'localhost:8081/api/short-url/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "url": "https://aws.random.cat/"
}'
```


##### Redireccionar a URL larga
```
curl --location --request GET 'localhost:8081/api/short-url/{short_url}'
```


##### Borrar URL corta
```
curl --location --request DELETE 'localhost:8081/api/short-url/{short_url}'
```


##### Explicación de la solución
      
- Crear URL: al crear la URL se almacenará la URL corta y original en la BD (postgres), de igual forma se almacenará en redis la URL original. La key utilizada será la URL corta y tendrá un ttl de 30 días.

- Redireccionar a URL original: Se obtiene de redis la URL original y se hace la redirección a dicha URL.

- Borrar URL: Se borra la key de redis y al consultarse una URL borrada se retornar un error 404.

##### Futuras mejoras

- Aumentar cobertura de tests.
- Agregar un proceso async que permita restaurar en redis URL cortas.
- Agregar migraciones automáticas de la BD.
- Documentar api (swagger)
