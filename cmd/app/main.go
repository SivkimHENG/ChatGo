package main


import (
    "github.com/go-chi/chi/v5"
    "net/http"
    "github.com/joho/godotenv"
    "os"
    "chatMessages/handlers"
    "log"
    "golang.org/x/exp/slog"
)


func main(){

    if err :=  godotenv.Load(); err != nil {
      log.Fatal(err);
    }


router := chi.NewMux();

router.Get("/home-page", handlers.homeHandler);
    listenAddr := os.Getenv("LISTEN_ADDR");
    slog.Info("HTTP server started", "listenAddr", listenAddr);
    http.ListenAndServe(listenAddr, router);





}
