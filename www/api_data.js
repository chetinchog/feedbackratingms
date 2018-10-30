define({ "api": [
  {
    "type": "get",
    "url": "/v1/",
    "title": "Check",
    "name": "FeedbackRating",
    "group": "Server_Status",
    "description": "<p>Verify MicroService Status</p>",
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
    "filename": "./controllers/controller.go",
    "groupTitle": "Server_Status"
  }
] });
