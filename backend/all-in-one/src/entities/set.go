package entities

// 从界面以json格式传入后台
type SetJsonIn struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"id"  db:"id"  ui:"" `
    Code   string   `json:"code"  db:"-"`  
	JpName string    `json:"jp_name"  db:"-"`
    CnName   string   `json:"cn_name"  db:"-"`  
	JpDescription string    `json:"jp_description"  db:"-"`
	CnDescription string    `json:"cn_description"  db:"-"`
	ImageUrl   string   `json:"image_url"  db:"-"` 
	Dishes   []string   `json:"dishes"  db:"-"` 
}

// 从后台以json格式返回前端
type SetJsonOut struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"id"  db:"id"  ui:"" `
    Code   string   `json:"code"  db:"-"`  
	JpName string    `json:"jp_name"  db:"-"`
    CnName   string   `json:"cn_name"  db:"-"`  
	JpDescription string    `json:"jp_description"  db:"-"`
	CnDescription string    `json:"cn_description"  db:"-"`
	ImageUrl   string   `json:"image_url"  db:"-"` 
	Dishes   []DishJson   `json:"dishes"  db:"-"` 
	CreatedOn   []string   `json:"created_on"  db:"-"`
}

// 写入数据的格式
type SetDB struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"-"  db:"id"  ui:"" `
    Code   string   `json:"-"  db:"code"`  
	JpName string    `json:"-"  db:"jp_name"`
    CnName   string   `json:"-"  db:"cn_name"`  
	JpDescription string    `json:"-"  db:"jp_description"`
	CnDescription string    `json:"-"  db:"cn_description"`
	ImageUrl   string   `json:"-"  db:"image_url"` 
	Dishes   string   `json:"-"  db:"dishes"` 
}
