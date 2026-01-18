AI Agent Task List for Gen-Concept Platform
===========================================

This document contains a series of detailed, sequential prompts designed to guide an AI coding agent (like Gemini, Cursor, or GitHub Copilot) through the implementation of the Gen-Concept platform refactor.

**Usage:** Execute these tasks in order. Do not proceed to the next phase until the previous one compiles and runs.

Phase 1: The Unified Graph (Solving the "Zoom")
-----------------------------------------------

**Goal:** Replace the recursive JourneyStep architecture with a flat, queryable Graph model (JourneyNode, JourneyEdge) to enable infinite zoom and better performance.

### Task 1.1: Create Graph Models

Context: We need new GORM models to represent the graph.

File: src/domain/model/journey\_graph.go

Prompt:

> Create a new file src/domain/model/journey\_graph.go.
> 
> Define two new structs, JourneyNode and JourneyEdge, which represent a graph structure for Journeys.
> 
> **Requirements:**
> 
> 1.  Both structs must embed BaseModel (from src/domain/model/base.go or base\_model.go - check which one is used by other models like Journey).
>     
> 2.  **JourneyNode Struct:**
>     
>     *   JourneyID: uuid.UUID with a GORM index.
>         
>     *   Type: string (size 50) - e.g., "START", "PROCESS", "API\_CALL".
>         
>     *   Label: string (size 255).
>         
>     *   Level: string (size 20) with a GORM index. This supports zoom levels like "HIGH", "MEDIUM", "LOW".
>         
>     *   ParentNodeID: \*uuid.UUID with a GORM index. This is a pointer to allow null values for root nodes.
>         
>     *   BlueprintID: \*uuid.UUID (pointer). This links the node to a specific generation blueprint.
>         
>     *   Metadata: \[\]byte with GORM type jsonb. This stores frontend layout positions (x, y coordinates).
>         
> 3.  **JourneyEdge Struct:**
>     
>     *   JourneyID: uuid.UUID with a GORM index.
>         
>     *   SourceID: uuid.UUID.
>         
>     *   TargetID: uuid.UUID.
>         
>     *   Label: string (size 50).
>         
> 
> Ensure all necessary packages (like github.com/google/uuid) are imported.

### Task 1.2: Update Database Migrations

Context: The new models need to be created in the PostgreSQL database.

File: src/infra/persistence/migration/1\_Init.go (or wherever AutoMigrate is called).

Prompt:

> Open src/infra/persistence/migration/1\_Init.go (or the file responsible for database migrations).
> 
> Action:
> 
> Add &model.JourneyNode{} and &model.JourneyEdge{} to the list of models being auto-migrated.
> 
> Ensure that this change allows GORM to create the journey\_nodes and journey\_edges tables in the database on startup.

### Task 1.3: Implement Graph Service

Context: We need business logic to fetch nodes based on their "Zoom Level".

File: src/domain/service/journey\_graph\_service.go

Prompt:

> Create a new service file src/domain/service/journey\_graph\_service.go.
> 
> **Requirements:**
> 
> 1.  Define a struct JourneyGraphService that holds a reference to your Repository (or gorm.DB if strictly necessary, but prefer the repository pattern if src/domain/repository allows).
>     
> 2.  Implement a method GetGraph(journeyID uuid.UUID, level string, parentID \*uuid.UUID) (\[\]model.JourneyNode, \[\]model.JourneyEdge, error).
>     
> 3.  **Logic:**
>     
>     *   Query journey\_nodes:
>         
>         *   Filter by journey\_id.
>             
>         *   If level is provided (not empty), filter by level.
>             
>         *   If parentID is provided, filter by parent\_node\_id.
>             
>     *   Query journey\_edges:
>         
>         *   Fetch edges where journey\_id matches.
>             
>         *   **Optimization:** Only return edges where _both_ SourceID and TargetID exist in the list of fetched nodes (to avoid dangling edges in the UI).
>             
> 4.  Return the nodes and edges.
>     

### Task 1.4: Expose API Endpoints

Context: The frontend needs to access this logic.

Files: src/api/handler/journey\_graph.go, src/api/router/journey.go

Prompt:

> Step 1: Create Handler
> 
> Create src/api/handler/journey\_graph.go.
> 
> *   Define a JourneyGraphHandler struct.
>     
> *   Implement GetGraph(c \*gin.Context).
>     
> *   Parse query parameters: level (string) and parent\_id (string/uuid).
>     
> *   Call the JourneyGraphService.
>     
> *   Return JSON: { "nodes": \[...\], "edges": \[...\] }.
>     
> 
> Step 2: Register Route
> 
> Open src/api/router/journey.go.
> 
> *   Add a new route GET /:id/graph to the journey group.
>     
> *   Bind it to JourneyGraphHandler.GetGraph.
>     

Phase 2: The Smart Generation Engine
------------------------------------

**Goal:** Transform the generation engine from a simple string replacer to a context-aware Go template engine capable of handling loops, logic, and imports.

### Task 2.1: Define Generation Context

Context: Define the data structure passed to blueprints.

File: src/domain/service/generation\_context.go

Prompt:

> Create src/domain/service/generation\_context.go.
> 
> **Define these structs:**
> 
> 1.  GenContext:
>     
>     *   ProjectName (string)
>         
>     *   Entity (GenEntity)
>         
>     *   Imports (\[\]string) - A list of Go import paths.
>         
> 2.  GenEntity:
>     
>     *   Name (string)
>         
>     *   VarName (string) - The lowerCamelCase version of the name (e.g., "user").
>         
>     *   Fields (\[\]GenField)
>         
>     *   PrimaryKey (string)
>         
> 3.  GenField:
>     
>     *   Name (string)
>         
>     *   Type (string)
>         
>     *   JSONTag (string)
>         
>     *   ValidateTag (string)
>         

### Task 2.2: Implement Context Builder

Context: Map the database Entity model to the GenContext.

File: src/domain/service/generation\_service.go

Prompt:

> Open src/domain/service/generation\_service.go.
> 
> **Add a new method:** BuildContext(entity model.Entity) GenContext.
> 
> **Logic:**
> 
> 1.  Initialize GenContext with the entity name and a derived VarName (lowercase).
>     
> 2.  Iterate over entity.EntityFields.
>     
> 3.  Map each model.EntityField to GenField.
>     
> 4.  **Smart Imports:**
>     
>     *   If a field type is enum.Date or similar, append "time" to GenContext.Imports.
>         
>     *   If a field has validation rules, generate the binding:"..." or validate:"..." tags in GenField.ValidateTag.
>         
> 5.  Return the populated GenContext.
>     

### Task 2.3: Integrate Template Engine

Context: Switch from strings.Replace to text/template.

File: src/domain/service/generation\_service.go

Prompt:

> Modify the GenerateCode method in src/domain/service/generation\_service.go.
> 
> **Refactoring Steps:**
> 
> 1.  Instead of iterating over placeholders and using strings.ReplaceAll:
>     
> 2.  Parse blueprint.TemplatePath (assuming it contains the template string for now) using Go's text/template package.
>     
> 3.  Call s.BuildContext(entity) to get the data object.
>     
> 4.  Execute the template into a bytes.Buffer.
>     
> 5.  Return the string result.
>     
> 
> **Note:** Ensure the template engine allows access to the fields defined in Task 2.1.

Phase 3: Library Discovery (The "Harvester")
--------------------------------------------

**Goal:** Create an agent that scans Git repositories to find reusable corporate functions.

### Task 3.1: Create Library Definition Model

Context: We need to store metadata about discovered functions.

File: src/domain/model/library.go

Prompt:

> Open src/domain/model/library.go.
> 
> **Add a new struct:** LibraryDefinition.
> 
> *   BaseModel embedding.
>     
> *   PackageName (string).
>     
> *   FunctionName (string).
>     
> *   Signature (string) - The full function signature.
>     
> *   Description (string).
>     
> *   Tags (\[\]string) - Use GORM serializer:json.
>     
> *   RepoURL (string).
>     

### Task 3.2: Implement AST Harvester

Context: Parse Go code to find exported functions.

File: src/domain/service/harvester.go

Prompt:

> Create src/domain/service/harvester.go.
> 
> **Implement ScanRepo(codeContent string) (\[\]model.LibraryDefinition, error):**
> 
> 1.  Use go/token and go/parser to parse the codeContent string.
>     
> 2.  Use go/ast to walk the syntax tree.
>     
> 3.  Look for \*ast.FuncDecl nodes.
>     
> 4.  Filter for functions where Name.IsExported() is true.
>     
> 5.  Extract the function Name and recreate the Signature string from the AST node.
>     
> 6.  Return a slice of LibraryDefinition objects populated with this data.
>     

### Task 3.3: Integrate Harvester into Generation

Context: Use discovered libraries during code generation.

File: src/domain/service/generation\_service.go

Prompt:

> Update the BuildContext method in src/domain/service/generation\_service.go.
> 
> **Logic:**
> 
> 1.  Check if the entity has specific flags (e.g., IsSensitive or ImplementsAudit).
>     
> 2.  If IsSensitive is true:
>     
>     *   Query the LibraryDefinition repository (you may need to inject this dependency).
>         
>     *   Find a library function tagged with "encryption".
>         
>     *   If found, add the library's package to GenContext.Imports.
>         
>     *   Add the function name to a new map in GenContext (e.g., LibraryFunctions\["Encrypt"\]).
>