package controllers

import (
	// "fmt"
	"gorm.io/datatypes"
	// "encoding/json"
	"app0/db"
	"app0/models"
	"net/http"

	"github.com/gin-gonic/gin"


	// "bytes"
	"os"
	"image"
	"image/jpeg"
	"image/png"
	// "path/filepath"
	// "strings"
	"github.com/nfnt/resize"

	"path/filepath"
)

func CreateStudent(c *gin.Context){
	var student models.Student
	
	if err := c.ShouldBindBodyWithJSON(&student);err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : err.Error(),
		})
		return
	}

	if err := db.DB.Create(&student).Error;err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": "Failed to create student",
		})
		return
	}

	c.JSON(http.StatusCreated,gin.H{
		"message":"Student created successfuly",
		"data": student,
	})
}

func CreateTrainer(c *gin.Context){
	var trainer models.Trainer
	if err := c.ShouldBindBodyWithJSON(&trainer);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return 
	}
	if err := db.DB.Create(&trainer).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to create trainer",	
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"Trainer created successfully",
		"data" : trainer,
	})
}

func CreateEnquiry(c *gin.Context){
	var enquiry models.Enquiry
	if err := c.ShouldBindBodyWithJSON(&enquiry);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return 
	}
	if err := db.DB.Create(&enquiry).Error; err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Failed to create trainer",	
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":"Trainer created successfully",
		"data" : enquiry,
	})
}


func GetAllTrainersForStudent(c *gin.Context) {
	studentID := c.Param("id")

	var student struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Batches datatypes.JSON `json:"batches"`
	}
	if err := db.DB.Raw("SELECT id, name, batches FROM students WHERE id = $1", studentID).Scan(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	type TrainerDetails struct {
		ID        uint            `json:"id"`
		Name      string          `json:"name"`
		Email     datatypes.JSON `json:"email"`
		Mobile    uint            `json:"mobile"`
		Password  string          `json:"password"`
		Branch    string          `json:"branch"`
		Location  datatypes.JSON `json:"location"`
		Expertise datatypes.JSON `json:"expertise"`
	}

	var trainers []TrainerDetails

	query := `
SELECT DISTINCT
	t.id,
	t.name,
	t.email,
	t.mobile,
	t.password,
	t.branch,
	t.location,
	t.expertise
FROM students s
LEFT JOIN LATERAL jsonb_array_elements(s.batches) AS batch ON true
LEFT JOIN trainers t
	ON t.name = batch ->> 'trainer'
WHERE s.id = $1
ORDER BY t.name
`

	if err := db.DB.Raw(query, studentID).Scan(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student_id":   student.ID,
		"student_name": student.Name,
		"batches": student.Batches,
		"trainers":     trainers,
	})
}

// func Getstudents(c *gin.Context) {
// 	id := c.Param("id")

// 	if id != "" {
// 		var stud []models.Student

// 		if err := db.DB.First(&stud , id).Error; err != nil{
// 			c.JSON(http.StatusNotFound, gin.H{
// 				"error":"Student not found",
// 			})
// 			return
// 		}
// 		c.JSON(http.StatusOK,stud)
// 		return
// 	}
// 	var stud []models.Student
// 	if err := db.DB.Find(&stud).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError,gin.H{
// 			"error":err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK,stud)
// }



// func UploadImage(c *gin.Context) {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Ensure uploads directory exists
// 	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// File path
// 	path := filepath.Join("uploads", file.Filename)

// 	// Save file ONLY ONCE
// 	if err := c.SaveUploadedFile(file, path); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Save metadata in DB
// 	imageData := models.Image{
// 		FileName: file.Filename,
// 		FilePath: path,
// 		FileType: file.Header.Get("Content-Type"),
// 	}

// 	if err := db.DB.Create(&imageData).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message":   "Uploaded successfully",
// 		"id":        imageData.ID,
// 		"filename":  file.Filename,
// 		"full_url":  "/image/" + file.Filename,
// 		"thumb_url": "/thumbnail/" + file.Filename,
// 	})
// }

// // FullImage serves full image
// func FullImage(c *gin.Context) {
// 	filename := c.Param("filename")
// 	path := filepath.Join("uploads", filename)
// 	c.File(path)
// }

// // Thumbnail serves 100x100 thumbnail
// func Thumbnail(c *gin.Context) {
// 	filename := c.Param("filename")
// 	path := filepath.Join("uploads", filename)

// 	file, err := os.Open(path)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
// 		return
// 	}
// 	defer file.Close()

// 	img, format, err := image.Decode(file)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image"})
// 		return
// 	}

// 	thumb := resize.Thumbnail(100, 100, img, resize.Lanczos3)
// 	buf := new(bytes.Buffer)
// 	if strings.ToLower(format) == "png" {
// 		png.Encode(buf, thumb)
// 		c.Data(http.StatusOK, "image/png", buf.Bytes())
// 	} else {
// 		jpeg.Encode(buf, thumb, nil)
// 		c.Data(http.StatusOK, "image/jpeg", buf.Bytes())
// 	}
// }

func ImageHandler(c *gin.Context) {

	// ==========================
	// 👉 UPLOAD IMAGE (POST)
	// ==========================
	if c.Request.Method == "POST" {

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": "file is required"})
			return
		}

		// Prevent duplicate
		var existing models.Image
		if err := db.DB.Where("file_name = ?", file.Filename).
			First(&existing).Error; err == nil {
			c.JSON(409, gin.H{"error": "Image already exists"})
			return
		}

		os.MkdirAll("uploads/thumbs", os.ModePerm)

		origPath := filepath.Join("uploads", file.Filename)
		thumbPath := filepath.Join("uploads/thumbs", file.Filename)

		// Save original
		if err := c.SaveUploadedFile(file, origPath); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Create thumbnail
		src, _ := os.Open(origPath)
		defer src.Close()

		img, format, err := image.Decode(src)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid image"})
			return
		}

		thumb := resize.Thumbnail(100, 100, img, resize.Lanczos3)

		thumbFile, _ := os.Create(thumbPath)
		defer thumbFile.Close()

		if format == "png" {
			png.Encode(thumbFile, thumb)
		} else {
			jpeg.Encode(thumbFile, thumb, nil)
		}

		// Save DB
		imageData := models.Image{
			FileName:  file.Filename,
			FilePath:  origPath,
			ThumbPath: thumbPath,
			FileType:  file.Header.Get("Content-Type"),
		}

		db.DB.Create(&imageData)

		c.JSON(200, gin.H{
			"message": "uploaded successfully",
			"image":   imageData,
		})
		return
	}

	// ==========================
	// 👉 DISPLAY IMAGE (GET)
	// ==========================
	filename := c.Param("filename")
	showThumb := c.Query("thumb") == "true"

	var img models.Image
	if err := db.DB.Where("file_name = ?", filename).
		First(&img).Error; err != nil {
		c.JSON(404, gin.H{"error": "image not found"})
		return
	}

	path := img.FilePath
	if showThumb {
		path = img.ThumbPath
	}

	c.File(path)
}




