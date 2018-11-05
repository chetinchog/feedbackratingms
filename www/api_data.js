define({ "api": [
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
          "content": "HTTP/1.1 200 OK\n{\n\t\"id\": \"{article's id}\",\n\t\"history\": [\n\t    {\n\t        \"rate\": \"{rate's value}\",\n\t        \"userId\": \"{user's id}\",\n\t        \"created\": \"{creation date}\"\n\t    }\n\t]\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controllers/rates.go",
    "groupTitle": "Valoraci_n"
  }
] });