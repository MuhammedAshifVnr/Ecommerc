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
        "/admin/category": {
            "get": {
                "description": "Admin can see the listed  category in ecommerce website",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Category"
                ],
                "summary": "get all category list by admin",
                "responses": {}
            },
            "post": {
                "description": "Add a new category with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Category"
                ],
                "summary": "Add a new category",
                "operationId": "add-category",
                "parameters": [
                    {
                        "description": "Category Add",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.CategoryData"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully added.",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Duplicate found.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/category/{ID}": {
            "put": {
                "description": "Admin can edit the listed  category in ecommerce website",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Category"
                ],
                "summary": "Admin can edit the category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category Edit",
                        "name": "Form",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.CategoryData"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Deletes a category by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Category"
                ],
                "summary": "Delete a category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "Admin can Block the listed  category in ecommerce website",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Category"
                ],
                "summary": "Admin can block the category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/coupon": {
            "get": {
                "description": "Admin side list all the coupons",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Coupons"
                ],
                "responses": {
                    "200": {
                        "description": "Coupons",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Admin can add the coupon in ecommerce website",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Coupons"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coupon Code",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Coupon Amount",
                        "name": "amount",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Limit Amount",
                        "name": "limit",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coupon added",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/coupon/{ID}": {
            "delete": {
                "description": "Admin side delete a coupon",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Coupons"
                ],
                "summary": "Admin can delete a coupon",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Coupon ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Coupon deleted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
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
        "/admin/offer": {
            "get": {
                "description": "Product Offer Listing",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Offer"
                ],
                "summary": "Product Offer Listing",
                "responses": {}
            }
        },
        "/admin/offer/{ID}": {
            "post": {
                "description": "Product Offer Adding",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Offer"
                ],
                "summary": "Offer Adding",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Offer",
                        "name": "offer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.OfferProductData"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Product Offer Deleting",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Offer"
                ],
                "summary": "Product Offer Deleting",
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
        "/admin/order": {
            "get": {
                "description": "All Orders  are listed here",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin-Order"
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
                "tags": [
                    "Admin-Order"
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
        "/admin/salesreport": {
            "get": {
                "description": "Admin can download Sales Reoprt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Download Sales Reoprt",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter",
                        "name": "filter",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/admin/sort": {
            "get": {
                "description": "Admin can find best product and category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Finding best product and category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
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
        },
        "/user/address": {
            "get": {
                "description": "User can list Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Address"
                ],
                "summary": "Address listing",
                "responses": {}
            },
            "post": {
                "description": "User can Add Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Address"
                ],
                "summary": "Address Adding",
                "parameters": [
                    {
                        "description": "Address",
                        "name": "address",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.AddressData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/address/{ID}": {
            "delete": {
                "description": "User can Delete Address",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Address"
                ],
                "summary": "Address Delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "User can Edite Address",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Address"
                ],
                "summary": "Address Edit",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Address ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.AddressData"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/cart": {
            "get": {
                "description": "User can get the cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Cart"
                ],
                "summary": "Get cart",
                "responses": {}
            }
        },
        "/user/cart/{ID}": {
            "post": {
                "description": "User can add product in cart",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Cart"
                ],
                "summary": "Adding Product in cart",
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
            },
            "delete": {
                "description": "User can delete the product from cart",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Cart"
                ],
                "summary": "Delete in Cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "description": "User can update the quantity of product in cart",
                "consumes": [
                    "multipart/form-Data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Cart"
                ],
                "summary": "Update Quantity in Cart",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Quantity",
                        "name": "quantity",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/checkout": {
            "post": {
                "description": "Checkout Page",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Cart"
                ],
                "summary": "Checkout Page",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Payment Method",
                        "name": "payment",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Address",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "coupon",
                        "name": "coupon",
                        "in": "formData"
                    }
                ],
                "responses": {}
            }
        },
        "/user/home": {
            "get": {
                "description": "Get All Products",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Product"
                ],
                "summary": "Get All Products",
                "responses": {}
            }
        },
        "/user/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "user",
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
        "/user/logout": {
            "get": {
                "description": "Logout",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/order": {
            "get": {
                "description": "Order Listing",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Order"
                ],
                "summary": "Order Listing",
                "responses": {}
            }
        },
        "/user/order-item/{ID}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Order"
                ],
                "summary": "Order item listing",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/order/{ID}": {
            "patch": {
                "description": "Order Cancelation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Order"
                ],
                "summary": "Order Cancelation",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Order ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/productDetail/{ID}": {
            "get": {
                "description": "Product Detail",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Product"
                ],
                "summary": "Product Detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user/profile": {
            "get": {
                "description": "Get user profile",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profile"
                ],
                "summary": "User Profile",
                "responses": {}
            },
            "patch": {
                "description": "User can Edite Profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User-Profile"
                ],
                "summary": "Edite Profile",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/database.UserData"
                        }
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "database.AddressData": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "street": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "zip": {
                    "type": "integer"
                }
            }
        },
        "database.CategoryData": {
            "type": "object",
            "properties": {
                "catagory": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                }
            }
        },
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
        },
        "database.OfferProductData": {
            "type": "object",
            "properties": {
                "expirey": {
                    "type": "string"
                },
                "percentage": {
                    "type": "number"
                }
            }
        },
        "database.UserData": {
            "type": "object",
            "properties": {
                "gender": {
                    "type": "string"
                },
                "mobile": {
                    "type": "integer"
                },
                "username": {
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
