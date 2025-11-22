# Entities Documentation

This document outlines the structure and properties of an **Entity** within the Gen-Concept system. An Entity represents a core business object or data model (e.g., User, Order, Product).

## Properties

| Property | Type | Description |
| :--- | :--- | :--- |
| **Entity Name** | `String` | The unique name of the entity (e.g., `User`). |
| **Entity Description** | `String` | A description of what the entity represents. |
| **Implements RBAC** | `Boolean` | If `true`, the entity includes Role-Based Access Control mechanisms. |
| **Is Authentication Required** | `Boolean` | If `true`, interactions with this entity require user authentication. |
| **Implements Audit** | `Boolean` | If `true`, changes to this entity are audited and logged. |
| **Implements Change Management** | `Boolean` | If `true`, the entity supports change management workflows (e.g., approvals for updates). |
| **Is Read Only** | `Boolean` | If `true`, the entity cannot be modified after creation. |
| **Is Independent Entity** | `Boolean` | If `true`, the entity exists independently. If `false`, it may depend on other entities (e.g., a weak entity). |
| **Version** | `String` | The version number of the entity definition (e.g., `1.0`). |
| **Is Backend Only** | `Boolean` | If `true`, this entity is not exposed directly to the frontend API. |
| **Preferred DB** | `Enum` | The preferred database technology for storing this entity (e.g., `Postgres`, `MongoDB`). |
| **Mode Of DB Interaction** | `Enum` | How the application interacts with the database for this entity (e.g., `ORM`, `StoredProcedures`). |

## Relationships

- **Depends On Entities**: A list of other entities this entity depends on (e.g., foreign key relationships).
- **Entity Fields**: An entity is composed of multiple [Fields](./FIELDS.md) that define its attributes.
