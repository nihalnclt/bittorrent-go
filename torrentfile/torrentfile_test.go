package torrentfile

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	torrentFile, err := Open("testdata/archlinux-2019.12.01-x86_64.iso.torrent")
	require.Nil(t, err)

	fmt.Println(torrentFile)
}
