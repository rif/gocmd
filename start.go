package main

import (
	gs "github.com/ungerik/go-start"
)

var (
	Homepage = &view.Page{
		OnPreRender: func(page *Page, context *Context) (err error) {
			context.Data = &PerPageData{} // Set global page data at request context
		},
		WriteTitle: func(context *Context, writer io.Writer) (err error) {
			writer.Write([]byte(context.Data.(*PerPageData).DynamicTitle))
			return nil
		},
		CSS:         HomepageCSS,
		WriteHeader: RSS("go-start.org RSS Feed", &RssFeed),
		WriteScripts: PageWriters(
			Config.Page.DefaultWriteScripts,
			JQuery, // jQuery/UI is built-in
			JQueryUI,
			JQueryUIAutocompleteFromURL(".select-username", IndirectURL(&API_Usernames), 2),
			GoogleAnalytics(GoogleAnalyticsID), // Google Analytics is built-in
		),
		Content: Views{},
	}

	Admin_Auth = NewBasicAuth("go-start.org", "admin", "password123")
)

func Paths() *ViewPath {
	return &ViewPath{View: Homepage, Sub: []ViewPath{ // /
		{Name: "style.css", View: HomepageCSS}, // /style.css
		{Name: "feed", View: RssFeed},          // /feed/
		{Name: "admin", View: Admin, Auth: Admin_Auth, Sub: []ViewPath{ // /admin/
			{Name: "user", Args: 1, View: Admin_User, Auth: Admin_Auth}, // /admin/user/<USER_ID>/
		}},
		{Name: "api", Sub: []ViewPath{ // 404 because no view defined
			{Name: "users.json", View: API_Usernames}, // /api/users.json
		}},
	},
	}
}

func main() {
	view.Init("go-start.org", CookieSecret, "pkg/myproject", "pkg/gostart") // Set site name, cookie secret and static paths
	view.Config.RedirectSubdomains = []string{"www"}                        // Redirect from www.
	view.Config.Page.DefaultMetaViewport = "width=960px"                    // Page width for mobile devices
	view.RunConfigFile(Paths(), "run.config")
}
