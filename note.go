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

const (
	READ_HEADER int = 0
	READ_DATA int = 1
)

var (
	source = flag.String("source", "posted", "source dir.")
	target = flag.String("target", "html", "target dir.")
)

func InitHandler(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if (info == nil) {
			return err;
		}
		if (info.IsDir()) {
			return nil;
		}

		file, err := os.Open(path);
		if err != nil {
			panic(err);
		}

		defer file.Close();
		buffer, err := ioutil.ReadAll(file);
		if err != nil {
			panic(err);
		}
		br := bufio.NewReader(strings.NewReader(string(buffer)));
		task := READ_HEADER;

		model := make(map[string]string);
		model["title"] = "Insert title here";

		contentBuffer := bytes.NewBufferString("");
		for {
			temp, _, err := br.ReadLine()
			if err != nil {
				break;
			}
			line := string(temp);

			if (task == READ_DATA) {
				contentBuffer.WriteString(line + "\n");
				continue;
			}

			// Empty line. Next read data
			if (line == "") {
				task = READ_DATA;
				continue;
			}

			// Handle head
			head := strings.TrimSpace(line)
			index := strings.Index(head, ":");
			if (index > 0 && index < len(head)) {
				key := strings.TrimSpace(head[0:index]);
				value := strings.TrimSpace(head[index + 1:len(head)]);
				model[key] = value;
				log.Println(key + " = " + value);
				continue;
			}
		}

		model["content"] = contentBuffer.String();

		filename := strings.Replace(model["title"], " ", "-", -1) + ".html";
		log.Println("filename: " + filename);

		tpl := model["template"];
		if (tpl == "") {
			tpl = "index.tpl";
		}
		log.Println("template: " + tpl);

		t, err := template.ParseFiles("templates/" + tpl);
		if (err != nil) {
			panic(err);
		}
		wr := bytes.NewBufferString("")
		err = t.Execute(wr, model);
		if (err != nil) {
			panic(err);
		}
		err = ioutil.WriteFile(*target + string(os.PathSeparator) + filename, wr.Bytes(), os.ModeAppend);
		if (err != nil) {
			panic(err);
		}
		return nil
	});
	if err != nil {
		return err;
	}
	return nil;
}

func InstallHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]");
	err := InitHandler(*source);
	if err != nil {
		w.Write([]byte("failed."));
	} else {
		w.Write([]byte("success."));
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]");

	path := r.RequestURI;
	index := strings.LastIndex(path, "/");
	log.Println(index)
	if (index <= 0 || index > len(path)) {
		http.NotFound(w, r);
		return;
	}
	name := path[index:len(path)];
	log.Println(name)

	f, err := os.Open(*target + string(os.PathSeparator) + name);
	if err != nil {
		http.NotFound(w, r);
		return;
	}
	w.Header().Set("content-type", "text/html; charset=utf-8");
	defer f.Close();
	io.Copy(w, f)
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]");
	w.Write([]byte("article"));
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("console", "request", "[" + time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05") + "][" +
		r.RemoteAddr + "][" + r.UserAgent() + "][" + r.Host + r.RequestURI + "]");

	if (r.RequestURI != "/" && strings.ToLower(r.RequestURI) != "/index.html") {
		http.NotFound(w, r);
		return;
	}

	f, err := os.Open(*target + string(os.PathSeparator) + "index.html");
	if err != nil {
		http.NotFound(w, r);
		return;
	}
	w.Header().Set("content-type", "text/html; charset=utf-8");
	defer f.Close();
	io.Copy(w, f)
}

func main() {
	flag.Parse();
	if *source == "" || *target == "" {
		flag.PrintDefaults();
		return;
	}
	os.Mkdir(*source, os.ModeDir)
	os.Mkdir(*target, os.ModeDir)

	mux := http.NewServeMux();
	mux.HandleFunc("/install", InstallHandler);
	mux.HandleFunc("/view/*", ViewHandler);
	mux.HandleFunc("/article", ArticleHandler);
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))));
	mux.HandleFunc("/", IndexHandler);

	log.Println("template: templates, source: " + *source + ", target: " + *target);
	go func() {
		http.ListenAndServe(":9090", mux);
	}()
	select {}
}
