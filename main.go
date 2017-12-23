package main
import ( 
  "log"
  "fmt" 
  "net/http"
  "os"
  "encoding/json"
)

type Name struct {
  First string `json:"first"`
  Last string  `json:"last"`
}

type Answer struct {
  Email string              `json:"email"`
  Name Name                 `json:"name"`
  Website string            `json:"website"`
  Github_repo_link string   `json:"github_repo_link"`
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

  name := Name{}
  name.First = "Jarnin"
  name.Last = "Fang"

  answer := Answer{}
  answer.Email = "jarninfang@gmail.com"
  answer.Website = "https://jarninfang.github.io"
  answer.Github_repo_link = "https://google.com"
  answer.Name = name

  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(answer)
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
