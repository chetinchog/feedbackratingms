<a name="top"></a>
# FeedbackRating Service vv0.14.7

FeedbackRatingMS Manager

- [RabbitMQ_GET](#rabbitmq_get)
	- [Buscar Validación de Artículo](#buscar-validación-de-artículo)
	- [Buscar Reseña](#buscar-reseña)
	- [Logout de Usuarios](#logout-de-usuarios)
	
- [RabbitMQ_POST](#rabbitmq_post)
	- [Product Validation](#product-validation)
	- [Notificación de cambio de Valoración de Artículo](#notificación-de-cambio-de-valoración-de-artículo)
	- [Notificación de Valoración Alta](#notificación-de-valoración-alta)
	- [Alerta de Valoración Baja](#alerta-de-valoración-baja)
	
- [Reglas_de_Valoraci_n](#reglas_de_valoraci_n)
	- [Asignar Parámetro a Artículo](#asignar-parámetro-a-artículo)
	
- [Sistema](#sistema)
	- [Check](#check)
	
- [Valoraci_n](#valoraci_n)
	- [Buscar Historial de Artículo](#buscar-historial-de-artículo)
	


# <a name='rabbitmq_get'></a> RabbitMQ_GET

## <a name='buscar-validación-de-artículo'></a> Buscar Validación de Artículo
[Back to top](#top)

<p>Listen validation product messages from catalog.</p>

	DIRECT feedback/article-validation





### Success Response

Message

```
		{
     	"type": "article-exist",
			"message" :
				{
					"articleId": "{articleId}",
					"valid": true|false
				}
     }
```


## <a name='buscar-reseña'></a> Buscar Reseña
[Back to top](#top)

<p>Escucha los mensajes de creación de Feedback para obtener las valoraciones</p>

	DIRECT feedback/new-feedback





### Success Response

Message

```
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


## <a name='logout-de-usuarios'></a> Logout de Usuarios
[Back to top](#top)

<p>Escucha de mensajes de logout desde auth.</p>

	FANOUT auth/logout





### Success Response

Message

```
{
   "type": "logout",
   "message": "{tokenId}"
}
```


# <a name='rabbitmq_post'></a> RabbitMQ_POST

## <a name='product-validation'></a> Product Validation
[Back to top](#top)

<p>Sending a validation request for a product.</p>

	DIRECT cart/article-exist





### Success Response

Message

```
    {
			"type": "article-exist",
			"queue": "catalog",
			"exchange": "",
			"message" : {
				"referenceId": "{referenceId}",
            	"articleId": "{articleId}",
			}
		}
```


## <a name='notificación-de-cambio-de-valoración-de-artículo'></a> Notificación de cambio de Valoración de Artículo
[Back to top](#top)

<p>Se notifica cada vez que cambia el promedio de la valoración de un artículo.</p>

	TOPIC rates/article-change-rate





### Success Response

Message

```
{
   "type": "article-change-rate",
   "queue": "rates"
   "message": {
        "articleId": "{article's id}",
        "newRate": "{article rate's value}",
        "feedAmount": "{amount of califications}"
    }
}
```


## <a name='notificación-de-valoración-alta'></a> Notificación de Valoración Alta
[Back to top](#top)

<p>Si una reseña supera la regla de una buena valoración, se notifica.</p>

	TOPIC rates/high-rate





### Success Response

Message

```
{
   "type": "high-rate",
   "queue": "rates"
   "message": {
        "articleId" : "{article's id}",
        "userId" : "{user's id}",
        "rate": "{article rate's value}",
    }
}
```


## <a name='alerta-de-valoración-baja'></a> Alerta de Valoración Baja
[Back to top](#top)

<p>Si una reseña supera la regla de una mala valoración, se genera una alerta.</p>

	TOPIC rates/low-rate





### Success Response

Message

```
{
   "type": "low-rate",
   "queue": "rates"
   "message": {
        "articleId" : "{article's id}",
        "userId" : "{user's id}",
        "rate": "{article rate's value}",
    }
}
```


# <a name='reglas_de_valoraci_n'></a> Reglas_de_Valoraci_n

## <a name='asignar-parámetro-a-artículo'></a> Asignar Parámetro a Artículo
[Back to top](#top)

<p>ABM Reglas</p>

	POST /v1/rates/:articleId/rules



### Examples

Body

```
   {
     "lowRate": "{bad rate's value}",
		"highRate": "{good rate's value}"
   }
```


### Success Response

Response

```
		HTTP/1.1 200 OK
		{
  		"articleId": "{article's id}",
 		"lowRate": "{bad rate's value}",
			"highRate": "{good rate's value}",
			"created": "{creation date}",
  		"modified": "{modification date}"
		}
```


# <a name='sistema'></a> Sistema

## <a name='check'></a> Check
[Back to top](#top)

<p>Verifica estado del Sistema</p>

	GET /v1/





### Success Response

Response

```
HTTP/1.1 200 OK
{
	"msg": "Running",
}
```


# <a name='valoraci_n'></a> Valoraci_n

## <a name='buscar-historial-de-artículo'></a> Buscar Historial de Artículo
[Back to top](#top)

<p>ABM Reglas</p>

	GET /v1/rates/:articlesd/history





### Success Response

Response

```
HTTP/1.1 200 OK
{
	"articleId": "{article's id}",
	"history": [
	    {
	        "rate": "{rate's value}",
	        "userId": "{user's id}",
	        "created": "{creation date}"
	    }
	]
}
```


