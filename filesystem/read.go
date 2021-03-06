package filesystem

import (
	"github.com/vmihailenco/msgpack"
	"io/ioutil"
	"os"
)

func Load(path Filepath, dest interface{}) error {
	f, err := os.OpenFile(path.Unwrap(), os.O_RDONLY, 0664)
	if err != nil {
		return err
	}
	dec := msgpack.NewDecoder(f)
	return dec.Decode(dest)
}

func LoadRaw(path Filepath) ([]byte, error) {
	data, err := ioutil.ReadFile(path.Unwrap())
	return data, err
}
