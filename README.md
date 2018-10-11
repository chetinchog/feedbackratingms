<a name="top"></a>
<h3>FeedbackRatingMS v0.14.4</h3>

# <a name='ms'></a> Microservicio de Valoración de Artículos

El Microservicio de Valoración de Artículos es el encargado de la administración de las valoraciones realizadas, por los usuarios, a los artículos y de su categorización según las reglas parametrizables.
Se guardan las valoraciones realizadas, se realiza el cálculo de la valoración promedio para cada artículo.
Se notifica cuando una valoración cumple alguna de las reglas, explicadas a continuación, y cuando se modifica la valoración promedio del artículo.

Las reglas parametrizables son:
- Si la valoración es mayor a un valor "x", se categoriza como "Buena Valoración"
- Si la valoración es menor a un valor "y", se categoriza como "Mala Valoración"
- Ambas se notifican por RabbitMQ

Para realizar estas acciones, el microservicio se comunica con los siguientes recursos:
- Auth: Para la administración de los parámetros de las reglas, se debe validar que el usuario sea Admin
- Catalog: Se realizan las validaciones de los artículos contenidos en las reseñas recibidas
- UserFeedback: No se comunica directamente, pero recibe las reseñas mediante RabbitMQ para su procesamiento
- RabbitMQ: Envía las notificaciones y alertas sobre las valoraciones buenas o malas y los cambios en los promedios

Y cuenta con las siguientes funcionalidades:
- [Reglas de Valoración](#rate-params)
    - [Asignar Parámetro General](#upsert-general-param)
    - [Buscar Parámetro General](#view-general-param)
    - [Asignar Parámetro a Artículo](#upsert-article-param)
    - [Buscar Parámetro a Artículo](#view-article-param)

- [Valoración](#rating)
    - [Buscar Valoración de Artículo](#view-article-rate)
    - [Buscar Historial de Artículo](#view-article-history)

- [RabbitMQ_GET](#rabbitmq_get)
    - [Buscar Reseña](#feedback-request)
    - [Buscar Validación de Artículo](#article-validation-response)
    - [Cerrar Sesión](#logout)

- [RabbitMQ_POST](#rabbitmq_post)
    - [Notificación de Valoración Alta](#high-rating-notify)
    - [Alerta de Valoración Baja](#low-rating-warning-alert)
    - [Notificación de cambio de Valoración de Artículo](#change-rating-notify)
    - [Solicitar Validación de Artículo](#article-validation-request)

# <a name='rate-params'></a> Reglas de Valoración

## <a name='upsert-general-param'></a> Asignar Parámetro General
[Inicio](#top)
<p>Administración de la categorización por valoración general.</p>
<p>Estos parámetros se utilizan si no existen los propios del artículo.</p>
<ul>
    <li>Si no existe lo crea o, por el contrario, lo actualiza.</li>
    <li>En el caso de que se asigne un "0", se anularía la regla y no se realizarían las  categorizaciones o notificaciones.</li>
</ul>
<h5>URL:</h5>
<p>POST /v1/rates/rules</p>

### Ejemplos
Body
```bash
{
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}"
}
```
Header de Autorización
```bash
Authorization=bearer {token}
```
### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}",
    "created": "{creation date}",
    "modified": "{modification date}"
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

## <a name='view-general-param'></a> Buscar Parámetro General
[Inicio](#top)
<p>Ver los parámetros generales de categorización por valoración.</p>
<h5>URL:</h5>
<p>GET /v1/rates/rules</p>

### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}",
    "created": "{creation date}",
    "modified": "{modification date}"
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

## <a name='upsert-article-param'></a> Asignar Parámetro a Artículo
[Inicio](#top)
<p>Administración de la categorización por valoración para un artículo específico.</p>
<p>Si se realizó la asignación de estos valores entonces los parámetros generales no se tienen en cuenta a la hora de categorizar.</p>
<ul>
    <li>Si no existe lo crea o, por el contrario, lo actualiza.</li>
    <li>En el caso de que se asigne un "0", se anularía la regla y no se realizarían las  categorizaciones o notificaciones.</li>
</ul>
<h5>URL:</h5>
<p>POST /v1/rates/:articleId/rules</p>

### Ejemplos
Body
```bash
{
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}"
}
```
Header de Autorización
```bash
Authorization=bearer {token}
```
### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "articleId": "{article's id}",
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}",
    "created": "{creation date}",
    "modified": "{modification date}"
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

## <a name='view-article-param'></a> Buscar Parámetro a Artículo
[Inicio](#top)
<p>Ver los parámetros para un artículo de su categorización por valoración.</p>
<h5>URL:</h5>
<p>GET /v1/rates/:articleId/rules</p>

### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "articleId": "{article's id}",
    "lowRate": "{bad rate's value}",
    "highRate": "{good rate's value}",
    "created": "{creation date}",
    "modified": "{modification date}"
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

# <a name='rating'></a> Valoración

## <a name='view-article-rate'></a> Buscar Valoración de Artículo
[Inicio](#top)
<p>Ver valoración de un artículo</p>
<ul>
    <li>Promedio de valoraciones hechas.</li>
    <li>Cantidad de valoraciones realizadas</li>
    <li>Clasificación del artículo según la categorización por valoración actual.</li>
</ul>
<h5>URL:</h5>
<p>GET /v1/rates/:articleId/</p>

### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "articleId": "{article's id}",
    "rate": "{article rate's value}",
    "feedAmount": "{amount of feedbacks made}",
    "badRate": "{is this category (boolean)}",
    "goodRate": "{is this category (boolean)}",
    "created": "{creation date}",
    "modified": "{modification date}"
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

## <a name='view-article-history'></a> Buscar Historial de Artículo
[Inicio](#top)
<p>Ver historial de un artículo</p>
<ul>
    <li>Todas las valoraciones realizadas.</li>
</ul>
<h5>URL:</h5>
<p>GET /v1/rates/:articlesd/history</p>

### Respuesta de éxito
Respuesta
```bash
# HTTP/1.1 200 OK
{
    "id": "{article's id}",
    "history": [
        {
            "rate": "{rate's value}",
            "userId": "{user's id}",
            "created": "{creation date}"
        }
    ]
}
```
### Respuesta de Error
401 Unauthorized
```bash
# HTTP/1.1 401 Unauthorized
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "path" : "{property name}",
    "message" : "{error cause}"
}
```
400 Bad Request
```bash
# HTTP/1.1 400 Bad Request
{
    "error" : "{error cause}"
}
```
500 Server Error
```bash
# HTTP/1.1 500 Server Error
{
    "error" : "{error cause}"
}
```

# <a name='rabbitmq_get'></a> RabbitMQ_GET

## <a name='feedback-request'></a> Buscar Reseña
[Inicio](#top)
<p>Escucha los mensajes de creación de Feedback para obtener las valoraciones.</p>

    DIRECT feedback/new-feedback

### Respuesta de éxito 
Mensaje
```bash
{
    "type": "new-feedback",
    "message": {
        "id" : "{feedback's id}"
        "idUser" : "{user's id}",
        "idProduct" : "{product's id}",
        "rate" : "{feedback's rate}",
        "created" : "{creation date}",
        "modified" : "{modification date}"
    }
}
```

## <a name='article-validation-response'></a> Buscar Validación de Artículo
[Inicio](#top)
<p>Escucha los mensajes de validación de productos solicitados.</p>

    DIRECT rates/article-exist

### Respuesta de éxito
Mensaje
```bash
{
  "type": "article-exist",
  "message" : {
      "articleId": "{articleId}",
      "valid": True|False
  }
}
```

## <a name='logout'></a> Cerrar Sesión
[Inicio](#top)
<p>Escucha de mensajes logout desde auth. Invalida sesiones en cache.</p>

	FANOUT auth/logout

### Respuesta de éxito
Mensaje
```bash
{
  "type": "article-exist",
  "message" : "tokenId"
}
```


# <a name='rabbitmq_post'></a> RabbitMQ_POST

## <a name='high-rating-notify'></a> Notificación de Valoración Alta
[Inicio](#top)
<p>Si una reseña supera la regla de una buena valoración, se notifica.</p>

    TOPIC rates/high-rate

### Ejemplo
Mensaje
```bash
{
   "type": "high-rate",
   "queue": "rates"
   "message": {
        "feedbackId" : "{feedback's id}",
        "notificationDate": "{notification's date}"
    }
}
```

## <a name='low-rating-warning-alert'></a> Alerta de Valoración Baja
[Inicio](#top)
<p>Si una reseña supera la regla de una mala valoración, se genera una alerta.</p>

    TOPIC rates/low-rate

### Ejemplo
Mensaje
```bash
{
   "type": "low-rate",
   "queue": "rates"
   "message": {
        "feedbackId": "{feedback's id}",
        "notificationDate": "{notification's date}"
    }
}
```

## <a name='change-rating-notify'></a> Notificación de cambio de Valoración de Artículo
[Inicio](#top)
<p>Se notifica cada vez que cambia el promedio de la valoración de un artículo.</p>

    TOPIC rates/article-change-rate

### Ejemplo
Mensaje
```bash
{
   "type": "article-change-rate",
   "queue": "rates"
   "message": {
        "articleId": "{article's id}",
        "newRate": "{article rate's value}",
        "notificationDate": "{notification's date}"
    }
}
```

## <a name='article-validation-request'></a> Solicitar Validación de Artículo
[Inicio](#top)
<p>Antes de iniciar las operaciones se validan los artículos contra el catalogo para verificar su existencia.</p>

    DIRECT catalog/article-exist

### Ejemplo
Mensaje
```bash
{
    "type": "article-exist",
    "queue": "rates",
    "message": {
        "articleId": "{articleId}"
    }
}