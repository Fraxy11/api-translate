// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/translate-text-libre": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Translate"
                ],
                "parameters": [
                    {
                        "description": "PARAM",
                        "name": "parameter",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Translate"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "model.Translate": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string",
                    "example": "id"
                },
                "text": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "to": {
                    "type": "string",
                    "example": "en"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "API TRANSLATE",
	Description:      "Service For API TRANSLATE",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
