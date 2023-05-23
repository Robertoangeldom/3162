// main.go
package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/MejiaFrancis/3161/3162/quiz-2/recsystem/internal/models" // this will be change to file from another folder models
	"github.com/alexedwards/scs/v2"                                      // needs to be downladed when I have internet
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Share data across our handlers

type application struct {
	errorLog       *log.Logger
	infoLog        *log.Logger
	reservations   models.ReservationModel
	user           models.UserModel
	equipments     models.EquipmentModel
	feedback       models.FeedbackModel
	sessionManager *scs.SessionManager
}

func main() {
	// configure our server
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("RCSYSTEM_DB_DSN"), "PostgreSQL DSN (Data Source Name)")
	flag.Parse()

	// get a database connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Print(err)
		return
	}
	//create instances of error log and info log
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//set up a new session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode

	// share data across our handlers
	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		reservations:   models.ReservationModel{DB: db},
		user:           models.UserModel{DB: db},
		equipments:     models.EquipmentModel{DB: db},
		feedback: models.FeedbackModel{DB: db},
		sessionManager: sessionManager,
	}
	// cleanup the connection pool
	defer db.Close()
	// acquired a database connection pool
	infoLog.Printf("database connection pool established")
	// create and start a custom web server
	infoLog.Printf("starting server on %s", *addr)
	//configure TLS
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	//err = srv.ListenAndServe()
	log.Fatal(err)
}

// The openDB() function returns a database connection pool or error
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// create a context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// test the DB connection
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}
