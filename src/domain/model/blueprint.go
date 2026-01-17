package model

type Blueprint struct {
	BaseModel
	StandardName    string          `gorm:"size:255"`
	Type            string          `gorm:"size:100"`
	Description     string          `gorm:"size:1000"`
	Functionalities []Functionality `gorm:"foreignKey:BlueprintID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Libraries       []Library       `gorm:"many2many:blueprint_libraries"`
}

type Functionality struct {
	BaseModel
	Category           string `gorm:"size:100"`
	Type               string `gorm:"size:100"`
	Provider           string `gorm:"size:100"`
	ImplementsGenerics bool
	FilePathsCSV       string `gorm:"size:1000"`
	BlueprintID        uint
	Operations         []FunctionalOperation `gorm:"foreignKey:FunctionalityID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type FunctionalOperation struct {
	BaseModel
	Name            string `gorm:"size:255"`
	Description     string `gorm:"size:1000"`
	FunctionalityID uint
}
