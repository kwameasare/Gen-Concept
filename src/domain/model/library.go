package model

type Library struct {
	BaseModel
	Name                   string                 `gorm:"size:255;uniqueIndex" json:"standardName"`
	Version                string                 `gorm:"size:50" json:"version"`
	Description            string                 `gorm:"size:1000" json:"description"`
	RepositoryURL          string                 `gorm:"size:500" json:"repositoryURL"`
	Namespace              string                 `gorm:"size:255" json:"namespace"`
	ExposedFunctionalities []LibraryFunctionality `gorm:"foreignKey:LibraryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"exposedFunctionalities"`
	Blueprints             []Blueprint            `gorm:"many2many:blueprint_libraries" json:"blueprints,omitempty"`
	OrganizationID         *uint                  `json:"organizationID,omitempty"`
	Organization           *Organization          `json:"organization,omitempty"`
	TeamID                 *uint                  `json:"teamID,omitempty"`
	Team                   *Team                  `json:"team,omitempty"`
	// Advanced Discovery Fields
	GitReference string `gorm:"size:255" json:"gitReference"` // e.g., branch or specific ref
	CommitHash   string `gorm:"size:100" json:"commitHash"`
	Tag          string `gorm:"size:100" json:"tag"`
}

type LibraryFunctionality struct {
	BaseModel
	Name        string `gorm:"size:255" json:"name"`
	Type        string `gorm:"size:100" json:"type"` // Utility, Service, Helper, etc.
	Description string `gorm:"size:1000" json:"description"`
	LibraryID   uint   `json:"libraryID"`
}

type BlueprintLibrary struct {
	BaseModel
	BlueprintID     uint   `json:"blueprintID"`
	LibraryID       uint   `json:"libraryID"`
	RequiredVersion string `gorm:"size:50" json:"requiredVersion"` // Which version this blueprint requires
}

type LibraryDefinition struct {
	BaseModel
	PackageName  string   `gorm:"size:150" json:"packageName"`
	FunctionName string   `gorm:"size:150;index" json:"functionName"`
	Signature    string   `gorm:"type:text" json:"signature"`
	Description  string   `gorm:"size:1000" json:"description"`
	Tags         []string `gorm:"serializer:json" json:"tags"`
	RepoURL      string   `gorm:"size:500" json:"repoURL"`
}
