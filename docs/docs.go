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
        "contact": {
            "name": "Airren, Peto",
            "email": "renqiqiang@outlook.com, peto1"
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
        "/file/download/{id}": {
            "post": {
                "description": "Download by file ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Download a file",
                "parameters": [
                    {
                        "type": "string",
                        "format": "octet-stream",
                        "description": "FILE ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    }
                }
            }
        },
        "/file/list": {
            "post": {
                "description": "List files by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "List files",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.File"
                            }
                        }
                    }
                }
            }
        },
        "/file/listByIds": {
            "get": {
                "description": "List files by file ids",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "List files",
                "parameters": [
                    {
                        "description": "FILE ID LIST",
                        "name": "ids",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.IdsReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.File"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    }
                }
            }
        },
        "/file/update/": {
            "put": {
                "description": "update file info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "update file info",
                "parameters": [
                    {
                        "type": "file",
                        "description": "FILE",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.FileVO"
                        }
                    }
                }
            }
        },
        "/file/upload/": {
            "post": {
                "description": "Upload a file",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "file"
                ],
                "summary": "Upload a file",
                "parameters": [
                    {
                        "type": "file",
                        "description": "FILE",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vo.FileVO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/vo.BaseVO"
                        }
                    }
                }
            }
        },
        "/task/create_by_file": {
            "post": {
                "description": "create task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Algorithm"
                ],
                "summary": "create task",
                "parameters": [
                    {
                        "description": "task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Algorithm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Algorithm"
                        }
                    }
                }
            }
        },
        "/task/list": {
            "post": {
                "description": "query task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Algorithm"
                ],
                "summary": "query task",
                "parameters": [
                    {
                        "description": "task req",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AlgorithmReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Algorithm"
                            }
                        }
                    }
                }
            }
        },
        "/task/update": {
            "put": {
                "description": "update task",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Job"
                ],
                "summary": "update task",
                "parameters": [
                    {
                        "description": "task",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Job"
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
        "/task/{id}": {
            "get": {
                "description": "Get details of task by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Algorithm"
                ],
                "summary": "get task by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "task id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Algorithm"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AlgoParameter": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "algorithmId": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "label": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "required": {
                    "type": "boolean"
                },
                "type": {
                    "$ref": "#/definitions/model.ParamType"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                },
                "value_list": {
                    "type": "string"
                }
            }
        },
        "model.Algorithm": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "command": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "favourite": {
                    "type": "integer"
                },
                "group_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AlgoParameter"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "model.File": {
            "type": "object",
            "properties": {
                "MD5": {
                    "type": "string"
                },
                "URLPath": {
                    "type": "string"
                },
                "account_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "file_type": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isPublic": {
                    "description": "allowed access by other user",
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "model.Job": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "string"
                },
                "algorithm": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "inputFile": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "outPutFile": {
                    "type": "string"
                },
                "parameter": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        },
        "model.ParamType": {
            "type": "string",
            "enum": [
                "string",
                "file",
                "radio",
                "select"
            ],
            "x-enum-varnames": [
                "ParamString",
                "ParamFile",
                "ParamRadio",
                "ParamSelect"
            ]
        },
        "req.AlgorithmReq": {
            "type": "object",
            "properties": {
                "asc": {
                    "type": "boolean"
                },
                "command": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "document": {
                    "type": "string"
                },
                "favourite": {
                    "type": "integer"
                },
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "image": {
                    "type": "string"
                },
                "label": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "order_by": {
                    "type": "string"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "page_size": {
                    "type": "integer",
                    "example": 10
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.AlgoParameter"
                    }
                },
                "price": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "req.IdsReq": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "vo.BaseVO": {
            "type": "object",
            "properties": {
                "asc": {
                    "type": "boolean"
                },
                "data": {},
                "error_code": {
                    "type": "integer"
                },
                "error_message": {
                    "type": "string"
                },
                "order_by": {
                    "type": "string"
                },
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "page_size": {
                    "type": "integer",
                    "example": 10
                },
                "success": {
                    "type": "boolean"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "vo.FileVO": {
            "type": "object",
            "properties": {
                "URLPath": {
                    "type": "string"
                },
                "account_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "deleted_by": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "file_type": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "isPublic": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "org": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "http://echo-bio.cn",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "Echo-Bio-Cloud",
	Description:      "Order Manager",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
