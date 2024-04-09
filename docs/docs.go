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
        "/admin/home": {
            "get": {
                "description": "after login show this page",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin Home Page",
                "responses": {}
            }
        },
        "/admin/login": {
            "post": {
                "description": "Login an admin user with username and password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login an admin user",
                "operationId": "admin-login",
                "parameters": [
                    {
                        "description": "Admin login details",
                        "name": "admin",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.Logging"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/admin/logout": {
            "get": {
                "description": "Logout an admin clearing the cookie",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Logout an admin user",
                "responses": {}
            }
        },
        "/admin/order": {
            "get": {
                "description": "All Orders  are listed here",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Orders Listing",
                "responses": {}
            }
        },
        "/admin/order/update/{ID}": {
            "patch": {
                "description": "Update the status of an order by ID.",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update Order Status",
                "operationId": "update-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "New status",
                        "name": "status",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/product": {
            "get": {
                "description": "get all product list by admin",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Product"
                ],
                "summary": "Admin can see the listed  products in ecommerce website",
                "responses": {}
            },
            "post": {
                "description": "Add a new product to the database",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Product"
                ],
                "summary": "Add a new product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the product",
                        "name": "product",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the category the product belongs to",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quantity of the product",
                        "name": "quantity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Price of the product",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Size of the product",
                        "name": "size",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Description of the product",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Images of the product (upload at least 3 images)",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Product added successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
        },
        "/admin/product/{ID}": {
            "put": {
                "description": "Edits a product by its ID, including updating its category, quantity, price, size, description, and images",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Product"
                ],
                "summary": "Edit a product",
                "operationId": "edit-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of the product",
                        "name": "product",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product category name",
                        "name": "category",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product quantity",
                        "name": "quantity",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product price",
                        "name": "price",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product size",
                        "name": "size",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Product description",
                        "name": "description",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "array",
                        "items": {
                            "type": "file"
                        },
                        "collectionFormat": "csv",
                        "description": "Product images",
                        "name": "images",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Deletes a product by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Product"
                ],
                "summary": "Delete a product",
                "operationId": "delete-product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/users": {
            "get": {
                "description": "Admin can see the users in ecommerce website",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Users"
                ],
                "responses": {}
            }
        },
        "/admin/users/{ID}": {
            "patch": {
                "description": "Admin can Block and Unblock the users in ecommerce website",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Users"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "database.Logging": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "useremail": {
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
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "ECOM",
	Description:      "This is a sample Gin API with Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
