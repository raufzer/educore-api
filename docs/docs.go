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
        "/auth/login": {
            "post": {
                "description": "Authenticate user and generate access token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Login",
                "parameters": [
                    {
                        "description": "Login Credentials",
                        "name": "loginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.LoginResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Invalid username or password",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "User Registration",
                "parameters": [
                    {
                        "description": "User Registration Details",
                        "name": "userRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully created user",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "400": {
                        "description": "Invalid registration details",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Retrieves a list of all users in the system",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List all users",
                "responses": {
                    "200": {
                        "description": "Users retrieved successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/response.UserResponse"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "User Creation Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUsersRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        },
        "/users/{name}": {
            "get": {
                "description": "Fetches user details for a specific username",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Retrieve a user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User found",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates user details by username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User Update Request",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated successfully",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad request - invalid input",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a user by username",
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/response.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.CreateUsersRequest": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.LoginRequest": {
            "type": "object",
            "required": [
                "name",
                "password"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "EduCore API",
	Description:      "Complete API for EduCore Platform",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
