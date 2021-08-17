package persistence

import "errors"

var ErrFind error = errors.New("見つかりませんでした。")
var ErrSave error = errors.New("登録に失敗しました。")
var ErrDelete error = errors.New("削除に失敗しました。")
