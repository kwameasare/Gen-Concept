package dto

// {
//     "projectName": "My Enterprise Project",//input
//     "projectDescription": "This is a description of my project", //input
//     "projectType": "Enterprise", //Select from a list of options
//     "isMultiTenant": true, //toggle boolean
//     "isMultiLingual": true, //toggle boolean
//     "entities": [
//         {
//             "entityName": "User",   //input
//             "entityDescription": "This is a description of my entity",  //input
//             "implementsRBAC": true, //toggle boolean
//             "isAuthenticationRequired": true,   //toggle boolean
//             "implementsAudit": true,    //toggle boolean
//             "implementsChangeManagement": true, //toggle boolean
//             "isReadOnly": false,    //toggle boolean
//             "isIndependentEntity": true,    //toggle boolean
//             "dependsOnEntities": [],    //select from a list of entities
//             "version": "1.0",   //input
//             "isBackendOnly": false, //toggle boolean
//             "preferredDB": "Postgres",  //select from a list of options
//             "modeOfDBInteraction": "ORM",   //select from ENUM
//             "entityFields": [
//                 {
//                     "fieldName": "name",  //input
//                     "displayName": "Name",  //input
//                     "fieldType": "String",  //select from ENUM
//                     "fieldDescription": "This is a description of my field",    //input
//                     "isMandatory": true,    //toggle boolean
//                     "isUnique": false,  //toggle boolean
//                     "isReadOnly": false,        //toggle boolean
//                     "isSensitive": false,   //toggle boolean
//                     "isEditable": true, //toggle boolean
//                     "isDerived": false, //toggle boolean
//                     "isCollection": false,    //toggle boolean
//                     "collectionType": "None", //select from ENUM
//                     "isEnum": false,    //toggle boolean
//                     "enumValues": [],   //input
//                     "derivativeType": "None",   //select from ENUM
//                     "derivativeExpression": "None", //input
//                     "isBackendOnly": false, //toggle boolean
//                     "displayStatus": "Show",    //select from ENUM
//                     "sampleData": "John Doe",   //input
//                     "inputValidation":{
//                         "description":"This is a description of my validation", //input
//                         "abortOnFailure":true,  //toggle boolean
//                         "customErrorMessage":"This is a custom error message"   //input

//                     }
//                 },
//                 {
//                     "fieldName": "email",   //input
//                     "displayName": "Email", //input
//                     "fieldType": "String",  //select from ENUM
//                     "fieldDescription": "This is a description of my field",    //input
//                     "isMandatory": true,    //toggle boolean
//                     "isUnique": false,  //toggle boolean
//                     "isReadOnly": false,    //toggle boolean
//                     "isSensitive": false,   //toggle boolean
//                     "isEditable": true,     //toggle boolean
//                     "isDerived": false,    //toggle boolean
//                     "isCollection": false,  //toggle boolean
//                     "collectionType": "None",
//                     "isEnum": false,    //toggle boolean
//                     "enumValues": [],   //input
//                     "derivativeType": "None",   //select from ENUM
//                     "derivativeExpression": "None", //input
//                     "isBackendOnly": false, //toggle boolean
//                     "displayStatus": "Detail",  //select from ENUM
//                     "sampleData":"john.doe@email.com",  //input
//                     "inputValidation":{
//                         "description":"This is a description of my validation", //input
//                         "abortOnFailure":true,  //toggle boolean
//                         "customErrorMessage":"This is a custom error message"   //input

//                     }
//                 },
//                 {
//                     "fieldName": "password",    //input
//                     "displayName": "Password",  //input
//                     "fieldType": "String", //select from ENUM
//                     "fieldDescription": "This is a description of my field",    //input
//                     "isMandatory": true,    //toggle boolean
//                     "isUnique": false,      //toggle boolean
//                     "isReadOnly": false,    //toggle boolean
//                     "isSensitive": true,    //toggle boolean
//                     "isEditable": true,    //toggle boolean
//                     "isDerived": false,   //toggle boolean
//                     "isCollection": false,  //toggle boolean
//                     "collectionType": "None",   //select from ENUM
//                     "isEnum": false,    //toggle boolean
//                     "enumValues": [],   //input
//                     "derivativeType": "None",   //select from ENUM
//                     "derivativeExpression": "None", //input
//                     "isBackendOnly": false, //toggle boolean
//                     "displayStatus": "hide",    //select from ENUM
//                     "sampleData":"password",    //input
//                     "inputValidation":{
//                         "description":"This is a description of my validation", //input
//                         "abortOnFailure":true,  //toggle boolean
//                         "customErrorMessage":"This is a custom error message"   //input
//                     }

//                 }
//             ]
//         },
//         {
//             "entityName":"DependentEntity",     //input
//             "entityDescription":"This is a description of my entity",   //input
//             "implementsRBAC":true,      //toggle boolean
//             "isAuthenticationRequired":true,    //toggle boolean
//             "implementsAudit":true,    //toggle boolean
//             "implementsChangeManagement":true,  //toggle boolean
//             "isReadOnly":false,    //toggle boolean
//             "isIndependentEntity":false,    //toggle boolean
//             "dependsOnEntities":[{ "entityName":"User",     //select from a list of entities
//             "fieldName":"users",    //input
//             "relationType":"OneToMany"  //select from ENUM
//             }
//         ],
//             "preferredDB":"Postgres",   //select from a list of options
//             "modeOfDBInteraction":"StoredProcedures",   //select from ENUM
//             "version":"1.0",    //input
//             "isBackendOnly":false,    //toggle boolean
//             "entityFields":[
//                 {
//                     "fieldName":"name",   //input
//                     "displayName":"Name",   //input
//                     "fieldType":"String",   //select from ENUM
//                     "fieldDescription":"This is a description of my field",   //input
//                     "isMandatory":true,   //toggle boolean
//                     "isUnique":false,   //toggle boolean
//                     "isReadOnly":false,  //toggle boolean
//                     "isSensitive":false,        //toggle boolean
//                     "isEditable":true,  //toggle boolean
//                     "isDerived":false,  //toggle boolean
//                     "isCollection":false,   //toggle boolean
//                     "collectionType":"None",    //select from ENUM
//                     "isEnum": false,    //toggle boolean
//                     "enumValues": [],   //input
//                     "derivativeType":"None",    //select from ENUM
//                     "derivativeExpression":"None",  //input
//                     "isBackendOnly":false,  //toggle boolean
//                     "displayStatus":"Show", //select from ENUM
//                     "sampleData":"Test Name",   //input
//                     "inputValidation":{
//                         "description":"This is a description of my validation", //input
//                         "abortOnFailure":true,  //toggle boolean
//                         "customErrorMessage":"This is a custom error message"   //input

//                     }
//                 },
//                 {
//                     "fieldName":"users",    //input
//                     "displayName":"Users",  //input
//                     "fieldType":"List",
//                     "fieldDescription":"This is a description of my field",  //input
//                     "isMandatory":true, //toggle boolean
//                     "isUnique":false,       //toggle boolean
//                     "isReadOnly":false,     //toggle boolean
//                     "isSensitive":false,    //toggle boolean
//                     "isEditable":true,  //toggle boolean
//                     "isDerived":false,  //toggle boolean
//                     "isCollection":true,    //toggle boolean
//                     "collectionType":"User",    //select from ENUM
//                     "isEnum": false,    //toggle boolean
//                     "enumValues": [],   //input
//                     "derivativeType":"None",    //select from ENUM
//                     "derivativeExpression":"None",  //input
//                     "isBackendOnly":false,  //toggle boolean
//                     "displayStatus":"Detail",   //select from ENUM
//                     "inputValidation":{
//                         "description":"This is a description of my validation", //input
//                         "abortOnFailure":true,  //toggle boolean
//                         "customErrorMessage":"This is a custom error message"   //input

//                     }

//                 }
//             ]
//         }

//     ]

// }

type Project struct {
	ProjectName        string   `json:"projectName"`
	ProjectDescription string   `json:"projectDescription"`
	ProjectType        string   `json:"projectType"`
	IsMultiTenant      bool     `json:"isMultiTenant"`
	IsMultiLingual     bool     `json:"isMultiLingual"`
	Entities           []Entity `json:"entities"`
}

type Entity struct {
	EntityName                 string            `json:"entityName"`
	EntityDescription          string            `json:"entityDescription"`
	ImplementsRBAC             bool              `json:"implementsRBAC"`
	IsAuthenticationRequired   bool              `json:"isAuthenticationRequired"`
	ImplementsAudit            bool              `json:"implementsAudit"`
	ImplementsChangeManagement bool              `json:"implementsChangeManagement"`
	IsReadOnly                 bool              `json:"isReadOnly"`
	IsIndependentEntity        bool              `json:"isIndependentEntity"`
	DependsOnEntities          []DependsOnEntity `json:"dependsOnEntities"`
	Version                    string            `json:"version"`
	IsBackendOnly              bool              `json:"isBackendOnly"`
	PreferredDB                string            `json:"preferredDB"`
	ModeOfDBInteraction        string            `json:"modeOfDBInteraction"`
	EntityFields               []EntityField     `json:"entityFields"`
}

type DependsOnEntity struct {
	EntityName   string `json:"entityName"`
	FieldName    string `json:"fieldName"`
	RelationType string `json:"relationType"`
}

type EntityField struct {
	FieldName            string          `json:"fieldName"`
	DisplayName          string          `json:"displayName"`
	FieldDescription     string          `json:"fieldDescription"`
	FieldType            string          `json:"fieldType"`
	IsMandatory          bool            `json:"isMandatory"`
	IsUnique             bool            `json:"isUnique"`
	IsReadOnly           bool            `json:"isReadOnly"`
	IsSensitive          bool            `json:"isSensitive"`
	IsEditable           bool            `json:"isEditable"`
	IsDerived            bool            `json:"isDerived"`
	IsCollection         bool            `json:"isCollection"`
	CollectionType       string          `json:"collectionType"`
	IsEnum               bool            `json:"isEnum"`
	EnumValues           []string        `json:"enumValues"`
	DerivativeType       string          `json:"derivativeType"`
	DerivativeExpression string          `json:"derivativeExpression"`
	IsBackendOnly        bool            `json:"isBackendOnly"`
	DisplayStatus        string          `json:"displayStatus"`
	SampleData           string          `json:"sampleData"`
	InputValidation      InputValidation `json:"inputValidation"`
}

type InputValidation struct {
	Description        string `json:"description"`
	AbortOnFailure     bool   `json:"abortOnFailure"`
	CustomErrorMessage string `json:"customErrorMessage"`
}
