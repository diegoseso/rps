package server

import(
	"net"
	"log"
	"io"
	"time"
	"net/http"
	"os"
)

type Server struct{

}

func NewServer() *Server{
	return &Server{}
}

var addr string

func(S *Server)Run(string configPath){
	
	err := config.Load(configPath)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	loggerConfig := config.GetLoggerConfig()
	if loggerConfig.Level != "" {
		logger.SetLevel(loggerConfig.Level)
	}

	logger.Info("TangeloGame Middleware")
	
	c := GetSocketServerConfig()
    serverAddress := c.Host + ":" + c.Port + c.Path 

	hub := newHub()
	go hub.run()
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		LoginHandler(w, r)
	})
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		connect(hub, w, r)
	})

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleConn( c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}

func(S *Server)Stop(){
    os.Exit(1)
}