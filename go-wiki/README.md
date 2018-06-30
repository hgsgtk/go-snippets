# go-wiki

## basic code
Learning at https://golang.org/doc/articles/wiki/

## must feature
- login/logout (authentication)
    - using gorilla/session

## todo to make better
- Store templates in tmpl/ and page data in data/.
- Add a handler to make the web root redirect to /view/FrontPage.
- Spruce up the page templates by making them valid HTML and adding some CSS rules.
- Implement inter-page linking by converting instances of [PageName] to 
<a href="/view/PageName">PageName</a>. (hint: you could use regexp.ReplaceAllFunc to do this)

## todo to solve problem
- don't use regrep
- replace template engine

## reference
- https://gist.github.com/mschoebel/9398202
- https://gist.github.com/alyssaq/75d6678d00572d103106
