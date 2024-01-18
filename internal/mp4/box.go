package mp4

import (
	"encoding/binary"
	"errors"
	"fmt"
)

type Box struct {
	Type   string
	Raw    []byte
	Offset int64
	Size   int64
	Child  []Box // child boxes
}

func (b *Box) Find(t string) *Box {
	for i := range b.Child {
		child := &b.Child[i]
		if child.Type == t {
			return child
		}
	}
	return nil
}

func (b *Box) unpack() error {
	if b.Type != "moov" && b.Type != "trak" && b.Type != "mdia" {
		return nil
	}

	for off := 0; off < len(b.Raw); {
		if len(b.Raw[off:]) < 8 {
			return fmt.Errorf("%s unpack", b.Type)
		}

		c := Box{
			Offset: b.Offset + int64(off),
			Size:   int64(binary.BigEndian.Uint32(b.Raw[off:])),
			Type:   string(b.Raw[off+4 : off+8]),
		}

		off += 8
		datalen := c.Size - 8
		c.Raw = b.Raw[off : off+int(datalen)]
		b.Child = append(b.Child, c)
		off += int(datalen)
	}

	for i := range b.Child {
		c := &b.Child[i]
		if err := c.unpack(); err != nil {
			return err
		}
	}
	return nil
}

func (b *Box) WidthAndHeight() (width, height int, err error) {
	tkhd := b.Find("tkhd")
	if tkhd == nil {
		return 0, 0, fmt.Errorf("tkhd not found")
	}

	hd, err := NewTKHD(tkhd.Raw)
	if err != nil {
		return 0, 0, fmt.Errorf("decode TKHD error %w", err)
	}

	return int(hd.Width >> 16), int(hd.Height >> 16), nil
}

func (b *Box) AudioBitrate() (bitrate uint32, err error) {
	mdia := b.Find("mdia")
	if mdia == nil {
		return 0, errors.New("mdia not found")
	}

	m, err := NewMVHD(mdia.Raw)
	if err != nil {
		return 0, fmt.Errorf("decode MVHD error %w", err)
	}

	return m.Bitrate, nil
}
