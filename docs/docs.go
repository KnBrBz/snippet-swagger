// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
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
        "/api/v1/endpoint": {
            "post": {
                "description": "Post entity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entity"
                ],
                "summary": "Post",
                "parameters": [
                    {
                        "description": "create entity",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Success"
                        }
                    }
                }
            }
        },
        "/api/v1/endpoint/{id}": {
            "get": {
                "description": "get all entities",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entities"
                ],
                "summary": "Get All",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Entity ID",
                        "name": "id",
                        "in": "path"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Posts"
                        }
                    }
                }
            },
            "put": {
                "description": "update entity",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entity"
                ],
                "summary": "Update",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Entity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "update entity",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/endpoint.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Success"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Success"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete entity",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "entity"
                ],
                "summary": "Delete",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Entity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/endpoint.Success"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "endpoint.Post": {
            "type": "object",
            "properties": {
                "alias": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "endpoint.Posts": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/endpoint.Post"
                    }
                }
            }
        },
        "endpoint.Success": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
