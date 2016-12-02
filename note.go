package main

import (
	"flag"
	"log"
	"net/http"
	"html/template"
	"time"
	"strings"
	"os"
	"path/filepath"
	"bufio"
	"io/ioutil"
	"bytes"
	"io"
)

func init() {
	os.Mkdir("templates", os.ModeDir)
}

var (
	source = flag.String("source", "posted", "source dir.")
	target = flag.String("target", "html", "target dir.")
)

// Core
func Parse(buffer string) map[string]string {
	const (
		READ_HEADER int = 0
		READ_DATA int = 1
	)
	status := READ_HEADER

	br := bufio.NewReader(strings.NewReader(buffer))

	model := make(map[string]string)
	model["title"] = "Insert title here"

	contentBuffer := bytes.NewBufferString("")
	for {
		temp, _, err := br.ReadLine()
		if err != nil {
			break
		}
		line := string(temp)

		if status == READ_DATA {
			contentBuffer.WriteString(line + "\n")
			continue
		}
		// Empty line. Next read data
		if strings.TrimSpace(line) == "" {
			status = READ_DATA
			continue
		}
		// Handle head
		head := strings.TrimSpace(line)
		index := strings.Index(head, ":")
		if index > 0 && index < len(head) {
			key := strings.TrimSpace(head[0:index])
			value := strings.TrimSpace(head[index + 1:])
			model[key] = value
			log.Println(key + " = " + value)
			continue
		}
	}
	model["content"] = contentBuffer.String()
	return model
}

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
	return Parse(string(buffer)), nil
}

func InitHandler(path string) error {
	const ext = ".md"
	fn := func(path string, info os.FileInfo, err error) error {
		if strings.ToLower(filepath.Ext(path)) != ext || info == nil || info.IsDir() {
			return nil
		}
		model, err := ParseFile(path)
		if err != nil {
			return err
		}

		filename := strings.Replace(model["title"], " ", "-", -1)
		log.Println("filename: " + filename)

		dir := path[len(*source):]
		dir = dir[0:strings.LastIndex(dir, string(os.PathSeparator))]
		dir = strings.TrimSpace(dir)
		log.Println("dir: " + dir)
		if dir != "" {
			os.Mkdir(*target + dir, os.ModeDir)
			filename = dir + string(os.PathSeparator) + filename
		}
		filename = *target + string(os.PathSeparator) + filename + ".html"
		log.Println("filename: " + filename)

		tpl := model["template"]
		if tpl == "" {
			tpl = "root.html"
		}
		log.Println("template: " + tpl)

		t, err := template.ParseFiles("templates/" + tpl)
		if err != nil {
			return err
		}
		wr := bytes.NewBufferString("")
		err = t.Execute(wr, model)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filename, wr.Bytes(), os.ModeAppend)
		if err != nil {
			return err
		}
		return nil
	}

	err := filepath.Walk(path, fn)
	if err != nil {
		return err
	}
	return nil
}

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]")
	err := InitHandler(*source)
	if err != nil {
		w.Write([]byte("failed. " + err.Error()))
	} else {
		w.Write([]byte("success."))
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]")

	filename := "index.html"

	path := r.RequestURI
	switch path {
	case "/", "/index.html":
		break
	default:
		filename = strings.Replace(path, "../", "/", -1)
		filename = strings.Replace(filename, "%20", " ", -1)
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
	path := r.RequestURI
	preLen := len("/preview/");
	if len(path) <= preLen {
		http.NotFound(w, r)
		return
	}
	path = path[preLen:]
	filename := strings.Replace(path, "../", "/", -1)
	filename = strings.Replace(filename, "%20", " ", -1)
	filename = strings.Replace(filename, "/", string(os.PathSeparator), -1)
	log.Println(filename)

	model, err := ParseFile(*source + string(os.PathSeparator) + filename)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	tpl := model["template"]
	if tpl == "" {
		tpl = "root.html"
	}
	log.Println("template: " + tpl)

	t, err := template.ParseFiles("templates/" + tpl)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	t.Execute(w, model)
}

func main() {
	flag.Parse()
	if *source == "" || *target == "" {
		flag.PrintDefaults()
		return
	}
	os.Mkdir(*source, os.ModeDir)
	os.Mkdir(*target, os.ModeDir)

	mux := http.NewServeMux()
	mux.HandleFunc("/preview/", PreviewHandler)
	mux.HandleFunc("/install", InstallHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", ViewHandler)

	log.Println("template: templates, source: " + *source + ", target: " + *target)
	go func() {
		http.ListenAndServe(":9090", mux)
	}()
	select {}
}
