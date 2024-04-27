package config

import "os"

func DBUser() string {
	return os.Getenv("DB_USER")
}

func DBPassword() string {
	return os.Getenv("DB_PASSWORD")
}

func DBHost() string {
	return os.Getenv("DB_HOST")
}

func DBPort() string {
	return os.Getenv("DB_PORT")
}

func DBName() string {
	return os.Getenv("DB_NAME")
}

func AWSAccessKey() string {
	return os.Getenv("AWS_ACCESS_KEY")
}

func AWSSecretKey() string {
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func AWSRegion() string {
	return os.Getenv("AWS_REGION")
}

func AWSBucket() string {
	return os.Getenv("AWS_BUCKET_NAME")
}
