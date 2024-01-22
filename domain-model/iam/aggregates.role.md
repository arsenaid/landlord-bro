# IAM Domain Model

```go
package role

type Role struct {
	// TODO: Add methods
}
```

```json
{
  "id": "00000000-0000-0000-0000-000000000000",
  "name": "USER_ROLE",
  "displayName": "My Role",
  "description": "Best role ever",
  "groupIds": [
    "00000000-0000-0000-0000-000000000000"
  ],
  "permissions": [{
      "id": "00000000-0000-0000-0000-000000000000",
      "name": "ACCESS_ORGANIZATION",
      "action": "VIEW",
      "featureId": "00000000-0000-0000-0000-000000000000"
  }],
  "createdDateTime": "2020-01-01T00:00:00.0000000Z",
  "updatedDateTime": "2020-01-01T00:00:00.0000000Z"
}
```


