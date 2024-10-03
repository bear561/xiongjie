package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Book struct represents a book with its details.
type Book struct {      
	ID       int    `json:"id"`      //图书id
	Title    string `json:"title"`     //图书标题
	Author   string `json:"author"`     //图书作者
	Borrowed bool   `json:"borrowed"`   //是否借书
	Borrower string `json:"borrower"`    //借书者
}
// 六花结构体表示一个使用文件存储数据的管理器
type Sixflower struct {      
	Books []Book
	FileName string
}

// 创造六花实例
func NewSixflower(fileName string) *Sixflower {
	return &Sixflower{
		Books:   []Book{},  //定义切片存储book
		FileName: fileName, //存储书籍数据的文件名
	}
	}


// 添加新书
func (a *Sixflower) AddBook() {
	var book Book
	fmt.Print("图书id是 ")
	fmt.Scan(&book.ID)
	fmt.Print("图书标题是: ")
	fmt.Scan(&book.Title)
	fmt.Print("图书作者是: ")
	fmt.Scan(&book.Author)
	book.Borrowed = false //默认没被借走
	a.Books = append(a.Books, book)
	a.SaveBooks()            //保存
}

// DeleteBook removes a book from the manager by its ID.
func (a *Sixflower) DeleteBook() {
	var id int
	fmt.Print("请输入图书id: ")
	fmt.Scan(&id)
	for i, book := range a.Books {
		if book.ID == id {
			a.Books = append(a.Books[:i], a.Books[i+1:]...)//处理索引为i的书其他全保留
			a.SaveBooks()            //保存
			fmt.Println("已经删除该本书")
			return
		}
	}
	fmt.Println("没有找到这本书")
}

// 借书方法
func (a *Sixflower) BorrowBook() {
	var id int
	var borrower string
	fmt.Print("请输入图书id")
	fmt.Scan(&id)
	fmt.Print("请输入您的姓名: ")
	fmt.Scan(&borrower)
	for i, book := range a.Books {
		if book.ID == id && !book.Borrowed {
			a.Books[i].Borrowed = true    
			a.Books[i].Borrower = borrower
			a.SaveBooks()
			fmt.Println("已经成功借的此书")
			return
		}
	}
	fmt.Println("此书不存在.")
}

// 归还图书的方法
func (a *Sixflower) ReturnBook() {
	var id int
	fmt.Print("请输入图书id ")
	fmt.Scan(&id)
	for i, book := range a.Books {
		if book.ID == id && book.Borrowed {
			a.Books[i].Borrowed = false
			a.Books[i].Borrower = ""
			a.SaveBooks()
			fmt.Println("图书已经归还")
			return
		}
	}
	fmt.Println("此书不存在")
}

// 列出此时的图书详情
func (a *Sixflower) ListBooks() {
	for _, book := range a.Books {
		status := "可借"
		if book.Borrowed {
			status = " 已借出 " + book.Borrower
		}
		fmt.Printf("ID: %d, Title: %s, Author: %s, Status: %s\n", book.ID, book.Title, book.Author, status)
	}
}

// 加载数据
func (a *Sixflower) LoadBooks() {
	data, err := os.ReadFile(a.FileName)
	if err != nil {
		fmt.Println("读取文件错误:", err)
		return
	}
	json.Unmarshal(data, &a.Books)
}

// 存储数据为文件名
func (a *Sixflower) SaveBooks() {
	data, err := json.MarshalIndent(a.Books, "", " ")
	if err != nil {
		fmt.Println("Error marshalling data:", err)
		return
	}
	err = os.WriteFile(a.FileName, data, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
}
// A 接口定义了管理书籍所需的方法。
type A interface {
	AddBook()        // 添加
	DeleteBook()      // 删除
	BorrowBook()     // 借阅
	ReturnBook()     // 归还
	ListBooks()       // 显示
	LoadBooks()      // 加载
	SaveBooks()       // 保存
}


func main() {
	manager := NewSixflower("Books.json")// 创建书籍管理器实例
	manager.LoadBooks()     //加载书记数据
	for {
		fmt.Println("\nOptions:")
		fmt.Println("1. Add Book")
		fmt.Println("2. Delete Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Books")
		fmt.Println("6. Exit")
		var choice string
		fmt.Scan(&choice)
		switch choice {
		case "1":
			manager.AddBook()
		case "2":
			manager.DeleteBook()
		case "3":
			manager.BorrowBook()
		case "4":
			manager.ReturnBook()
		case "5":
			manager.ListBooks()
		case "6":
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("输入无效 请再试一次")
		}
	}
}