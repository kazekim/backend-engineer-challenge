package grcgitscanner

type ScannerInput struct {
	FilePath string `json:"file_path"`
	Data string `json:"data"`
	Line int64 `json:"line"`
}

//NewScannerInput new ScannerInput data from parameters
func NewScannerInput(data, filePath string, line int64) ScannerInput {
	m := ScannerInput{
		FilePath: filePath,
		Data: data,
		Line: line,
	}
	return m
}