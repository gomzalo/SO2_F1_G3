package configs

import "os"

func EnvMongoUser() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return os.Getenv("MONGO_USER")
}

func EnvMongoPass() string {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	return os.Getenv("MONGO_PASS")
}