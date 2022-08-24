package pkg

type fileObject struct {
	filepath string
	filename string
	filetype string
}

func (fo *fileObject) create(filename string) error {
	fo.filepath = "./statics"
	fo.filename = filename
	fo.filetype = ".csv"
	return nil
}
