package main
import ( 
  "log"
  "fmt" 
  "net/http"
  "os"
  "encoding/json"
)

type Name struct {
  first string
  last string
}

type Answer struct {
  email string
  name Name
  website string
  github_repo_link string
}

func determineListenAddress() (string, error) {
  port := os.Getenv("PORT") 
  if port == "" {
    return "", fmt.Errorf("$PORT not set") 
  } 
  return ":" + port, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=utf-8")
  w.Header().Set("Access-Control-Allow-Origin", "*")
  if origin := r.Header.Get("Origin"); origin != "" {
    w.Header().Set("Access-Control-Allow-Origin", origin)
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers",
    "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
  }

  name := Name{}
  name.first = "Jarnin"
  name.last = "Fang"

  answer := Answer{}
  answer.email = "jarninfang@gmail.com"
  answer.website = "https://jarninfang.github.io"
  answer.github_repo_link = "https://google.com"
  answer.name = name
  //answerJson,_ := json.Marshal(answer)
  json.NewEncoder(w).Encode(answer)
  w.WriteHeader(http.StatusCreated)
  //w.Write(answerJson)
}

func main() {
  addr, err := determineListenAddress()
  if err != nil {
    log.Fatal(err)
  }
  http.HandleFunc("/",handler)
  log.Printf("Listening on %s...\n", addr)
  if err := http.ListenAndServe(addr, nil); err != nil {
    panic(err)
  }
}
