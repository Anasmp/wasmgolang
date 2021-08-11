package main

import (
	"syscall/js"

	"github.com/PaulRosset/go-hacknews"
)

var document = js.Global().Get("document")

func getElementByID(id string) js.Value {
	return document.Call("getElementById", id)
}

func topStories() []hacknews.Post {
	init := hacknews.Initializer{Story: "topstories", NbPosts: 10}
	codes, err := init.GetCodesStory()
	if err != nil {
		panic(err)
	}
	posts, err := init.GetPostStory(codes)
	if err != nil {
		panic(err)
	}
	return posts
}
func renderPost(post hacknews.Post, parent js.Value) {
	println(post.Title)
	li := document.Call("createElement", "li")
	li.Set("style", "padding: 20px;border-bottom: 1px solid #ddd;")
	a := document.Call("createElement", "a")
	text := document.Call("createTextNode", post.Title)
	a.Set("href", post.Url)
	a.Call("appendChild", text)
	li.Call("appendChild", a)
	parent.Call("appendChild", li)
}

func renderPosts(posts []hacknews.Post, parent js.Value) {
	ul := document.Call("createElement", "ul")
	ul.Set("style", "list-style-type:none;")
	parent.Call("appendChild", ul)
	for _, post := range posts {
		renderPost(post, ul)
	}
}

func registerCallbacks(posts []hacknews.Post) {
	stopButton := getElementByID("start")
	buttonText := getElementByID("buttonText")
	buttonText.Set("textContent", "Start ")
	stopButton.Set("disabled", false)
	stopButton.Set("onclick", js.FuncOf(func(js.Value, []js.Value) interface{} {
		println("stopping")
		stopButton.Set("style", "display:none;")
		body := document.Get("body")
		renderPosts(posts, body)
		return nil
	}))

}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Initialized !!")

	// register functions
	posts := topStories()

	registerCallbacks(posts)

	<-c
}
