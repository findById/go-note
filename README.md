## Using "html/template" to output.
### like: markdown to html.
#### I can do more...

## Example: markdown to html
### Source
```
...header or information...
key:value
\r\n
...content markdown...
```
@see posted/index.md

### Template
@see templates/index.tmpl or [https://golang.org/pkg/html/template/](https://golang.org/pkg/html/template/)

## Post Headers
|header|required|purpose|example|
|:----|:----|:----|:----|
|template|required||template: index|
|title|required|String at top of post|title: This post title|
|authors|required||authors: |
|tags|required||tags: tech|
|create_at|required||create_at: 2006-01-02 15:04:05 -0700|
|private|optional||private: true|
|keywords|optional||keywords: keywords|
|description|optional||description: description|
|thumbnail|optional||thumbnail: https://cdn-images-1.medium.com/max/1600/1*Nst9mLu02tSxQCvkAo3L0A.png|
|authors|optional||authors: |