package cloudinary

import (
	"context"
	"log"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Interface interface {
	UploadToCloudinary(file *multipart.FileHeader, merchant_name string) (string, error)
}

type cloudinaryClient struct {
    client *cloudinary.Cloudinary
}

func Init() Interface{
	cldSecret := os.Getenv("CLOUD_API_SECRET")
	cldName := os.Getenv("CLOUD_NAME")
	cldKey := os.Getenv("CLOUD_API_KEY")

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)

	if err != nil {
		log.Fatalf("Error connecting to cloudinary: %v", err)
	}

	return &cloudinaryClient{client: cld}
}


func(c *cloudinaryClient) UploadToCloudinary(file *multipart.FileHeader, merchant_name string) (string, error){
	filename := uuid.NewString()
	var ctx = context.Background()
	res, err := c.client.Upload.Upload(ctx, file, uploader.UploadParams{
			Folder: "/nbid-intermediate/" + merchant_name + "/products/",
			PublicID: filename,
			
		})

	if err != nil {
		return "", err
	}
	log.Println(res)
	// check if there are any eager in response
	if len(res.Eager) > 0 {
		// will return secure url with transformation
		return res.Eager[0].SecureURL, nil
	}

	// if no, will use secure url (without transformation)
	url := res.SecureURL

	return url, nil


} 