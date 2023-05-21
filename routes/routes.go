package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mthsnts/go-API/controllers"
	"github.com/mthsnts/go-API/middleware"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	http.Handle("/", r)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	r.HandleFunc("/api/personalidades/new", controllers.CriaUmaNovaPersonalidade).Methods("Post")
	r.HandleFunc("/api/personalidades/delete/{id}", controllers.RemoveUmaPersonalidade).Methods("Delete")
	r.HandleFunc("/api/personalidades/edit/{id}", controllers.EditaPersonalidade).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}

//Deodato Petit Wertheimer foi um médico e político brasileiro, seus primeiros anos de vida foram em São Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Colégio Anchieta dos jesuítas.
