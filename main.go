package main

import (
	"context"
	"flag"
	"fmt"
	_ "image/jpeg"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
)

var bucket = flag.String("b", " ", "Name of bucket")
var file = flag.String("f", " ", "File to upload")

func main() {
	flag.Parse()
	fmt.Printf("%s %s", *bucket, *file)

	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Printf("Error creating storage clinet %s\n", err.Error())
		return
	}

	// Check the bucket's attributes
	// bkt := client.Bucket("gol-dev-bucket")
	bkt := client.Bucket(*bucket)
	attrs, err := bkt.Attrs(ctx)
	if err != nil {
		log.Printf("Error getting bucket attributes %s\n", err.Error())
		return
	}

	fmt.Printf("Bucket: %s\n Created at %s\n Located in %s\n Storage class %s\n", attrs.Name, attrs.Created, attrs.Location, attrs.StorageClass)

	// Upload image to storage
	// file, err := os.Open("wave.jpg")
	file, err := os.Open(*file)
	obj := bkt.Object("wave.jpg")
	w := obj.NewWriter(ctx)
	io.Copy(w, file)
	if err := w.Close(); err != nil {
		log.Printf("Cannot close writer %s\n", err.Error())
		return
	}

	log.Println("OK writing to storage!")

}
