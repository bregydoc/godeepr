package godeepr


type Query struct {
	Key string
	Attributes []string
}

type Method struct {
	Name string
	Alias string
	Params map[string]interface{}
	Attributes []string
}