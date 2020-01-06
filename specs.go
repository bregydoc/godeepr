package godeepr

const contextualizeOperator = "=>"
const paramsKey = "()"
const contextOperator = "=>"

const voidAttr = ""
const dotAttr = "."
const dashAttr = "-"


func attributeHasValidName(attrName string) bool {
	if attrName == voidAttr {
		return false
	} else if attrName == dotAttr {
		return false
	} else if attrName == dashAttr {
		return false
	}
	return true
}

