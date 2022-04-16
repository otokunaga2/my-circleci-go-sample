package main

import "testing"

func TestSaveArticle(t *testing.T) {
	blog := New()
	blog.SaveArticle(Article{"My title", "My post body"})
	if blog.Articles[0].Title != "Mytitle" {
		t.Errorf("Item was not expected title")
	}
}
