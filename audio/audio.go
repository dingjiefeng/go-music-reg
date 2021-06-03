package audio

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
)

type Audio struct {
	fileName string
	Wave
}

func (audio *Audio) SetFile(fileLoc string) {
	audio.fileName = fileLoc
}

func (audio *Audio) readWavHead(b []byte) WaveHeader {
	header := WaveHeader{}
	header.ChunkId = string(b[0:4])
	if header.ChunkId != "RIFF" {
		panic("Invalid file")
	}
	sizeInfo := b[4:8]
	var size uint32
	buf := bytes.NewReader(sizeInfo)
	err := binary.Read(buf, binary.LittleEndian, &size)
	if err != nil {
		panic(err)
	}
	header.ChunkSize = int(size)

	format := b[8:12]
	header.Format = string(format)
	if header.Format != "WAVE" {
		panic("Format should be WAVE")
	}
	return header
}

func (audio *Audio) readFmt(b []byte) WaveFormat {
	waveFormat := WaveFormat{}
	id := b[12:16]
	waveFormat.ID = string(id)
	waveFormat.Size = bits32ToInt(b[16:20])
	waveFormat.AudioFormat = bits16ToInt(b[20:22])
	waveFormat.NumChannels = bits16ToInt(b[22:24])
	waveFormat.SampleRate = bits32ToInt(b[24:28])
	waveFormat.ByteRate = bits32ToInt(b[28:32])
	waveFormat.BlockAlign = bits16ToInt(b[32:34])
	waveFormat.BitsPerSample = bits16ToInt(b[34:36])
	if waveFormat.Size != 16 {
		panic("sub chunk size != 16")
	}

	return waveFormat
}

func (audio *Audio) readData(b []byte) WaveData {
	waveData := WaveData{}

	chunkID := string(b[36:40])
	var listSize = 0
	// skip LIST chunk
	if chunkID == "LIST" {
		listSize = bits32ToInt(b[40:44])
	}
	// seek data chunk
	waveData.ID = string(b[44+listSize : 48+listSize])
	waveData.Size = bits32ToInt(b[48+listSize : 52+listSize])
	waveData.RawData = b[52+listSize:]
	return waveData
}

// 为了简便 双声道 取平均转变为 单声道
func parseRawData(wfmt WaveFormat, rawData []byte) []Frame {
	bytesPerSample := wfmt.BitsPerSample / 8
	bitsNumPerSegment := wfmt.BitsPerSample * wfmt.NumChannels
	frames := []Frame{}
	for i := 0; i < len(rawData); i += bitsNumPerSegment {
		frame := Frame(0)
		for j := 0; j < wfmt.NumChannels; j++ {
			frame += Frame(byteSizeToIntFunc[8*bytesPerSample](rawData[i:i+bytesPerSample]) / wfmt.NumChannels)
		}
		frames = append(frames, frame)
	}

	return frames
}

func (audio *Audio) ResloveContent() {
	data, err := ioutil.ReadFile(audio.fileName)
	if err != nil {
		audio.Wave = Wave{}
		return
	}
	audio.Header = audio.readWavHead(data)
	audio.Format = audio.readFmt(data)
	audio.Content = audio.readData(data)
	audio.Content.Frames = parseRawData(audio.Format, audio.Content.RawData)
}
