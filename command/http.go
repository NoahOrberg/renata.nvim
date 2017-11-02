package command

import (
	"strconv"

	"github.com/neovim/go-client/nvim"
)

func (r *Renata) RenataHttp(v *nvim.Nvim, args []string) error {
	v.Command("echom " + strconv.Itoa(len(args)))
	v.Command("echom " + args[0])
	// バッファを読む
	// それが妥当なJSONか確認する
	// 引数一番目にメソッド、二番目にURLを受け取る
	// 第三引数以降でヘッダなどを追加できるようにする
	return nil
}
