package pkg

type CsvParser interface {
	Create(filename string) error
}

type FileObject struct {
	filepath string
	filename string
	filetype string
}

func (fo *FileObject) Create(filename string) error {
	fo.filepath = "./statics"
	fo.filename = filename
	fo.filetype = ".csv"
	return nil
}
