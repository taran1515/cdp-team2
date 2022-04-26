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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "API to check products-admin health",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Health Check Route",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "API to Login Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Login"
                ],
                "summary": "Login Route",
                "parameters": [
                    {
                        "description": "Username/email of the user",
                        "name": "username",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "Password of the user",
                        "name": "password",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/logout": {
            "post": {
                "description": "API to delete acess token from user databse and logout user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "Logout"
                ],
                "summary": "Logout Route",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Bearer \u003cAdd access token here\u003e",
                        "description": "Insert your access token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/oAuth/token": {
            "post": {
                "description": "API to Generate access token using basic token provided by login API",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "plain/text"
                ],
                "tags": [
                    "OAuth"
                ],
                "summary": "OAuth Access Token",
                "parameters": [
                    {
                        "type": "string",
                        "default": "Basic \u003cAdd access token here\u003e",
                        "description": "Insert your basic token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8088",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Swagger Frontstore Auth Microservice",
	Description:      "Micorservice for handling Frontstore Auth.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
