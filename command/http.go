package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/neovim/go-client/nvim"
)

func (r *Renata) RenataHttp(v *nvim.Nvim, args []string) error {
	if len(args) < 2 {
		return errors.New("usage: :Http <METHOD> <URL> ( <HEADER>:<VALUE>... )")
	}
	method := args[0]
	if method != "GET" && method != "POST" && method != "PUT" && method != "PATCH" && method != "DELETE" {
		return fmt.Errorf("unknown this command %s", method)
	}

	url := args[1]

	// バッファを読む
	buf, err := v.CurrentBuffer()
	if err != nil {
		return err
	}

	lines, err := v.BufferLines(buf, 0, -1, true)
	size := 0
	for _, l := range lines {
		size += len(l)
	}

	line := make([]byte, 0, size+len(lines))
	for _, l := range lines {
		l = []byte(strings.TrimSpace(string(l)))
		line = append(line, l...)
	}

	payload := string(line)

	switch method {
	case "GET":
		getRequest(url, payload)
	}
	return nil
}

func postRequest(url, payload string, header map[string]string) error {
	return nil
}
