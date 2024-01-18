package mp4

import (
	"encoding/binary"
	"errors"
	"io"
)

type File struct {
	Box
	Moov *Box
}

func (f *File) GetTracks() []Box {
	if f.Moov == nil {
		return []Box{}
	}

	return f.Moov.Child
}

func Parse(r io.Reader) (file *File, err error) {
	file = &File{
		Box: Box{Type: "root", Size: -1},
	}

	p := parser{
		r: r,
		f: file,
	}

	err = p.Parse()
	if err != nil {
		return nil, err
	}

	moov := file.Box.Find("moov")
	if moov == nil {
		return nil, errors.New("moov missing - wrong mp4 file")
	}

	err = moov.unpack()
	if err != nil {
		return nil, err
	}

	file.Moov = moov

	return file, nil
}

type parser struct {
	r   io.Reader
	f   *File
	off int64
}

func (p *parser) Parse() (err error) {
	for {
		b := Box{}

		x := make([]byte, 8)
		var n int
		b.Offset = p.off
		n, err = io.ReadFull(p.r, x)
		p.off += int64(n)

		b.Size = int64(binary.BigEndian.Uint32(x))
		b.Type = string(x[4:])

		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		size := b.Size - 8
		if b.Type == "moov" {
			b.Raw = make([]byte, int(size))
			if _, err := io.ReadFull(p.r, b.Raw); err != nil {
				return err
			}
			p.off += int64(len(b.Raw))
		} else {
			if err := p.skip(size); err != nil {
				return err
			}
		}
		p.f.Child = append(p.f.Child, b)
	}
}

func (p *parser) skip(n int64) error {
	var err error
	if s, ok := p.r.(io.Seeker); ok {
		_, err = s.Seek(n, 1)
	}

	if err == nil {
		p.off += n
	}
	return err
}
