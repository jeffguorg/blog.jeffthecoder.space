package page

type postMeta struct {
	Title string
	Brief string
	Hash  string
}

type post struct {
	postMeta
	Content string
}

// data types used in template layout
type helpPageData struct {
	ContentTemplate string
}

func newHelpPageData() helpPageData {
	return helpPageData{
		ContentTemplate: "help",
	}
}

type postListPageData struct {
	ContentTemplate string
	Posts           []postMeta
}

func newPostListPageData() postListPageData {
	return postListPageData{
		ContentTemplate: "post-list",
		Posts:           nil,
	}
}
