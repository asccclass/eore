// router.go
package main

import(
   "github.com/gorilla/mux"
   "github.com/asccclass/staticfileserver"
   "github.com/asccclass/serverstatus"
   "./libs/template"
)

// Create your Router function
func NewRouter(srv *SherryServer.ShryServer, documentRoot string)(*mux.Router) {
   router := mux.NewRouter()

   // Add your router here.....
/*
   // Takeinlist 
   tml, err := sherryTemplate	
   if err != nil {
      panic(err)
   }
   tpl.AddTemplateRouter(router)
*/

   //logger
   router.Use(SherryServer.ZapLogger(srv.Logger))

   // health check
   m := serverstatus.NewServerStatus()
   router.HandleFunc("/healthz", m.Healthz).Methods("GET")

   // Static File server
   staticfileserver := SherryServer.StaticFileServer{documentRoot, "index.html"}
   router.PathPrefix("/").Handler(staticfileserver)

   return router
}

