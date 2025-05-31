package main

import (
	"gameverse/configs"
	"gameverse/internal/auth"
	"gameverse/pkg/db"
	"net/http"
)

func main() {
	conf := configs.LoadConfig()
	_ = db.NewDb(conf)

	router := http.NewServeMux()
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config: conf,
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	println("Backend service is running...")

	server.ListenAndServe()
}

// Строка
// var s string = "hello"

// // Число
// var n int = 42
// var f float64 = 3.14

// // Булевое значение
// var b bool = true

// // Структура
// type User struct {
//     Name string
//     Age  int
// }
// u := User{Name: "Alex", Age: 30}

// // Массив (фиксированный размер)
// var arr [3]int = [3]int{1, 2, 3}

// // Срез (slice, динамический массив)
// sl := []int{4, 5, 6}

// // Срез из массива
// sl2 := arr[1:3] // элементы с 1 по 2

// // Карта (map, ассоциативный массив)
// m := map[string]int{"a": 1, "b": 2}

// // Канал
// ch := make(chan int)

// // Буферизированный канал
// chBuf := make(chan string, 5)

// // Горутина
// go func() {
//     println("Hello from goroutine")
// }()

// // Указатель
// var p *int = &n

// // Интерфейс
// type Stringer interface {
//     String() string
// }

// // Функция как значение
// fn := func(x int) int { return x * 2 }

// // Множество (set) — через map
// set := map[int]struct{}{1: {}, 2: {}}

// // Мьютекс (синхронизация)
// import "sync"
// var mu sync.Mutex

// // Время
// import "time"
// now := time.Now()
