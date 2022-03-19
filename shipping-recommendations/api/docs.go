// Package api GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package api

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
        "/v1/shipping-recommendations": {
            "get": {
                "description": "List recommended shipping options sorted by the best combination of cost and time.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Shipping Recommendations"
                ],
                "summary": "List recommended shipping options",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.ShippingRecommendationsResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "presenter.ShippingOptionResponse": {
            "type": "object",
            "properties": {
                "cost": {
                    "type": "number"
                },
                "estimated_days": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "presenter.ShippingRecommendationsResponse": {
            "type": "object",
            "properties": {
                "shipping_recommendations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/presenter.ShippingOptionResponse"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Shipping Recommendations API",
	Description:      "A RESTful API able to recommend the same shipping options ordered by the best combination of cost and time.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}