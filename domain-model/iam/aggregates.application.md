# IAM Domain Model

```go
package product

type Application struct {
	// TODO: Add methods
}
```

```json
{
  "id": "00000000-0000-0000-0000-000000000000",
  "name": "My App",
  "display_name": "MY_APP",
  "description": "App description",
  "productIds": [
    "00000000-0000-0000-0000-000000000000"
  ],
  "features": [{
    "id": "00000000-0000-0000-0000-000000000000",
    "name":"BASIC_AUDIENCES",
    "permissionIds": [
    "00000000-0000-0000-0000-000000000000"
    ]
  }],
  "createdDateTime": "2020-01-01T00:00:00.0000000Z",
  "updatedDateTime": "2020-01-01T00:00:00.0000000Z"
}
```


