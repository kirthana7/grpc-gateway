package genswagger

import (
	"encoding/json"
	"testing"

	"github.com/go-openapi/spec"
)

func TestAtlasPatch(t *testing.T) {
	input := `
{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/bank/address/{resource_id}": {
      "get": {
        "operationId": "Read",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/exampleRead",
              "additionalProperties": null
            }
          }
        },
        "parameters": [
          {
            "name": "application_name",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "resource_type",
            "in": "path",
            "required": true,
            "type": "string"
          },
          {
            "name": "resource_id",
            "in": "path",
            "required": true,
            "type": "string",
            "format": "int64"
          },
          {
            "name": "id",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          }
        ],
        "tags": [
          "Service"
        ]
      }
    },
    "/api/v1/examples/{id.value}": {
      "get": {
        "tags": [
          "Service"
        ],
        "operationId": "Read",
        "parameters": [
          {
            "type": "string",
            "description": "The string value.",
            "name": "id.value",
            "in": "path",
            "required": true
          },
          {
            "name": "id",
            "in": "query",
            "required": false,
            "type": "string",
            "format": "int64"
          }
        ],
        "responses": {
          "200": {
            "description": "GET operation response",
            "schema": {
              "$ref": "#/definitions/exampleRead"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "exampleRead": {
      "type": "object",
      "additionalProperties": null
    }
  }
}
`
	res := atlasSwagger([]byte(input), false, false)
	var swSpec spec.Swagger
	if err := json.Unmarshal([]byte(res), &swSpec); err != nil {
		t.Fatalf("can't parse result back: %v", err)
	}

	for _, path := range swSpec.Paths.Paths {
		if path.Get != nil {
			var resourceIDPresent bool
			for _, param := range path.Get.Parameters {
				if param.Name == "application_name" || param.Name == "resource_type" {
					t.Error("atlasPatch should filter out required params that are not part of URL")
				}
				if param.Name == "id.value" {
					t.Error("atlasPatch should replace id.value in URL with just id")
				}
				if param.Name == "id" && param.In == "query" {
					t.Error("atlasPatch should filter out id present in query")
				}
				if param.Name == "resource_id" || (param.Name == "id" && param.In == "path") {
					resourceIDPresent = true
				}
			}

			if !resourceIDPresent {
				t.Error("atlasPatch should not filter parameter with name resource_id")
			}
		}
	}
}
