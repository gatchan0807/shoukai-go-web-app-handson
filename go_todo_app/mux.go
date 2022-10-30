package main

import (
	"net/http"

	"github.com/gatchan0807/go_todo_app/handler"
	"github.com/gatchan0807/go_todo_app/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator"
)

func NewMux() http.Handler {
	// memo: MUXってなに？ => もともと電子回路の文脈でふたつの信号を1つにまとめる装置の意味で使われる用語。　( https://wa3.i-3-i.info/word13228.html )
	// Goでは "ServerMux is an HTTP request multiplexer." と表現されるように、ルーティングを担当する処理が ServeMux に詰め込まれている
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するため明示的に戻り値を捨てている
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	// memo: add-task の略の at
	at := &handler.AddTask{Store: store.Tasks, Validator: v}
	mux.Post("/tasks", at.ServeHTTP)
	// memo: list-task の略の lt
	lt := &handler.ListTask{Store: store.Tasks}
	mux.Get("/tasks", lt.ServeHTTP)

	return mux
}
