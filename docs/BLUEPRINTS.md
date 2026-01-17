# Blueprints Documentation

**Blueprints** are reusable code templates used by the **Generation Engine** to produce implementation code. They allow the system to generate standardized code patterns for common tasks.

## Properties

| Property | Type | Description |
| :--- | :--- | :--- |
| **Standard Name** | `String` | Unique identifier for the blueprint (e.g., `gin-controller-basic`). |
| **Type** | `Enum` | Type of artifact generated: `CONTROLLER`, `SERVICE`, `REPOSITORY`, `DTO`. |
| **Description** | `String` | Purpose of the blueprint. |
| **Template Path** | `String` | Path to the template source code (or raw content in some modes). |

## Placeholders
Blueprints use `Placeholders` (format: `{{PlaceholderName}}`) to inject dynamic values during generation.

| Property | Type | Description |
| :--- | :--- | :--- |
| **Name** | `String` | The key used in the template (e.g., `EntityName`). |
| **Default Value** | `String` | Fallback value if no input is provided. |
| **Description** | `String` | Used as a **Prompt** for the AI Agent if the value is missing. |

## Generation Engine
The system exposes an API (`POST /api/v1/generation/preview`) that:
1.  Takes a Blueprint ID and a map of Input Values.
2.  Replaces `{{Placeholders}}` with inputs.
3.  **AI Fallback**: If an input is missing, the AI Agent generates a value based on the placeholder's description.

## Reverse Engineering
The system can "Import" existing code to create new Blueprints via `POST /api/v1/importer/parse`.
-   **Input**: Raw code snippet (e.g., Go Struct).
-   **Output**: Blueprint with placeholders automatically extracted from fields.
