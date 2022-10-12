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
        "/v1/article": {
            "get": {
                "description": "get articles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "List articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/models.Article"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    }
                }
            },
            "put": {
                "description": "update article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Update article",
                "parameters": [
                    {
                        "description": "Article body",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    }
                }
            },
            "post": {
                "description": "create new article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Create article",
                "parameters": [
                    {
                        "description": "Article body",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    }
                }
            }
        },
        "/v1/article/{id}": {
            "get": {
                "description": "get article by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Get article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.PackedArticleModel"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete article by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "articles"
                ],
                "summary": "Delete article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Article ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/models.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/models.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.JSONError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Article": {
            "type": "object",
            "required": [
                "author_id"
            ],
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "content": {
                    "$ref": "#/definitions/models.Content"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Author": {
            "type": "object",
            "required": [
                "firstname",
                "lastname"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "John"
                },
                "id": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "Doe"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.Content": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "models.CreateArticleModel": {
            "type": "object",
            "required": [
                "author_id"
            ],
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "content": {
                    "$ref": "#/definitions/models.Content"
                }
            }
        },
        "models.JSONError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.JSONResult": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "models.PackedArticleModel": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/models.Author"
                },
                "content": {
                    "$ref": "#/definitions/models.Content"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "models.UpdateArticleModel": {
            "type": "object",
            "required": [
                "author_id",
                "id"
            ],
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "content": {
                    "$ref": "#/definitions/models.Content"
                },
                "id": {
                    "type": "string"
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
