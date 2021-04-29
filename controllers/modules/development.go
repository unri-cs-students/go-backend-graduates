package modules

import (
	"github.com/SemmiDev/go-backend/controllers/middlewares"
	"github.com/SemmiDev/go-backend/entities"
	"github.com/SemmiDev/go-backend/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Development(db *gorm.DB, g *gin.Engine) {
	g.GET("/dummy", middlewares.Auth, func(c *gin.Context) {
		InsertDummy(db)
	})
	g.GET("/test", func(c *gin.Context) {
		repo := repositories.NewViewRepository(db)
		//results := []map[string]interface{}{}
		//var results []entities.Graduates
		// oid := c.Param("id")
		// ids := db.Table("wisudawan").Joins("INNER JOIN content ON content.wisudawan_id = wisudawan.id").Joins("INNER JOIN organization ON organization.id = content.organization_id").Where("organization_id = ?", oid).Distinct("wisudawan.id")
		// db.Preload(clause.Associations).Find(&results, "id IN (?)", ids)
		// c.JSON(http.StatusOK, results)
		// db.Find(&results).Order("nim desc").Limit(2).Find(&results)
		result, _ := repo.GetLast("a", "1")
		c.JSON(http.StatusOK, result)

	})
}

func InsertDummy(db *gorm.DB) {
	var MajorDummy [3]entities.Major
	MajorDummy[0] = entities.Major{
		Major:       	"Teknik Informatika",
		Faculty:      	"Sekolah Teknik Elektro Informatika",
		FacultyShort: 	"ILKOM",
		MajorShort:  	"HMIF",
	}
	MajorDummy[1] = entities.Major{
		Major:       "Teknik Lingkungan",
		Faculty:      "Faculty Teknik Sipil dan Lingkungan",
		MajorShort:  "TL",
		FacultyShort: "FTSL",
	}

	MajorDummy[2] = entities.Major{
		Major:      	"Matematika",
		Faculty:      	"Faculty Matematika dan Ilmu Pengetahuan Alam",
		FacultyShort: 	"FMIPA",
		MajorShort:  	"MA",
	}
	db.Create(&MajorDummy)

	var MajorId [3]string
	MajorId[0] = MajorDummy[0].ID
	MajorId[1] = MajorDummy[1].ID
	MajorId[2] = MajorDummy[2].ID

	var OrganizationDummy [3]entities.Orgz
	OrganizationDummy[0] = entities.Orgz{
		Name:          	"Himpunan Mahasiwswa Teknik Informatika",
		Slug:          	"HMIF",
		Category:      	"HMJ",
		Logo:          	"/path/to/logo",
		FacultyShort: 	"ILKOM",
	}
	OrganizationDummy[1] = entities.Orgz{
		Name:     "Kabinet",
		Slug:     "Kabinet",
		Category: "Kabinet_KM_UNRI",
		Logo:     "/path/to/logo",
	}
	OrganizationDummy[2] = entities.Orgz{
		Name:          	"Himpunan Mahasiswa Elektro",
		Slug:          	"HME",
		Category:      	"HMJ",
		Logo:          	"/path/to/logo",
		FacultyShort: 	"ILKOM",
	}

	db.Create(&OrganizationDummy)

	var OrganizationId [3]string
	OrganizationId[0] = OrganizationDummy[0].ID
	OrganizationId[1] = OrganizationDummy[1].ID
	OrganizationId[2] = OrganizationDummy[2].ID

	var GraduatesDummy [3]entities.Graduates
	GraduatesDummy[0] = entities.Graduates{
		Identifier:     13519000,
		Name:         	"Sebuah nama",
		NickName:    	"Sebuah panggilan",
		ThesisTitle:    "Belum TA",
		Incoming:     	16,
		MajorID:    	MajorId[0],
		Instagram:    	"gak ada",
		DateOfBirth: 	time.Now(),
		Photo:        	"path/to/foto",
	}
	GraduatesDummy[1] = entities.Graduates{
		Identifier:     13519001,
		Name:         	"Name Lain",
		NickName:    	"NickName Lain",
		ThesisTitle:    "Sudah TA",
		Incoming:     	15,
		MajorID:    	MajorId[1],
		Instagram:    	"ada",
		DateOfBirth: 	time.Now(),
		Photo:        	"path/to/photo",
	}
	GraduatesDummy[2] = entities.Graduates{
		Identifier:     13519002,
		Name:         	"Name Lain lagi",
		NickName:    	"NickName Lain lagi",
		ThesisTitle:    "Sudah TA lagi",
		Incoming:     	14,
		MajorID:    	MajorId[2],
		Instagram:    	"sudah ada",
		DateOfBirth: 	time.Now(),
		Photo:        	"/path/to/newphoto",
	}
	db.Create(&GraduatesDummy)

	var idW [3]string
	idW[0] = GraduatesDummy[0].ID
	idW[1] = GraduatesDummy[1].ID
	idW[2] = GraduatesDummy[2].ID

	var MessageDummy [3]entities.Message
	MessageDummy[0] = entities.Message{
		ReceiverID: idW[0],
		Message:    "Halo",
		Sender:     "Anon",
	}

	MessageDummy[1] = entities.Message{
		ReceiverID: idW[0],
		Message:    "Miaw",
		Sender:     "Kucing UNRI",
	}

	MessageDummy[2] = entities.Message{
		ReceiverID: idW[2],
		Message:    "Halo sayang",
		Sender:     "Secret Admirer",
	}
	db.Create(&MessageDummy)

	var ContentDummy [5]entities.Content

	ContentDummy[0] = entities.Content{
		GraduatesID: idW[0],
		Type:        "PRESTASI",
		Headings:    "Imba aku cuk",
	}
	ContentDummy[1] = entities.Content{
		GraduatesID: idW[0],
		Type:        "TIPS_SUKSES",
		Headings:    "Swimming aja",
		Details:     "Berenang menyehatkan badan",
	}
	ContentDummy[2] = entities.Content{
		GraduatesID:    idW[0],
		OrganizationID: OrganizationId[1],
		Type:           "KONTRIBUSI",
		Headings:       "Swimming aja",
	}
	ContentDummy[3] = entities.Content{
		GraduatesID:    idW[2],
		OrganizationID: OrganizationId[0],
		Type:           "KONTRIBUSI",
		Headings:       "Swimming aja",
	}
	ContentDummy[4] = entities.Content{
		GraduatesID:    idW[2],
		OrganizationID: OrganizationId[0],
		Type:           "KONTRIBUSI",
		Headings:       "Lalala",
	}
	db.Create(&ContentDummy)
}