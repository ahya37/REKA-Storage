package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"reka-storage/src/auth"
	"reka-storage/src/shared/middleware"
	"reka-storage/src/storage"
	"reka-storage/src/user"
	"reka-storage/src/user/repositories"
	"reka-storage/src/user/services"
)

func main() {
	// ===============================
	// Load ENV
	// ===============================
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file loaded")
	}

	// ===============================
	// Init Gin
	// ===============================
	r := gin.Default()

	// ===============================
	// MongoDB Connection
	// ===============================
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(os.Getenv("MONGO_URI")),
	)
	if err != nil {
		log.Fatal("MongoDB connection failed:", err)
	}

	db := mongoClient.Database(os.Getenv("MONGO_DB"))
	log.Println("MongoDB connected")

	// ===============================
	// MinIO Client
	// ===============================
	minioEndpoint := os.Getenv("REKASTORAGE_HOST") + ":" + os.Getenv("REKASTORAGE_PORT")
	useSSL := os.Getenv("REKASTORAGE_USE_SSL") == "true"

	minioClient, err := minio.New(
		minioEndpoint,
		&minio.Options{
			Creds: credentials.NewStaticV4(
				os.Getenv("REKASTORAGE_ACCESS_KEY"),
				os.Getenv("REKASTORAGE_SECRET_KEY"),
				"",
			),
			Secure: useSSL,
		},
	)
	if err != nil {
		log.Fatal("MinIO connection failed:", err)
	}

	log.Println("MinIO connected:", minioEndpoint)

	//=== START Storage Wiring
	fileRepo := storage.NewRepository(db) // db: *mongo.Database ✔
	fileService := storage.NewService(
		minioClient,
		os.Getenv("REKASTORAGE_BUCKET"),
		fileRepo,
	)
	fileHandler := storage.NewHandler(fileService)
	//=== END Storage Wiring

	//=== START Auth Wiring
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)
	//=== END Auth Wiring

	//=== START User Wiring
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	//=== END User Wiring

	//=== START API Routes
	api := r.Group("/api")
	{
		//AUTH
		authGroup := api.Group("/auth")
		authGroup.POST("/login", authHandler.Login)

		//USER
		userGroup := api.Group("/user")
		userGroup.Use(middleware.AuthMiddleware())
		userGroup.GET("/profile", userHandler.Profile)
		userGroup.POST("/register", userHandler.Register)

		// STORAGE
		storageGroup := api.Group("/storage")
		storageGroup.Use(middleware.AuthMiddleware())
		storage.RegisterRoutes(storageGroup, fileHandler)
	}
	//=== END API Routes

	//=== START Health Check
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	//=== END Health Check

	//=== START Run Server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "4000"
	}
	log.Println("Server running at :" + port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
