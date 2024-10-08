package client

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/nihalnclt/bittorrent-go/bitfield"
	"github.com/nihalnclt/bittorrent-go/handshake"
	"github.com/nihalnclt/bittorrent-go/message"
	"github.com/nihalnclt/bittorrent-go/peers"
)

type Client struct {
	Conn     net.Conn
	Choked   bool
	Bitfield bitfield.Bitfield
	peer     peers.Peer
	infoHash [20]byte
	peerID   [20]byte
}

func completeHandshake(conn net.Conn, infoHash, peerID [20]byte) (*handshake.Handshake, error) {
	conn.SetDeadline(time.Now().Add(3 * time.Second))
	defer conn.SetDeadline(time.Time{}) // Disable the deadline

	req := handshake.New(infoHash, peerID)
	_, err := conn.Write(req.Serialize())
	if err != nil {
		return nil, err
	}

	res, err := handshake.Read(conn)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(res.InfoHash[:], infoHash[:]) {
		return nil, fmt.Errorf("expected infohash %x and got %x", res.InfoHash, infoHash)
	}

	return res, nil
}

func recvBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	conn.SetDeadline(time.Now().Add(2 * time.Second))
	defer conn.SetDeadline(time.Time{})

	msg, err := message.Read(conn)
	if err != nil {
		return nil, err
	}

	if msg.ID != message.MsgBitfield {
		err := fmt.Errorf("expected bitfield but got ID %d", msg.ID)
		return nil, err
	}

	return msg.Payload, nil
}

func New(peer peers.Peer, peerID, infoHash [20]byte) (*Client, error) {
	conn, err := net.DialTimeout("tcp", peer.String(), 3*time.Second)
	if err != nil {
		conn.Close()
		return nil, err
	}

	_, err = completeHandshake(conn, infoHash, peerID)
	if err != nil {
		conn.Close()
		return nil, err
	}

	bf, err := recvBitfield(conn)
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &Client{
		Conn:     conn,
		Choked:   true,
		Bitfield: bf,
		peer:     peer,
		infoHash: infoHash,
		peerID:   peerID,
	}, nil
}
