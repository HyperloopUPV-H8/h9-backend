package connection_transfer

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/HyperloopUPV-H8/Backend-H8/common"
	"github.com/HyperloopUPV-H8/Backend-H8/connection_transfer/models"
	"github.com/rs/zerolog"
	trace "github.com/rs/zerolog/log"
)

const (
	CONNECTION_TRANSFER_HANDLER_NAME = "connectionTransfer"
)

type ConnectionTransfer struct {
	writeMx     *sync.Mutex
	boardStatus map[string]models.Connection
	SendMessage func(topic string, payload any, target ...string) error
	updateTopic string
	trace       zerolog.Logger
}

type ConnectionTransferConfig struct {
	UpdateTopic string `toml:"update_topic"`
}

func New(config ConnectionTransferConfig) ConnectionTransfer {
	trace.Info().Msg("new connection transfer")

	return ConnectionTransfer{
		writeMx:     &sync.Mutex{},
		boardStatus: make(map[string]models.Connection),
		SendMessage: defaultSendMessage,
		updateTopic: config.UpdateTopic,
		trace:       trace.With().Str("component", CONNECTION_TRANSFER_HANDLER_NAME).Logger(),
	}
}

func (connectionTransfer *ConnectionTransfer) UpdateMessage(topic string, payload json.RawMessage, source string) {
	connectionTransfer.trace.Trace().Str("source", source).Str("topic", topic).Msg("got message")
	connectionTransfer.send()
}

func (connectionTransfer *ConnectionTransfer) SetSendMessage(sendMessage func(topic string, payload any, target ...string) error) {
	connectionTransfer.trace.Debug().Msg("set send message")
	connectionTransfer.SendMessage = sendMessage
}

func (connectionTransfer *ConnectionTransfer) HandlerName() string {
	return CONNECTION_TRANSFER_HANDLER_NAME
}

func (connectionTransfer *ConnectionTransfer) Update(name string, IsConnected bool) {
	connectionTransfer.writeMx.Lock()
	defer connectionTransfer.writeMx.Unlock()

	connectionTransfer.trace.Debug().Str("connection", name).Bool("isConnected", IsConnected).Msg("update connection state")

	connectionTransfer.boardStatus[name] = models.Connection{
		Name:        name,
		IsConnected: IsConnected,
	}

	connectionTransfer.send()
}

func (connectionTransfer *ConnectionTransfer) send() {
	connectionTransfer.trace.Debug().Msg("send connections")

	connArr := common.Values(connectionTransfer.boardStatus)

	if err := connectionTransfer.SendMessage(connectionTransfer.updateTopic, connArr); err != nil {
		connectionTransfer.trace.Error().Stack().Err(err).Msg("")
		return
	}
}

func defaultSendMessage(string, any, ...string) error {
	return errors.New("connection transfer must be registered before use")
}
