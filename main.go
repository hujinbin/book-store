package main

import (
	"net/http"
	"text/template"
	"book-store/controller"
)

func Indexhandler(w http.ResponseWriter, r *http.Request)  {
	t := template.Must(template.ParseFiles("page/index.html"))
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", Indexhandler)
	http.Handle("/files/",http.StripPrefix("/files/",http.FileServer(http.Dir("files"))))
	http.Handle("/page/",http.StripPrefix("/page/",http.FileServer(http.Dir("page"))))
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/checkUserName", controller.CheckUserName)

	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage) // 去更新图书的页面
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)   // 更新或添加图书
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)         // 添加图书到购物车

	http.HandleFunc("/getCartInfo", controller.GetCartInfo) // Get cart information
	http.HandleFunc("/deleteCart", controller.DeleteCart)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItem)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItem) // updateCartItem

	http.HandleFunc("/checkout", controller.Checkout)
	http.HandleFunc("/getOrders", controller.GetOrders)       // Get all orders
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfo) // Get all orders
	http.HandleFunc("/getOrderByUserID", controller.GetOrderByUserID)
	http.HandleFunc("/sendOrder", controller.SendOrder)
	http.HandleFunc("/takeOrder", controller.TakeOrder)

	http.ListenAndServe(":6060", nil)
}
