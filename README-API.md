<a name="top"></a>
# FeedbackRating Service vv0.14.7

FeedbackRatingMS Manager

- [Reglas_de_Valoraci_n](#reglas_de_valoraci_n)
	- [Asignar Parámetro a Artículo](#asignar-parámetro-a-artículo)
	
- [Sistema](#sistema)
	- [Check](#check)
	
- [Valoraci_n](#valoraci_n)
	- [Buscar Valoración de Artículo](#buscar-valoración-de-artículo)
	


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

## <a name='buscar-valoración-de-artículo'></a> Buscar Valoración de Artículo
[Back to top](#top)

<p>ABM Reglas</p>

	GET /v1/rates/:articleId/





### Success Response

Response

```
HTTP/1.1 200 OK
{
	"articleId": "{article's id}",
	"rate": "{article rate's value}",
	"ra1": "{amount of rates with value 1}",
	"ra2": "{amount of rates with value 2}",
	"ra3": "{amount of rates with value 3}",
	"ra4": "{amount of rates with value 4}",
	"ra5": "{amount of rates with value 5}",
	"feedAmount": "{amount of feedbacks made}",
	"badRate": "{is this category (boolean)}",
	"goodRate": "{is this category (boolean)}",
	"created": "{creation date}",
	"modified": "{modification date}"
}
```


