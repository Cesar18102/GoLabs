package TextIo

import (
	"fmt"
	"io/ioutil"
	"os"
)

type TextFile struct {
	path string;
};

func NewTextFile(path string) TextFile {
	return TextFile { path: path };
}

func (file *TextFile) Read() string {
	data, err := ioutil.ReadFile(file.path);

	if(err != nil) {
		fmt.Print("Failed opening file\n");
		return "";
	}

	return string(data);
}

func (file *TextFile) Write(text string) {
	os.Remove(file.path);
	descr, err := os.Create(file.path);

	defer func() {
		if(descr != nil) {
			descr.Close();
		}
	}()

	if(err != nil) {
		fmt.Printf("Failed writing file\n");
		return;
	}

	descr.WriteString(text);
}