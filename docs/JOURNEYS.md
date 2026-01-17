# Journeys Documentation

**Journeys** represent the functional workflows within the system. They map how data flows through the application, from user interaction to backend processing.

## Structure

A Journey is hierarchical, allowing for infinite nesting of steps to represent varying levels of abstraction (Business Logic $\rightarrow$ implementation details).

| Property | Type | Description |
| :--- | :--- | :--- |
| **Name** | `String` | Name of the journey/flow. |
| **Programming Language** | `String` | Target language for the implementation (e.g., `Golang`). |
| **Entity Journeys** | `List` | Grouping of flows by their related Entity. |

### Journey Steps
Steps are the building blocks of a journey.

| Property | Type | Description |
| :--- | :--- | :--- |
| **Type** | `Enum` | Type of step: `API_CALL`, `DATABASE_OPERATION`, `LOGIC`, `DECISION`. |
| **Description** | `String` | Human-readable explanation of what the step does. |
| **Level** | `Enum` | Complexity level: `HIGH` (Business), `MEDIUM` (Technical), `LOW` (Code). |
| **SubSteps** | `List` | Nested steps for "Drill-down" views. |
| **Condition** | `String` | Logic condition for branching (e.g., `if user.role == 'admin'`). |

## Features
-   **Zoomable Canvas**: Users can double-click high-level steps to see the detailed sub-steps.
-   **Dual Views**: Toggle between "Business View" (High-level) and "Technical View" (All details).
