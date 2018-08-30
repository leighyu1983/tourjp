package entities

type Shop struct {
    //Id string   `json:"-"  db:"id"  ui:"" `   // json:"-" 不显示
    Id string   `json:"id"  db:"id"  ui:"" `
    JpName   string   `json:"jp_name"  db:"jp_name"`  
	JpDescription string    `json:"jp_description"  db:"jp_description"`
    CnName   string   `json:"cn_name"  db:"cn_name"`  
    CnDescription string    `json:"cn_description"  db:"cn_description"`
    ImageUrl string `json:"image_url"  db:"image_url"`
}

