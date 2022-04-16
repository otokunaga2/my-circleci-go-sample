package main

import "testing"

func TestSaveArticle(t *testing.T) {
	blog := New()
	blog.SaveArticle(Article{"My title", "My post body"})
	testTitle := "My title"
	if blog.Articles[0].Title != testTitle {
		t.Errorf("Item was not expected title expect My title but actual %s", blog.Articles[0].Title)
	}
}

func TestFetchAllArticles(t *testing.T){
	blog := New()

	blog.SaveArticle(Article{"My title", "My post body"})

	articles := blog.FetchAll()

	if len(articles) == 0 {
		t.Errorf("Fetch all number of articles")
	}
}
