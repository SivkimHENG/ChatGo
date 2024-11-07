package handlers;


import (

    "fmt"
    "net/http"
    "chatMessages/view/homepage"
)


func homeHandler (w http.ResponseWriter, r *http.Request){
    return homepage.Index().Render(r.Context(), w);
}
