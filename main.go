package main
import (
  "log"
  "fmt"
  "net/http"
  "os"
  //"encoding/json"
  "github.com/gorilla/mux"
)

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT")
  if port == "" {
    return "", fmt.Errorf("$PORT not set")
  }
  return ":" + port, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  //m := map[string]string{
  //  "email": "jarninfang@gmail.com",
  //  "website": "jarninfang.github.io",
  //}
  fmt.Fprintln(w, "Hello World")
  //w.Header().Add("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  //_ = json.NewEncoder(w).Encode(m)
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  router.HandleFunc("/", handler)
  http.Handle("/", router)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, router); err != nil {
    panic(err)
  }
}
