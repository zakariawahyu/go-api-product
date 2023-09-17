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
        "/product": {
            "post": {
                "description": "Create new single product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create new product",
                "parameters": [
                    {
                        "description": "Create Product",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Get single product by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get product by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update single product by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update single product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Create Product",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateProduct"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Product"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete product by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Delete product by id",
                "parameters": [
                    {
                        "type": "string",
                        "description": "product id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateProduct": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "rating",
                "stock",
                "variety"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 3
                },
                "name": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 0
                },
                "stock": {
                    "type": "integer"
                },
                "variety": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateProduct": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "rating",
                "stock",
                "variety"
            ],
            "properties": {
                "description": {
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 3
                },
                "name": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 0
                },
                "stock": {
                    "type": "integer"
                },
                "variety": {
                    "type": "string"
                }
            }
        },
        "entity.Product": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "rating",
                "stock",
                "variety"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "maxLength": 500,
                    "minLength": 3
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "maxLength": 250,
                    "minLength": 3
                },
                "price": {
                    "type": "number"
                },
                "rating": {
                    "type": "integer",
                    "maximum": 10,
                    "minimum": 0
                },
                "slug": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "variety": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "errors": {},
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "response.SuccessResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "success": {
                    "type": "boolean",
                    "example": true
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}