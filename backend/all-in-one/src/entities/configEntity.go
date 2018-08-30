package entities


type ConfigEntity struct {
    // root:123456@tcp(192.168.121.186:3306)/tour_jp?charset=utf8
    MysqlUrl    string  
    MysqlMaxConnection  int
    // image url
    ImageUrlShop string
    ImageUrlSet string
    ImageUrlDish string
    // upload image path
    ImageFolder string
    // qrcode url
    UrlH5Employee string
    UrlH5Boss string
    UrlH5Customer string
    
}




