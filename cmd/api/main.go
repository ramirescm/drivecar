package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/go-chi/chi/v5"
	"github.com/ramirescm/drivecar/internal/auth"
	"github.com/ramirescm/drivecar/internal/bucket"
	"github.com/ramirescm/drivecar/internal/files"
	"github.com/ramirescm/drivecar/internal/folders"
	"github.com/ramirescm/drivecar/internal/queue"
	"github.com/ramirescm/drivecar/internal/users"
	"github.com/ramirescm/drivecar/pkg/database"
)

func main() {
	db, b, qc := getSessions()

	r := chi.NewRouter()

	r.Post("/auth", auth.HandlerAuth(func(login, password string) (auth.Authenticated, error) {
		return users.Authenticate(login, password)
	}))

	files.SetRoutes(r, db, b, qc)
	folders.SetRoutes(r, db)
	users.SetRoutes(r, db)

	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")), r)
}

func getSessions() (*sql.DB, *bucket.Bucket, *queue.Queue) {
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	// rabbit config
	qcfg := queue.RabbitMQConfig{
		URL:       os.Getenv("RABBIT_URL"),
		TopicName: os.Getenv("RABBIT_TOPIC_NAME"),
		Timeout:   time.Second * 30,
	}

	// create new queue
	qc, err := queue.New(queue.RabbitMQ, qcfg)
	if err != nil {
		log.Fatal(err)
	}

	bcfg := bucket.AwsConfig{
		Config: &aws.Config{
			Region:      aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_KEY"), os.Getenv("AWS_SECRET"), ""),
		},
		BucketDownload: "drivecar-raw",
		BucketUpload:   "drivercar-gzip",
	}

	// create new bucket session
	b, err := bucket.New(bucket.AwsProvider, bcfg)
	if err != nil {
		log.Fatal(err)
	}

	return db, b, qc
}
