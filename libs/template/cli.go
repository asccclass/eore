package main

import(
   "os"
   "fmt"
   "text/template"
)

func main() {
   dbName := "laitaian"
   table := "takeinlistobj"
   tpl, err := NewTemplate(dbName, table)
   if err != nil {
      fmt.Println(err)
      return 
   }
   tmpl := template.Must(template.ParseFiles("./frontend.tmpl"))
   params := tpl.ProcessSchema()
   // tmpl.Execute(f, params)
   tmpl.Execute(os.Stdout, params)	// write to command line
   // f.Close()
return
/*
   // 判斷目錄是否存在
   dir := "./" + dbName + "/" + table
   if _, err := os.Stat(dir); os.IsNotExist(err) {
      if err = os.MkdirAll(dir, 0755); err != nil {
         panic(err)
      }
   }
   f, err := os.Create(dir + "/backend/" + table + ".go")
   if err != nil {
      fmt.Errorf("Create %s error:%s\n", dir, err.Error())
      return
   }
   tmpl := template.Must(template.ParseFiles("./crud.tmpl"))
   params := tpl.ProcessSchema()
   tmpl.Execute(f, params)
   // tmpl.Execute(os.Stdout, params)	// write to command line
   // tmpl.Execute(w, nil)		// write to web
   f.Close()

   f, err = os.Create(dir + "/backend/" + table + "web.go")
   if err != nil {
      fmt.Errorf("Create %s error:%s\n", dir, err.Error())
      return
   }
   tmpl = template.Must(template.ParseFiles("./webinterface.tmpl"))
   params = tpl.ProcessSchema()
   tmpl.Execute(f, params)
   f.Close()
*/
   // 前端
/*
   f, err = os.Create(dir + "/frontend/" + table + ".go")
   if err != nil {
      fmt.Errorf("Create %s error:%s\n", dir, err.Error())
      return
   }
*/
}
