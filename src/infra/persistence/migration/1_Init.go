package migration

import (
	"fmt"
	"gen-concept-api/config"
	"gen-concept-api/constant"
	models "gen-concept-api/domain/model"
	database "gen-concept-api/infra/persistence/database"
	"gen-concept-api/pkg/logging"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const countStarExp = "count(*)"

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := database.GetDb()

	createTables(database)
	createDefaultUserInformation(database)
	createPropertyCategory(database)

}

func createTables(database *gorm.DB) {
	tables := []interface{}{}
	database = database.Debug()

	// Basic
	tables = addNewTable(database, models.File{}, tables)
	// Property
	tables = addNewTable(database, models.PropertyCategory{}, tables)
	tables = addNewTable(database, models.Property{}, tables)

	// User
	tables = addNewTable(database, models.User{}, tables)
	tables = addNewTable(database, models.Role{}, tables)
	tables = addNewTable(database, models.UserRole{}, tables)

	//Project
	tables = addNewTable(database, models.Project{}, tables)
	tables= addNewTable(database, models.Entity{}, tables)
	tables= addNewTable(database, models.DependsOnEntity{}, tables)
	tables= addNewTable(database, models.EntityField{}, tables)
	tables= addNewTable(database, models.InputValidation{}, tables)
	
	er := database.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error
	if er != nil {
		logger.Error(logging.Postgres, logging.Migration, er.Error(), nil)
	}
	
	err := database.Migrator().AutoMigrate(tables...)
	if err != nil {
		logger.Error(logging.Postgres, logging.Migration, err.Error(), nil)
	}
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}

func addNewTable(database *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	fmt.Print( "\n HEEEEEEEERRRRRRRRREEEEE >>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	fmt.Printf( "\n Checking table %s", model)
	if !database.Migrator().HasTable(model) {
		tables = append(tables, model)
	}
	return tables
}

func createDefaultUserInformation(database *gorm.DB) {

	adminRole := models.Role{Name: constant.AdminRoleName}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: constant.DefaultRoleName}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{Username: constant.DefaultUserName, FirstName: "Test", LastName: "Test",
		MobileNumber: "233557262205", Email: "admin@admin.com"}
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)

	createAdminUserIfNotExists(database, &u, adminRole.ID)

}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0
	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	if exists == 0 {
		database.Create(r)
	}
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId uint) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.Username).
		First(&exists)
	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.ID, RoleId: roleId}
		database.Create(&ur)
	}
}



func createPropertyCategory(database *gorm.DB) {
	// count := 0

	// database.
	// 	Model(&models.PropertyCategory{}).
	// 	Select(countStarExp).
	// 	Find(&count)
	// if count == 0 {
	// 	database.Create(&models.PropertyCategory{Name: "Body"})                    
	// 	database.Create(&models.PropertyCategory{Name: "Engine"})                  
	// 	database.Create(&models.PropertyCategory{Name: "Drivetrain"})              
	// 	database.Create(&models.PropertyCategory{Name: "Suspension"})              
	// 	database.Create(&models.PropertyCategory{Name: "Equipment"})               
	// 	database.Create(&models.PropertyCategory{Name: "Driver support systems"})  
	// 	database.Create(&models.PropertyCategory{Name: "Lights"})                  
	// 	database.Create(&models.PropertyCategory{Name: "Multimedia"})              
	// 	database.Create(&models.PropertyCategory{Name: "Safety equipment"})        
	// 	database.Create(&models.PropertyCategory{Name: "Seats and steering wheel"})
	// 	database.Create(&models.PropertyCategory{Name: "Windows and mirrors"})      
	// }
	// createProperty(database, "Body")
	// createProperty(database, "Engine")
	// createProperty(database, "Drivetrain")
	// createProperty(database, "Suspension")
	// createProperty(database, "Comfort")
	// createProperty(database, "Driver support systems")
	// createProperty(database, "Lights")
	// createProperty(database, "Multimedia")
	// createProperty(database, "Safety equipment")
	// createProperty(database, "Seats and steering wheel")
	// createProperty(database, "Windows and mirrors")

}

// func createProperty(database *gorm.DB, cat string) {
// 	count := 0
// 	catModel := models.PropertyCategory{}

// 	database.
// 		Model(models.PropertyCategory{}).
// 		Where("name = ?", cat).
// 		Find(&catModel)

// 	database.
// 		Model(&models.Property{}).
// 		Select(countStarExp).
// 		Where("category_id = ?", catModel.Id).
// 		Find(&count)

// 	if count > 0 || catModel.Id == 0 {
// 		return
// 	}
// 	var props *[]models.Property
// 	switch cat {
// 	case "Body":
// 		props = getBodyProperties(catModel.Id)

// 	case "Engine":
// 		props = getEngineProperties(catModel.Id)

// 	case "Drivetrain":
// 		props = getDrivetrainProperties(catModel.Id)

// 	case "Suspension":
// 		props = getSuspensionProperties(catModel.Id)

// 	case "Comfort":
// 		props = getComfortProperties(catModel.Id)

// 	case "Driver support systems":
// 		props = getDriverSupportSystemProperties(catModel.Id)

// 	case "Lights":
// 		props = getLightsProperties(catModel.Id)

// 	case "Multimedia":
// 		props = getMultimediaProperties(catModel.Id)

// 	case "Safety equipment":
// 		props = getSafetyEquipmentProperties(catModel.Id)

// 	case "Seats and steering wheel":
// 		props = getSeatsProperties(catModel.Id)

// 	case "Windows and mirrors":
// 		props = getWindowsProperties(catModel.Id)

// 	default:
// 		props = &([]models.Property{})
// 	}

// 	for _, prop := range *props {
// 		database.Create(&prop)
// 	}
// }

func Down1() {
	// nothing
}
