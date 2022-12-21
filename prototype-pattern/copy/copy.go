package copy

import (
	"encoding/json"
	"time"
)

// Keyword 搜索关键字
type Keyword struct {
	Word     string     `json:"word"`
	Visit    int        `json:"visit"`
	UpdateAt *time.Time `json:"update_at"`
}

// Clone 这里使用序列号和反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var nk Keyword
	b, _ := json.Marshal(k)
	_ = json.Unmarshal(b, &nk)
	return &nk
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (ks *Keywords) Clone(updatedWords []*Keyword) Keywords {
	nks := Keywords{}
	for k, v := range *ks {
		// 这里是浅拷贝，直接拷贝了地址
		nks[k] = v
	}

	// 替换掉需要更新的自旋，这里用的是深拷贝
	for _, w := range updatedWords {
		nks[w.Word] = w.Clone()
	}

	return nks
}
