package entities

type DishJson struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"id"  db:"id"  ui:"" `
    Code   string   `json:"code"  db:"-"`  
	JpName string    `json:"jp_name"  db:"-"`
    CnName   string   `json:"cn_name"  db:"-"`  
	JpDescription string    `json:"jp_description"  db:"-"`
	CnDescription string    `json:"cn_description"  db:"-"`
	ImageUrl   string   `json:"image_url"  db:"-"` 
	JpMaterials   []string   `json:"jp_materials"  db:"-"` 
	CnMaterials   []string   `json:"cn_materials"  db:"-"`
}

type DishDB struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"-"  db:"id"  ui:"" `
    Code   string   `json:"-"  db:"code"`  
	JpName string    `json:"-"  db:"jp_name"`
    CnName   string   `json:"-"  db:"cn_name"`  
	JpDescription string    `json:"-"  db:"jp_description"`
	CnDescription string    `json:"-"  db:"cn_description"`
	ImageUrl   string   `json:"-"  db:"image_url"` 
	JpMaterials   string   `json:"-"  db:"jp_materials"` 
	CnMaterials   string   `json:"-"  db:"cn_materials"`
}
