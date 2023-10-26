package view

import m "blogPost/model"

type BlogView struct {
	Blog *m.Blog
}

func (v *BlogView) Render() string {
	return ""
}
