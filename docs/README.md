## EORE 程式碼產生程式


### Schema 設計限制
* 系統自動設定日期
```
createDate	// 建立日期
lastupdate	// 更新日期
```


### 前端限制
* 更新
```
// 允許一次更新多筆，因此即便只有一筆資料也需要送出陣列
[{
   ...
}]
```


### 相關資訊
* 若樣板檔有註解可以參考：https://colobu.com/2019/11/05/Golang-Templates-Cheatsheet/
* [How to build your first web application with Go](https://freshman.tech/web-development-with-go/)
