package message

import (
	"encoding/binary"
	"fmt"
	"io"
)

type messageID uint8

const (
	// MsgChoke chokes the receiver
	MsgChoke messageID = 0
	// MsgUnchoke unchokes the receiver
	MsgUnchoke messageID = 1
	// MsgInterested expresses interest in receiving data
	MsgInterested messageID = 2
	// MsgNotInterested expresses disinterest in receiving data
	MsgNotInterested messageID = 3
	// MsgHave alerts the receiver that the sender has downloaded a piece
	MsgHave messageID = 4
	// MsgBitfield encodes which pieces that the sender has downloaded
	MsgBitfield messageID = 5
	// MsgRequest requests a block of data from the receiver
	MsgRequest messageID = 6
	// MsgPiece delivers a block of data to fulfill a request
	MsgPiece messageID = 7
	// MsgCancel cancels a request
	MsgCancel messageID = 8
)

type Message struct {
	ID      messageID
	Payload []byte
}

// Read parses a message from a stream. Returns `nil` on keep-alive message
func Read(r io.Reader) (*Message, error) {
	lengthBuf := make([]byte, 4)
	_, err := io.ReadFull(r, lengthBuf)
	if err != nil {
		return nil, err
	}
	fmt.Println(lengthBuf)
	length := binary.BigEndian.Uint32(lengthBuf)

	// keep-alive message
	if length == 0 {
		return nil, nil
	}

	messageBuf := make([]byte, length)
	_, err = io.ReadFull(r, messageBuf)
	if err != nil {
		return nil, err
	}

	m := Message{
		ID:      messageID(messageBuf[0]),
		Payload: messageBuf[1:],
	}

	return &m, nil
}
