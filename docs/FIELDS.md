# Entity Fields Documentation

This document outlines the structure and properties of an **Entity Field** within the Gen-Concept system. Fields define the individual attributes of an Entity (e.g., firstName, email, price).

## Properties

| Property | Type | Description |
| :--- | :--- | :--- |
| **Field Name** | `String` | The technical name of the field (camelCase, e.g., `firstName`). |
| **Display Name** | `String` | The human-readable label for the field (e.g., `First Name`). |
| **Field Type** | `Enum` | The data type of the field (e.g., `String`, `Integer`, `Boolean`, `Date`). |
| **Field Description** | `String` | A description of the field's purpose. |
| **Is Mandatory** | `Boolean` | If `true`, the field is required. |
| **Is Unique** | `Boolean` | If `true`, the field value must be unique across all records. |
| **Is Read Only** | `Boolean` | If `true`, the field cannot be modified by users. |
| **Is Sensitive** | `Boolean` | If `true`, the field contains sensitive data (e.g., PII, passwords) and should be handled with care. |
| **Is Editable** | `Boolean` | If `true`, the field can be edited after creation. |
| **Is Backend Only** | `Boolean` | If `true`, the field is not exposed to the frontend. |
| **Display Status** | `Enum` | Controls visibility in lists/details (e.g., `Show`, `Hide`, `Detail`). |
| **Sample Data** | `String` | Example data for documentation or seeding purposes. |

## Advanced Configuration

### Data Structure Types
These options are mutually exclusive.

#### 1. Enumerations
| Property | Type | Description |
| :--- | :--- | :--- |
| **Is Enum** | `Boolean` | If `true`, the field can only take one of a predefined set of values. |
| **Enum Values** | `List<String>` | The allowed values for the enum (e.g., `["Active", "Inactive"]`). |

#### 2. Collections
| Property | Type | Description |
| :--- | :--- | :--- |
| **Is Collection** | `Boolean` | If `true`, the field holds a collection of values. |
| **Collection Type** | `Enum` | The type of collection (e.g., `List`, `Set`, `Map`). |
| **Collection Item Type** | `Enum` | The type of items in the collection: `Primitive` or `Entity`. |
| **Collection Entity** | `UUID` | If Item Type is `Entity`, this specifies which Entity is referenced. |

#### 3. Derived Fields
| Property | Type | Description |
| :--- | :--- | :--- |
| **Is Derived** | `Boolean` | If `true`, the field's value is calculated from other fields. |
| **Derivative Type** | `Enum` | The method of derivation (e.g., `Computed`, `Concatenated`). |
| **Derivative Expression** | `String` | The logic or expression used to derive the value (e.g., `firstName + ' ' + lastName`). |

## Validation

| Property | Type | Description |
| :--- | :--- | :--- |
| **Input Validations** | `List` | A list of custom validation rules. |
| **Description** | `String` | Description of the validation rule. |
| **Abort On Failure** | `Boolean` | If `true`, stops processing if validation fails. |
| **Custom Error Message** | `String` | The message shown to the user upon failure. |
