# Get Password Policy

## Endpoint
```http
GET /internal/identity/admins/password-policy
```

---

**Description:**
This endpoint retrieves the current password policy configuration for the server. The password policy defines the complexity requirements for user passwords, including requirements for uppercase letters, numbers, symbols, and minimum length.

>**Note:** To access this endpoint, users must have the `ADMIN` role associated to them.

---

## Headers
| Key             | Value                         | Description                              |
| :-------------- | :---------------------------- | :----------------------------------------|
| Content-Type    | application/json              | Indicates that the response body is JSON.|
| Date            | Tue, 03 Dec 2024 19:38:16 GMT | The date and time the request was made.  |
| Authorization   | Bearer <token>                | The bearer token for authentication.     |

---

## Example Request
```http
GET /internal/identity/admins/password-policy HTTP/1.1
Accept: application/json
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...
```

---

## Responses

### Success Response
#### HTTP Status Code: `200 OK`
#### Response Body:
```json
{
  "require_upper": true,
  "require_number": true,
  "require_symbol": false,
  "min_length": 8
}
```

#### Response Fields:
| Field            | Type    | Description                                           |
|:-----------------|:--------|:------------------------------------------------------|
| `require_upper`  | `bool`  | Whether uppercase letters are required in passwords.  |
| `require_number` | `bool`  | Whether numbers are required in passwords.            |
| `require_symbol` | `bool`  | Whether symbols are required in passwords.            |
| `min_length`     | `int`   | Minimum required password length.                     |

---

## Error Responses

### 1. Insufficient Role
#### HTTP Status Code: `403 Forbidden`
#### Response Body:
```json
{
  "error": "insufficient_role",
  "error_description": "the request requires higher privileges than provided by the access token"
}
```

### 2. Unauthorized
#### HTTP Status Code: `401 Unauthorized`
#### Response Body:
```json
{
  "error": "invalid_token",
  "error_description": "the access token is invalid or has expired"
}
```

---

## Notes
- This endpoint returns the currently active password policy configuration.
- The password policy settings can be configured via environment variables or the configuration file.
- These settings apply to all password-related operations (registration, password reset, etc.).
- Default values if not configured:
  - `require_upper`: `false`
  - `require_number`: `false`
  - `require_symbol`: `false`
  - `min_length`: `5`
