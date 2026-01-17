# Libraries Documentation

**Libraries** represent external dependencies or modules that can be imported into the Project. The system features a **Discovery Agent** that can fetch library definition from Git repositories.

## Properties

| Property | Type | Description |
| :--- | :--- | :--- |
| **Name** | `String` | Name of the library. |
| **Type** | `Enum` | Type of library: `INTERNAL`, `EXTERNAL`, `PLUGIN`. |
| **Git Reference** | `String` | URL to the Git repository. |
| **Commit/Tag** | `String` | Specific version to use. |

## Discovery Mechanism
The **Library Discovery Service** allows importing libraries by providing a Repository URL.
-   **Endpoint**: `POST /api/v1/libraries/discover`
-   **Process**:
    1.  Fetches `gen_library.json` (metadata) from the repository root.
    2.  Parses the metadata to create a `Library` entity.
    3.  Supports **Private Repositories** via Personal Access Token (PAT).
