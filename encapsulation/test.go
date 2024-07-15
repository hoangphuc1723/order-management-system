package encapsulation

import "fmt"

// Encapsulation struct có thể exported ra bên ngoài pagekage này (Encapsulation viết hoa chữ cái đầu)
type Encapsulation struct{}

// Hàm Expose có thể exported ra bên ngoài pagekage này (Expose viết hoa chữ cái đầu)
func (e *Encapsulation) Expose() {
	fmt.Println("AHHHH! I'm exposed!")
}

// hàm hide chỉ có thể sử dụng trong package này (hide viết thường chữ cái đầu)
func (e *Encapsulation) hide() {
	fmt.Println("Shhhh... this is super secret")
}

// Unhide sử dụng hàm hide chưa được exported
func (e *Encapsulation) Unhide() {
	e.hide()
	fmt.Println("...jk")
}
