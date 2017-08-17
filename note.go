package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

func init() {
	os.Mkdir("templates", os.ModePerm)
}

const (
	TEMPLATE_SUFFIX    = ".html"
	IGNORE_APPEND_LIST = "index.html"

	KEY_HEADER_TEMPLATE    = "template"
	KEY_HEADER_TITLE       = "title"
	KEY_HEADER_AUTHORS     = "authors"
	KEY_HEADER_TAGS        = "tags"
	KEY_HEADER_CREATE_AT   = "create_at"
	KEY_HEADER_PRIVATE     = "private"
	KEY_HEADER_THUMBNAIL   = "thumbnail"
	KEY_HEADER_KEYWORDS    = "keywords"
	KEY_HEADER_DESCRIPTION = "description"
)

var (
	port   = flag.String("p", "9090", "accept port.")
	source = flag.String("source", "posted", "source dir.")
	target = flag.String("target", "html", "target dir.")
	debug  = flag.Bool("d", false, "debug model.")
)

func ParseFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buffer, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ParseMessage(string(buffer)), nil
}

func InitHandler(path string) error {
	docs := []Doc{}

	const prefix = "2006-01-02"
	const suffix = ".md"
	fn := func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(path)) != suffix || info == nil || info.IsDir() {
			return nil
		}
		model, err := ParseFile(path)
		if err != nil {
			return err
		}
		if model[KEY_HEADER_PRIVATE] == "true" {
			// ignored
			return nil
		}
		if model[KEY_HEADER_TITLE] == "" {
			model[KEY_HEADER_TITLE] = "Insert title here"
		}

		dir := path[len(*source):]
		dir = dir[0:strings.LastIndex(dir, string(os.PathSeparator))]
		dir = strings.TrimSpace(dir)

		name := info.Name()
		name = strings.Replace(name, " ", "-", -1)
		filename := name[0: len(name)-len(suffix)]
		if len(filename) > (len(prefix) + 1) {
			// yyyy-MM-dd-*.md ==> yyyy-MM-dd/*.html
			_, err := time.Parse(prefix, filename[0:len(prefix)])
			if err == nil {
				dir = dir + string(os.PathSeparator) + filename[0:len(prefix)]
				filename = filename[len(prefix)+1:]
			}
		}
		os.MkdirAll(*target+string(os.PathSeparator)+dir, os.ModePerm)
		filename = dir + string(os.PathSeparator) + filename

		// for list start
		doc := Doc{}
		doc.Title = model[KEY_HEADER_TITLE]
		doc.Author = model[KEY_HEADER_AUTHORS]
		doc.Tag = model[KEY_HEADER_TAGS]
		doc.Desc = model[KEY_HEADER_DESCRIPTION]
		doc.Date = model["date"]
		if model[KEY_HEADER_CREATE_AT] != "" {
			doc.Date = model[KEY_HEADER_CREATE_AT]
		}
		rp := strings.Replace(filename+".html", string(os.PathSeparator), "/", -1)
		if strings.HasPrefix(rp, "/") {
			rp = rp[1:]
		}
		doc.Permalink = rp
		if rp != IGNORE_APPEND_LIST {
			docs = append(docs, doc)
		}
		// for list end

		filename = *target + string(os.PathSeparator) + filename + ".html"
		// filename = strings.Replace(filename, string(os.PathSeparator) + string(os.PathSeparator), string(os.PathSeparator), -1)
		log.Println("source: " + path)
		log.Println("target: " + filename)

		tpl := model[KEY_HEADER_TEMPLATE]
		if tpl == "" {
			tpl = model["layout"]
		}
		if tpl == "" {
			tpl = "default"
		}
		log.Println("template: " + tpl)

		t, err := template.ParseFiles("templates/" + tpl + TEMPLATE_SUFFIX)
		if err != nil {
			return err
		}
		wr := bytes.NewBufferString("")
		err = t.Execute(wr, model)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, wr.Bytes(), os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}

	err := filepath.Walk(path, fn)
	if err != nil {
		return err
	}

	// for list start
	sort.Sort(sort.Reverse(DocSlice(docs)))

	model := make(map[string]interface{})
	model[KEY_HEADER_TITLE] = "Articles"
	model["items"] = docs
	t, err := template.ParseFiles("templates/list" + TEMPLATE_SUFFIX)
	if err != nil {
		return err
	}
	wr := bytes.NewBufferString("")
	err = t.Execute(wr, model)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(*target+string(os.PathSeparator)+"list.html", wr.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}
	// for list end

	return nil
}

var lastParseTime = time.Now().Unix()

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" + r.RemoteAddr + "][" +
		r.UserAgent() + "][" + r.Host + r.RequestURI + "]")

	if !*debug {
		parseTime := time.Now().Unix()
		r.ParseForm()
		if (parseTime - 5*60) <= lastParseTime {
			w.Write([]byte("Refuse to handle."))
			lastParseTime = parseTime
			return
		}
		lastParseTime = parseTime
	}

	go func() {
		err := InitHandler(*source)
		if err != nil {
			log.Println("failed. " + err.Error())
		} else {
			log.Println("success.")
		}
	}()
	w.Write([]byte("handled."))
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "]["+
		r.RemoteAddr+ "]["+ r.UserAgent()+ "]["+ r.Host+ r.RequestURI+ "]")

	filename := "index.html"

	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	path := uri.Path
	switch path {
	case "/", "/index.html":
		break
	default:
		filename = strings.Replace(path, "../", "/", -1)
		filename = strings.Replace(filename, "/", string(os.PathSeparator), -1)
		break
	}
	log.Println("filename: " + filename)

	f, err := os.Open(*target + string(os.PathSeparator) + filename)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	info, err := f.Stat()
	if err != nil || info.IsDir() {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("content-type", "text/html charset=utf-8")
	defer f.Close()
	io.Copy(w, f)
}

func PreviewHandler(w http.ResponseWriter, r *http.Request) {
	uri, err := url.ParseRequestURI(r.RequestURI)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	path := uri.Path
	preLen := len("/preview/")
	if len(path) <= preLen {
		http.NotFound(w, r)
		return
	}
	path = path[preLen:]
	filename := strings.Replace(path, "../", "/", -1)
	filename = strings.Replace(filename, "/", string(os.PathSeparator), -1)
	log.Println(filename)

	model, err := ParseFile(*source + string(os.PathSeparator) + filename)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tpl := model[KEY_HEADER_TEMPLATE]
	if tpl == "" {
		tpl = "default"
	}
	log.Println("template: " + tpl)

	t, err := template.ParseFiles("templates/" + tpl + TEMPLATE_SUFFIX)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, model)
}

func main() {
	flag.Parse()
	if *port == "" || *source == "" || *target == "" {
		flag.PrintDefaults()
		return
	}
	os.MkdirAll(*source, os.ModePerm)
	os.MkdirAll(*target, os.ModePerm)

	router := []Router{
		Router{"GET", "/preview/", PreviewHandler},
		Router{"GET", "/install/", InstallHandler},
		Router{"GET", "/", ViewHandler},
	}

	mux := NewMux(router)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Println("port: " + *port, "template: templates, source: " + *source + ", target: " + *target)
	go func() {
		http.ListenAndServe(":" + *port, mux)
	}()
	select {}
}
