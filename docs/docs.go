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
                                    "$ref": "#/definitions/main.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/main.Article"
                                            }
                                        }
                                    }
                                }
                            ]
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
                            "$ref": "#/definitions/main.UpdateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/main.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.JSONError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.JSONError"
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
                            "$ref": "#/definitions/main.CreateArticleModel"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/main.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/main.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/main.JSONError"
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
                                    "$ref": "#/definitions/main.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/main.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.JSONError"
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
                                    "$ref": "#/definitions/main.JSONResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/main.Article"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/main.JSONError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Article": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/main.Person"
                },
                "content": {
                    "$ref": "#/definitions/main.Content"
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
        "main.Content": {
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
        "main.CreateArticleModel": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/main.Person"
                },
                "content": {
                    "$ref": "#/definitions/main.Content"
                }
            }
        },
        "main.JSONError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "main.JSONResult": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "main.Person": {
            "type": "object",
            "required": [
                "firstname",
                "lastname"
            ],
            "properties": {
                "firstname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "John"
                },
                "lastname": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 2,
                    "example": "Doe"
                }
            }
        },
        "main.UpdateArticleModel": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/main.Person"
                },
                "content": {
                    "$ref": "#/definitions/main.Content"
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
