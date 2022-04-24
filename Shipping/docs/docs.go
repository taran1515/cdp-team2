// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/shipping-address/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "add shipping address to a customer",
                "operationId": "Create",
                "parameters": [
                    {
                        "description": "address",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/shipping-address/customer/{customerId}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "get shipping address of a customer",
                "operationId": "GetByCustomer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customerId",
                        "name": "customerId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        },
        "BasicAuth": {
            "type": "basic"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8003",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Shipping Address Microservice",
	Description:      "This is shipping address service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}