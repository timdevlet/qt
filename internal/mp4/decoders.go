package mp4

import "encoding/binary"

type MVHD struct {
	Duration uint32
	Bitrate  uint32
}

func NewMVHD(p []byte) (*MVHD, error) {
	m := &MVHD{}

	bp := &rawBox{data: p}

	bp.next(4 + 4 + 4 + 4)
	m.Duration = bp.Uint32()
	m.Bitrate = bp.Uint32()

	return m, nil
}

type TKHD struct {
	Duration uint32
	Width    uint32
	Height   uint32
}

func NewTKHD(p []byte) (*TKHD, error) {
	h := &TKHD{}

	bp := &rawBox{data: p}

	bp.next(4 + 4 + 4 + 4 + 8)
	h.Duration = bp.Uint32()
	bp.Skip(48)
	h.Width = bp.Uint32()
	h.Height = bp.Uint32()

	return h, nil
}

type rawBox struct {
	data    []byte
	i       int
	scratch [8]byte
}

func (p *rawBox) next(n int) []byte {
	i := p.i
	p.i += n
	if p.i <= len(p.data) {
		return p.data[i:p.i]
	}

	return p.scratch[:n]
}

func (p *rawBox) Skip(n int) {
	p.i += n
}

func (p *rawBox) Uint32() uint32 {
	return binary.BigEndian.Uint32(p.next(4))
}
