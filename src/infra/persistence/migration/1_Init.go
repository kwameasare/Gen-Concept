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

var logger = logging.NewLogger(config.GetConfig())

func Up1() {
	database := database.GetDb()

	createTables(database)
	createDefaultUserInformation(database)

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
	tables = addNewTable(database, models.Entity{}, tables)
	tables = addNewTable(database, models.DependsOnEntity{}, tables)
	tables = addNewTable(database, models.EntityField{}, tables)
	tables = addNewTable(database, models.InputValidation{}, tables)

	//Journey
	tables = addNewTable(database, models.Journey{}, tables)
	tables = addNewTable(database, models.EntityJourney{}, tables)
	tables = addNewTable(database, models.Operation{}, tables)
	tables = addNewTable(database, models.JourneyStep{}, tables)
	tables = addNewTable(database, models.FieldInvolved{}, tables)
	tables = addNewTable(database, models.RetryCondition{}, tables)
	tables = addNewTable(database, models.ResponseAction{}, tables)
	tables = addNewTable(database, models.Filter{}, tables)
	tables = addNewTable(database, models.Sort{}, tables)

	//Blueprint
	tables = addNewTable(database, models.Blueprint{}, tables)
	tables = addNewTable(database, models.Functionality{}, tables)
	tables = addNewTable(database, models.FunctionalOperation{}, tables)

	// Library
	tables = addNewTable(database, models.Library{}, tables)
	tables = addNewTable(database, models.LibraryFunctionality{}, tables)
	tables = addNewTable(database, models.BlueprintLibrary{}, tables)

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
	fmt.Print("\n HEEEEEEEERRRRRRRRREEEEE >>>>>>>>>>>>>>>>>>>>>>>>>>>\n")
	fmt.Printf("\n Checking table %s", model)
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

func Down1() {
	// nothing
}
