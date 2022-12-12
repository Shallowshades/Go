package compressor

type RawCompressor struct{}

func (_ RawCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (_ RawCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}
