package structs

type Admin_info struct {
	Admin_id       int    `json:"admin_id"`
	Admin_name     string `json:"admin_name"`
	Admin_email    string `json:"admin_email"`
	Admin_password string `json:"admin_password"`
}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Brands struct {
	Brand_id    int    `json:"brand_id"`
	Brand_title string `json:"brand_title"`
}

type Cart struct {
	Id      int `json:"id"`
	P_id    int `json:"p_id"` //product id
	User_id int `json:"user_id"`
	Qty     int `json:"qty"` // number of product in cart
}

type Categories struct {
	Cat_id    int    `json:"cat_id"`
	Cat_title string `json:"cat_title"`
}

type Email_info struct {
	Email_id int    `json:"email_id"`
	Email    string `json:"email"`
}

type Logs struct {
	Id      int    `json:"id"`
	User_id int    `json:"user_id"`
	Action  string `json:"action"`
	Date    string `json:"date"`
}

type Orders struct {
	Order_id   int    `json:"order_id"`
	User_id    int    `json:"user_id"`
	Product_id int    `json:"product_id"`
	Qty        int    `json:"qty"`
	Trx_id     string `json:"trx_id"` //transaction id
	P_status   string `json:"p_status"`
}
type Order_products struct {
	Order_pro_id int    `json:"order_pro_id"`
	Order_id     int    `json:"order_id"`
	Product_id   string `json:"product_id"`
	Qty          int    `json:"qty"`
	Amt          int    `json:"amt"`
}

type Orders_info struct {
	Order_id   int    `json:"order_id"`
	User_id    int    `json:"user_id"`
	F_name     string `json:"f_name"`
	Email      string `json:"email"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        int    `json:"zip"`
	Cardname   string `json:"cardname"`
	Cardnumber string `json:"user_id"`
	Expdate    string `json:"expdate"`
	Prod_count int    `json:"prod_count"`
	Total_amt  int    `json:"total_amt"`
	Cvv        int    `json:"cvv"`
}

type User_info struct {
	User_id    int    `json:"user_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
}

type User_info_backup struct {
	User_id    int    `json:"user_id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Mobile     string `json:"mobile"`
	Address1   string `json:"address1"`
	Address2   string `json:"address2"`
}
type Product struct {
	Product_id       int    `json:"product_id"`
	Product_cat      int    `json:"product_cat"`
	Product_brand    int    `json:"product_brand"`
	Product_title    string `json:"product_title"`
	Product_price    int    `json:"product_price"`
	Product_desc     string `json:"product_desc"`
	Product_image    string `json:"product_image"`
	Product_keywords string `json:"product_keywords"`
	Product_count    int    `jason:"product_count"`
}
