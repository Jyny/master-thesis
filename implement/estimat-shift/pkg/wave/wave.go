package wave

import (
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
)

func ReadWave(file string) (*audio.Float32Buffer, error) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	d := wav.NewDecoder(f)
	buf, err := d.FullPCMBuffer()
	if err != nil {
		return nil, err
	}

	return buf.AsFloat32Buffer(), nil
}
