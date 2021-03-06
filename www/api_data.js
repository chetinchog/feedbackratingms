define({ "api": [
  {
    "type": "direct",
    "url": "feedback/article-validation",
    "title": "Buscar Validación de Artículo",
    "group": "RabbitMQ_GET",
    "description": "<p>Listen validation product messages from catalog.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "\t\t{\n     \t\"type\": \"article-exist\",\n\t\t\t\"message\" :\n\t\t\t\t{\n\t\t\t\t\t\"articleId\": \"{articleId}\",\n\t\t\t\t\t\"valid\": true|false\n\t\t\t\t}\n     }",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitGET.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "DirectFeedbackArticleValidation"
  },
  {
    "type": "direct",
    "url": "feedback/new-feedback",
    "title": "Buscar Reseña",
    "group": "RabbitMQ_GET",
    "description": "<p>Escucha los mensajes de creación de Feedback para obtener las valoraciones</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "\t\t{\n   \t\t\"type\": \"new-feedback\",\n   \t\t\"message\": {\n       \t\t\"id\" : \"{feedback's id}\"\n     \t\t \t\"idUser\" : \"{user's id}\",\n       \t\t\"idProduct\" : \"{product's id}\",\n       \t\t\"rate\" : \"{feedback's rate}\",\n       \t\t\"created\" : \"{creation date}\",\n       \t\t\"modified\" : \"{modification date}\"\n   \t\t}\n\t\t}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitGET.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "DirectFeedbackNewFeedback"
  },
  {
    "type": "fanout",
    "url": "auth/logout",
    "title": "Logout de Usuarios",
    "group": "RabbitMQ_GET",
    "description": "<p>Escucha de mensajes de logout desde auth.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "{\n   \"type\": \"logout\",\n   \"message\": \"{tokenId}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitGET.go",
    "groupTitle": "RabbitMQ_GET",
    "name": "FanoutAuthLogout"
  },
  {
    "type": "direct",
    "url": "cart/article-exist",
    "title": "Product Validation",
    "group": "RabbitMQ_POST",
    "description": "<p>Sending a validation request for a product.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "    {\n\t\t\t\"type\": \"article-exist\",\n\t\t\t\"queue\": \"catalog\",\n\t\t\t\"exchange\": \"\",\n\t\t\t\"message\" : {\n\t\t\t\t\"referenceId\": \"{referenceId}\",\n            \t\"articleId\": \"{articleId}\",\n\t\t\t}\n\t\t}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitPOST.go",
    "groupTitle": "RabbitMQ_POST",
    "name": "DirectCartArticleExist"
  },
  {
    "type": "topic",
    "url": "rates/article-change-rate",
    "title": "Notificación de cambio de Valoración de Artículo",
    "group": "RabbitMQ_POST",
    "description": "<p>Se notifica cada vez que cambia el promedio de la valoración de un artículo.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "{\n   \"type\": \"article-change-rate\",\n   \"queue\": \"rates\"\n   \"message\": {\n        \"articleId\": \"{article's id}\",\n        \"newRate\": \"{article rate's value}\",\n        \"feedAmount\": \"{amount of califications}\"\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitPOST.go",
    "groupTitle": "RabbitMQ_POST",
    "name": "TopicRatesArticleChangeRate"
  },
  {
    "type": "topic",
    "url": "rates/high-rate",
    "title": "Notificación de Valoración Alta",
    "group": "RabbitMQ_POST",
    "description": "<p>Si una reseña supera la regla de una buena valoración, se notifica.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "{\n   \"type\": \"high-rate\",\n   \"queue\": \"rates\"\n   \"message\": {\n        \"articleId\" : \"{article's id}\",\n        \"userId\" : \"{user's id}\",\n        \"rate\": \"{article rate's value}\",\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitPOST.go",
    "groupTitle": "RabbitMQ_POST",
    "name": "TopicRatesHighRate"
  },
  {
    "type": "topic",
    "url": "rates/low-rate",
    "title": "Alerta de Valoración Baja",
    "group": "RabbitMQ_POST",
    "description": "<p>Si una reseña supera la regla de una mala valoración, se genera una alerta.</p>",
    "success": {
      "examples": [
        {
          "title": "Message",
          "content": "{\n   \"type\": \"low-rate\",\n   \"queue\": \"rates\"\n   \"message\": {\n        \"articleId\" : \"{article's id}\",\n        \"userId\" : \"{user's id}\",\n        \"rate\": \"{article rate's value}\",\n    }\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./rabbit/rabbitPOST.go",
    "groupTitle": "RabbitMQ_POST",
    "name": "TopicRatesLowRate"
  },
  {
    "type": "post",
    "url": "/v1/rates/:articleId/rules",
    "title": "Asignar Parámetro a Artículo",
    "name": "FeedbackRating",
    "group": "Reglas_de_Valoraci_n",
    "description": "<p>ABM Reglas</p>",
    "examples": [
      {
        "title": "Body",
        "content": "   {\n     \"lowRate\": \"{bad rate's value}\",\n\t\t\"highRate\": \"{good rate's value}\"\n   }",
        "type": "json"
      }
    ],
    "success": {
      "examples": [
        {
          "title": "Response",
          "content": "\t\tHTTP/1.1 200 OK\n\t\t{\n  \t\t\"articleId\": \"{article's id}\",\n \t\t\"lowRate\": \"{bad rate's value}\",\n\t\t\t\"highRate\": \"{good rate's value}\",\n\t\t\t\"created\": \"{creation date}\",\n  \t\t\"modified\": \"{modification date}\"\n\t\t}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/rules.go",
    "groupTitle": "Reglas_de_Valoraci_n"
  },
  {
    "type": "get",
    "url": "/v1/rates/:articleId/rules",
    "title": "Buscar Parámetro a Artículo",
    "name": "FeedbackRating",
    "group": "Reglas_de_Valoraci_n",
    "description": "<p>Get Reglas</p>",
    "success": {
      "examples": [
        {
          "title": "Response",
          "content": "\t\tHTTP/1.1 200 OK\n\t\t{\n  \t\t\"articleId\": \"{article's id}\",\n \t\t\"lowRate\": \"{bad rate's value}\",\n\t\t\t\"highRate\": \"{good rate's value}\",\n\t\t\t\"created\": \"{creation date}\",\n  \t\t\"modified\": \"{modification date}\"\n\t\t}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/rules.go",
    "groupTitle": "Reglas_de_Valoraci_n"
  },
  {
    "type": "get",
    "url": "/v1/",
    "title": "Check",
    "name": "FeedbackRating",
    "group": "Sistema",
    "description": "<p>Verifica estado del Sistema</p>",
    "success": {
      "examples": [
        {
          "title": "Response",
          "content": "HTTP/1.1 200 OK\n{\n\t\"msg\": \"Running\",\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/system.go",
    "groupTitle": "Sistema"
  },
  {
    "type": "get",
    "url": "/v1/rates/:articlesd/history",
    "title": "Buscar Historial de Artículo",
    "name": "FeedbackRating",
    "group": "Valoraci_n",
    "description": "<p>ABM Reglas</p>",
    "success": {
      "examples": [
        {
          "title": "Response",
          "content": "HTTP/1.1 200 OK\n{\n\t\"articleId\": \"{article's id}\",\n\t\"history\": [\n\t    {\n\t        \"rate\": \"{rate's value}\",\n\t        \"userId\": \"{user's id}\",\n\t        \"created\": \"{creation date}\"\n\t    }\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/rates.go",
    "groupTitle": "Valoraci_n"
  },
  {
    "type": "get",
    "url": "/v1/rates/:articleId/",
    "title": "Buscar Valoración de Artículo",
    "name": "FeedbackRating",
    "group": "Valoraci_n",
    "description": "<p>ABM Reglas</p>",
    "success": {
      "examples": [
        {
          "title": "Response",
          "content": "HTTP/1.1 200 OK\n{\n\t\"articleId\": \"{article's id}\",\n\t\"rate\": \"{article rate's value}\",\n\t\"ra1\": \"{amount of rates with value 1}\",\n\t\"ra2\": \"{amount of rates with value 2}\",\n\t\"ra3\": \"{amount of rates with value 3}\",\n\t\"ra4\": \"{amount of rates with value 4}\",\n\t\"ra5\": \"{amount of rates with value 5}\",\n\t\"feedAmount\": \"{amount of feedbacks made}\",\n\t\"badRate\": \"{is this category (boolean)}\",\n\t\"goodRate\": \"{is this category (boolean)}\",\n\t\"created\": \"{creation date}\",\n\t\"modified\": \"{modification date}\"\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/rates.go",
    "groupTitle": "Valoraci_n"
  },
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./www/main.js",
    "group": "_home_che_go_src_github_com_chetinchog_feedbackratingms_www_main_js",
    "groupTitle": "_home_che_go_src_github_com_chetinchog_feedbackratingms_www_main_js",
    "name": ""
  }
] });
