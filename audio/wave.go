package audio

type WaveHeader struct {
	ChunkId   string
	ChunkSize int
	Format    string
}

type WaveFormat struct {
	ID            string
	Size          int
	AudioFormat   int
	NumChannels   int
	SampleRate    int
	ByteRate      int
	BlockAlign    int
	BitsPerSample int
}

type Frame int64

type WaveData struct {
	ID      string
	Size    int
	RawData []byte
	Frames  []Frame
}

type Wave struct {
	Header  WaveHeader
	Format  WaveFormat
	Content WaveData
}
