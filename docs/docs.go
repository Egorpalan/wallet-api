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
        "/api/v1/wallet": {
            "post": {
                "description": "Deposit or withdraw funds from a wallet",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Perform wallet operation",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.TransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Operation successful",
                        "schema": {
                            "$ref": "#/definitions/handler.StandardResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to process transaction",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets": {
            "post": {
                "description": "Create a new wallet with initial balance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Create a wallet",
                "parameters": [
                    {
                        "description": "Wallet ID",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Wallet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Wallet created successfully",
                        "schema": {
                            "$ref": "#/definitions/handler.StandardResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request payload",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to create wallet",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/wallets/{walletId}": {
            "get": {
                "description": "Get the current balance of a wallet by walletId",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get wallet balance",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Wallet ID",
                        "name": "walletId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Wallet balance",
                        "schema": {
                            "$ref": "#/definitions/handler.BalanceResponse"
                        }
                    },
                    "500": {
                        "description": "Failed to get balance",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entity.TransactionRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "operationType": {
                    "type": "string"
                },
                "walletId": {
                    "type": "string"
                }
            }
        },
        "entity.Wallet": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "walletId": {
                    "type": "string"
                }
            }
        },
        "handler.BalanceResponse": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "number"
                }
            }
        },
        "handler.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "handler.StandardResponse": {
            "type": "object",
            "properties": {
                "message": {
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Wallet API",
	Description:      "This is a wallet management API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
