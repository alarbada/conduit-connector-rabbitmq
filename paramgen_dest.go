// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package rabbitmq

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	DestinationConfigDeliveryAppID           = "delivery.appID"
	DestinationConfigDeliveryContentEncoding = "delivery.contentEncoding"
	DestinationConfigDeliveryContentType     = "delivery.contentType"
	DestinationConfigDeliveryCorrelationID   = "delivery.correlationID"
	DestinationConfigDeliveryDeliveryMode    = "delivery.deliveryMode"
	DestinationConfigDeliveryExpiration      = "delivery.expiration"
	DestinationConfigDeliveryImmediate       = "delivery.immediate"
	DestinationConfigDeliveryMandatory       = "delivery.mandatory"
	DestinationConfigDeliveryMessageTypeName = "delivery.messageTypeName"
	DestinationConfigDeliveryPriority        = "delivery.priority"
	DestinationConfigDeliveryReplyTo         = "delivery.replyTo"
	DestinationConfigDeliveryUserID          = "delivery.userID"
	DestinationConfigExchangeAutoDelete      = "exchange.autoDelete"
	DestinationConfigExchangeDurable         = "exchange.durable"
	DestinationConfigExchangeInternal        = "exchange.internal"
	DestinationConfigExchangeName            = "exchange.name"
	DestinationConfigExchangeNoWait          = "exchange.noWait"
	DestinationConfigExchangeType            = "exchange.type"
	DestinationConfigQueueAutoDelete         = "queue.autoDelete"
	DestinationConfigQueueDurable            = "queue.durable"
	DestinationConfigQueueExclusive          = "queue.exclusive"
	DestinationConfigQueueName               = "queue.name"
	DestinationConfigQueueNoWait             = "queue.noWait"
	DestinationConfigRoutingKey              = "routingKey"
	DestinationConfigTlsCaCert               = "tls.caCert"
	DestinationConfigTlsClientCert           = "tls.clientCert"
	DestinationConfigTlsClientKey            = "tls.clientKey"
	DestinationConfigTlsEnabled              = "tls.enabled"
	DestinationConfigUrl                     = "url"
)

func (DestinationConfig) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		DestinationConfigDeliveryAppID: {
			Default:     "",
			Description: "AppID specifies the application that created the message.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryContentEncoding: {
			Default:     "",
			Description: "ContentEncoding specifies the encoding of the message content.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryContentType: {
			Default:     "application/json",
			Description: "ContentType specifies the MIME type of the message content. Defaults to \"application/json\".",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryCorrelationID: {
			Default:     "",
			Description: "CorrelationID is used to correlate RPC responses with requests.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryDeliveryMode: {
			Default:     "2",
			Description: "DeliveryMode indicates the message delivery mode. Non-persistent (1) or persistent (2). Default is 2 (persistent).",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryExpiration: {
			Default:     "",
			Description: "Expiration specifies the message expiration time, if any.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryImmediate: {
			Default:     "false",
			Description: "Immediate indicates if the message should be treated as immediate. If true, the message is not queued if no consumers are on the matching queue.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryMandatory: {
			Default:     "false",
			Description: "Mandatory indicates if the message is mandatory. If true, tells the server to return the message if it cannot be routed to a queue.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryMessageTypeName: {
			Default:     "",
			Description: "MessageTypeName specifies the message type name.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryPriority: {
			Default:     "0",
			Description: "Priority specifies the message priority. Ranges from 0 to 9. Default is 0.",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{
				config.ValidationGreaterThan{V: -1},
				config.ValidationLessThan{V: 10},
			},
		},
		DestinationConfigDeliveryReplyTo: {
			Default:     "",
			Description: "ReplyTo specifies the address to reply to.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigDeliveryUserID: {
			Default:     "",
			Description: "UserID specifies the user who created the message. Useful for publishers.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeAutoDelete: {
			Default:     "false",
			Description: "AutoDelete indicates if the exchange will be deleted when the last queue is unbound from it.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeDurable: {
			Default:     "true",
			Description: "Durable indicates if the exchange will survive broker restarts.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeInternal: {
			Default:     "false",
			Description: "Internal indicates if the exchange is used for internal purposes and cannot be directly published to by a client.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeName: {
			Default:     "",
			Description: "Name is the name of the exchange.",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeNoWait: {
			Default:     "false",
			Description: "NoWait indicates if the exchange should be declared without waiting for server confirmation.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigExchangeType: {
			Default:     "",
			Description: "Type is the type of the exchange (e.g., direct, fanout, topic, headers).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigQueueAutoDelete: {
			Default:     "false",
			Description: "AutoDelete indicates if the queue will be deleted when there are no more consumers.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigQueueDurable: {
			Default:     "true",
			Description: "Durable indicates if the queue will survive broker restarts.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigQueueExclusive: {
			Default:     "false",
			Description: "Exclusive indicates if the queue can be accessed by other connections.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigQueueName: {
			Default:     "",
			Description: "Name is the name of the queue to consume from / publish to",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		DestinationConfigQueueNoWait: {
			Default:     "false",
			Description: "NoWait indicates if the queue should be declared without waiting for server confirmation.",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigRoutingKey: {
			Default:     "",
			Description: "RoutingKey is the routing key to use when publishing to an exchange",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigTlsCaCert: {
			Default:     "",
			Description: "CACert is the path to the CA certificate to use for TLS",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigTlsClientCert: {
			Default:     "",
			Description: "ClientCert is the path to the client certificate to use for TLS",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigTlsClientKey: {
			Default:     "",
			Description: "ClientKey is the path to the client key to use for TLS",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		DestinationConfigTlsEnabled: {
			Default:     "false",
			Description: "Enabled indicates if TLS should be used",
			Type:        config.ParameterTypeBool,
			Validations: []config.Validation{},
		},
		DestinationConfigUrl: {
			Default:     "",
			Description: "URL is the RabbitMQ server URL",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
	}
}
