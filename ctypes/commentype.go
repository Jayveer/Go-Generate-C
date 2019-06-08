package ctypes

type CommentType struct {
	Comment string
}

func (c *CommentType) ToC() string {
	return "//" + c.Comment
}
