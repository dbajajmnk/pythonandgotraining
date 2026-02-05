package main

import (
	"log"
	"log/slog"
	"os"
)

func main() {
	log.SetPrefix("Team 5 App")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("FIRST PART IS DONE")

	file, err := os.OpenFile("mylogs.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0664)



	if err != nil {
		panic(err)
	}
	defer file.Close()

	log.SetOutput(file)
	log.Println("This logs goes into File")

	handler := slog.NewJSONHandler(os.Stdout,&slog.HandlerOptions{Level: slog.LevelInfo})
	logger := slog.New(handler)

	request_id :="my_request_11"
	logger.With("my_app_id",request_id)
	

}
