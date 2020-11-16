package common

import (
	"bytes"
	"encoding/json"
	"strconv"
)

// ProtocolVersion is the websocket protocol version
const ProtocolVersion = 1

// MessageType is the identifier code for websocket message types
type MessageType uint8

// 1 - 29 modify post model state
const (
	MessageInvalid MessageType = iota
	MessageInsertPost
	MessageAppend
	MessageBackspace
	MessageSplice
	MessageClosePost
	MessageInsertImage
	MessageSpoiler
	MessageModeratePost
)

// >= 30 are miscellaneous and do not write to post models
const (
	MessageSynchronise MessageType = 30 + iota
	MessageReclaim

	// Send new post ID to client
	MessagePostID

	// Concatenation of multiple websocket messages to reduce transport overhead
	MessageConcat

	// Message from the client meant to invoke no operation. Mostly used as a
	// one way ping, because the JS Websocket API does not provide access to
	// pinging.
	MessageNOOP

	// Transmit current synced IP count to client
	MessageSyncCount

	// Send current server Unix time to client
	MessageServerTime

	// Redirect the client to a specific board
	MessageRedirect

	// Send a notification to a client
	MessageNotification

	// Notify the client, he needs a captcha solved
	MessageCaptcha

	// Passes MeguTV playlist data
	MessageMeguTV

	// Used by the client to send it's protocol version and by the server to
	// send server and board configurations
	MessageConfigs
)

// Forwarded functions from "meguca/websockets/feeds" to avoid circular imports
var (
	// GetByIPAndBoard retrieves all Clients that match the passed IP on a board
	GetByIPAndBoard func(ip, board string) []Client

	// GetClientsByIP returns connected clients with matching ips
	GetClientsByIP func(ip string) []Client

	// SendTo sends a message to a feed, if it exists
	SendTo func(id uint64, msg []byte)

	// ClosePost closes a post in a feed, if it exists
	ClosePost func(id, op uint64, links []Link, commands []Command) error

	// Propagate a message about a post being moderated to connected clients
	PropagateModeration func(id, op uint64, entry ModerationEntry) error
)

// Client exposes some globally accessible websocket client functionality
// without causing circular imports
type Client interface {
	Send([]byte)
	Redirect(board string)
	IP() string
	LastTime() int64
	NewProtocol() bool
	Last100() bool
	Close(error)
}

// EncodeMessage encodes a message for sending through websockets or writing to
// the replication log.
func EncodeMessage(typ MessageType, msg interface{}) ([]byte, error) {
	data, err := json.Marshal(msg)
	switch {
	case err != nil:
		return nil, err
	case typ != MessageConcat && bytes.IndexRune(data, 0) != -1:
		return nil, ErrContainsNull
	default:
		return PrependMessageType(typ, data), nil
	}
}

// PrependMessageType prepends the encoded websocket message type to an already
// encoded message
func PrependMessageType(typ MessageType, data []byte) []byte {
	encoded := make([]byte, len(data)+2)

	// Ensure type string is always 2 chars long
	var i int
	if typ < 10 {
		encoded[0] = '0'
		i = 1
	}
	strconv.AppendUint(encoded[i:i], uint64(typ), 10)

	copy(encoded[2:], data)

	return encoded
}
