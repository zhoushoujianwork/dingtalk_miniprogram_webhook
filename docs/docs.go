// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "DevOps Team"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v2023-03/webhook/callback": {
            "post": {
                "description": "webhook callback",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "webhook"
                ],
                "summary": "webhook callback",
                "parameters": [
                    {
                        "type": "string",
                        "description": "signature",
                        "name": "signature",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "timestamp",
                        "name": "timestamp",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "nonce",
                        "name": "nonce",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "input",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.WebHookRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.WebHookRequest": {
            "type": "object",
            "properties": {
                "encrypt": {
                    "description": "encrypt",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "2023-03",
	Host:             "api.pop.patsnap.info",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "mini program webhook API",
	Description:      "Patsnap OPS API spec.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
