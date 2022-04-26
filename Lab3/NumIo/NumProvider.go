package NumIo

import (
	"errors"
	"fmt"
	"io"
	"os"
)

type NumFile struct {
	path string;
	values []float64;
	isRead bool;
	isDirty bool;
};

func (file *NumFile) read() []float64 {
	descr, err := os.Open(file.path);

	defer func() {
		if(descr != nil) {
			descr.Close();
		}
	}()

	if(err != nil) {
		errors.New(err.Error());
	}

	item := 0.0;
	content := make([]float64, 0);

	for {
		_, err := fmt.Fscanf(descr, "%f\n", &item);

		if(err != nil) {
			if err == io.EOF {
				break;
			}
			errors.New(err.Error());
		}

		content = append(content, item);
	}

	return content;
}

func (file *NumFile) write(content []float64) {
	descr, err := os.Open(file.path);

	defer func() {
		if(descr != nil) {
			descr.Close();
		}
	}()

	if(err != nil) {
		errors.New(err.Error());
	}

	for _, item := range content {
		_, err := fmt.Fprintf(descr, "%f\n", item);

		if(err != nil) {
			errors.New(err.Error());
		}
	}
}

func NewNumFile(path string) NumFile {
	return NumFile{ path: path };
}

func (file *NumFile) refresh() {
	if(!file.isRead) {
		file.values = file.read();
		file.isRead = true;
	}
}

func (file *NumFile) Add(n float64) {
	file.refresh();
	file.values = append(file.values, n);
	file.isDirty = true;
}

func (file *NumFile) RemoveAll(n float64) {
	file.refresh();

	for {
		if !file.RemoveFirst(n) {
			break;
		}
	}
}

func (file *NumFile) RemoveFirst(n float64) bool {
	file.refresh();
	for i, val := range file.values {
		if(val == n) {
			file.values = append(file.values[:i], file.values[i + 1:]...);
			file.isDirty = true;
			return true;
		}
	}
	return false;
}

func (file *NumFile) Clear() {
	file.refresh();
	file.values = make([]float64, 0);
	file.isDirty = true;
}

func (file *NumFile) Commit() {
	if(file.isDirty) {
		file.write(file.values);
		file.isDirty = false;
	}
}

func (file *NumFile) Get() []float64 {
	return file.values;
}