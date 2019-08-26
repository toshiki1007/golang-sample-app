package main

import (
    "html/template"
    "log"
    "net/http"
    "time"
    "io"
    "os"
)

func clockHandler(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("/home/ssm-user/go/src/localhost/gohttpserver/templates/clock.html.tpl"))

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "clock.html.tpl", time.Now()); err != nil {
        log.Fatal(err)
    }
    log.Printf("[Info]: Response is successed.")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
    // テンプレートをパース
    t := template.Must(template.ParseFiles("/home/ssm-user/go/src/localhost/gohttpserver/templates/error.html.tpl"))

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "error.html.tpl", "Error!"); err != nil {
        log.Fatal(err)
    }
    log.Printf("[Error]: Response is failed.")
}

func main() {
    logfile, err := os.OpenFile("./app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        panic("cannnot open app.log:" + err.Error())
    }
    defer logfile.Close()

    // io.MultiWriteで、
    // 標準出力とファイルの両方を束ねて、
    // logの出力先に設定する
    log.SetOutput(io.MultiWriter(logfile, os.Stdout))

    log.SetFlags(log.Ldate | log.Ltime)

    // ハンドラーを登録
    http.HandleFunc("/clock", clockHandler)
    http.HandleFunc("/error", errorHandler)
    // 先ほどの静的ファイルハンドラー
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("/home/ssm-user/go/src/localhost/gohttpserver/html"))))
    // サーバーを起動
    log.Fatal(http.ListenAndServe(":8080", nil))