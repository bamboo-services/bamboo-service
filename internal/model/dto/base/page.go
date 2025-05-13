package base

// Page 分页数据封装类型。
//
// 该类型用于封装分页数据，包括总页数、每页数量、当前页码以及泛型类型的数据内容。
//
// 泛型 E 表示数据内容类型。
type Page[E interface{}] struct {
	Total   int  `json:"total" dc:"总页数"`
	Limit   int  `json:"limit" dc:"每页数量"`
	Current int  `json:"current" dc:"当前页码"`
	Record  []*E `json:"record" dc:"数据"`
}

// NewPage 创建一个新的分页数据结构体实例。
//
// 参数:
//   - total: 数据总条数
//   - base: 当前页码
//   - limit: 每页条目数
//   - data: 泛型数据列表
//
// 返回:
//   - 指向包含分页数据的 Page 结构体的指针
func NewPage[E interface{}](total, page, limit int, data []*E) *Page[E] {
	return &Page[E]{
		Total:   total,
		Limit:   limit,
		Current: page,
		Record:  data,
	}
}
