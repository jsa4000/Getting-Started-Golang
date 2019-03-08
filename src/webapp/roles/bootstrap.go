package roles

import "encoding/json"

const bootstrapData = `[
    {
        "id": "ADMIN",
        "name": "ADMIN"
    },
    {
        "id": "WRITE",
        "name": "WRITE"
    },
    {
        "id": "READ",
        "name": "READ"
    }
]`

// BootstrapData Get the bootstrap data deserialized
func BootstrapData() ([]Role, error) {
	var result []Role
	err := json.Unmarshal([]byte(bootstrapData), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
