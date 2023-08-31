// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/students": {
            "get": {
                "description": "Get a list of all students",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Get all students",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AllStudentResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseStudentNotFound"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new student",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Create a new student",
                "parameters": [
                    {
                        "description": "Student object that needs to be added",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StudentInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccessCreate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseInvalid"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseFailedCreate"
                        }
                    }
                }
            }
        },
        "/students/{id}": {
            "get": {
                "description": "Get student details by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Get student by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.StudentResponseByID"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseStudentNotFound"
                        }
                    }
                }
            },
            "put": {
                "description": "Update student details by ID with JSON data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Update student by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated student object",
                        "name": "student",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.StudentUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccessUpdate"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseInvalid"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseStudentNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseFailedUpdate"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete student by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "students"
                ],
                "summary": "Delete student by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Student ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseSuccessDelete"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseStudentNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ResponseFailedDelete"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AllStudentResponse": {
            "type": "object",
            "properties": {
                "students": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.StudentResponse"
                    }
                }
            }
        },
        "models.ResponseFailedCreate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Failed create student!"
                }
            }
        },
        "models.ResponseFailedDelete": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Failed delete student!"
                }
            }
        },
        "models.ResponseFailedUpdate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Failed update student!"
                }
            }
        },
        "models.ResponseInvalid": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Invalid request!"
                }
            }
        },
        "models.ResponseStudentNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Student not found!"
                }
            }
        },
        "models.ResponseSuccessCreate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully create student!"
                }
            }
        },
        "models.ResponseSuccessDelete": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully delete student!"
                }
            }
        },
        "models.ResponseSuccessUpdate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "Successfully update student!"
                }
            }
        },
        "models.StudentInput": {
            "type": "object",
            "properties": {
                "jurusan": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "nim": {
                    "type": "string"
                }
            }
        },
        "models.StudentResponse": {
            "type": "object",
            "properties": {
                "Created_at": {
                    "type": "string"
                },
                "Deleted_at": {
                    "type": "string"
                },
                "ID": {
                    "type": "integer"
                },
                "Jurusan": {
                    "type": "string"
                },
                "Name": {
                    "type": "string"
                },
                "Nim": {
                    "type": "string"
                },
                "Updated_at": {
                    "type": "string"
                }
            }
        },
        "models.StudentResponseByID": {
            "type": "object",
            "properties": {
                "student": {
                    "$ref": "#/definitions/models.StudentResponse"
                }
            }
        },
        "models.StudentUpdate": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRUD Students API",
	Description:      "API for CRUD Students",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}