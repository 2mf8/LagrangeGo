package sign

type Provider func(string, uint32, []byte, map[string]string) (map[string]string, error)
