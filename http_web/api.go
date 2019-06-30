package main 

import (
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"encoding/json"
	"os"
	"strings"
	"strconv"
	"encoding/base64"
)

type Page struct {
	Title string
	Body  []byte
}

func encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}

func decode(in string) string {
	bs, _ := base64.StdEncoding.DecodeString(in)
	return string(bs)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HANDLE REQUEST")

	// p, _ := loadPage("TestPage")
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)

	fmt.Fprintf(w, "Hi 3there X")
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {


	//m := Message{"Alice", "Hello", 1294706395881547000}

	// f := Part{"dds", "ds"}
	// responseWithJson(w, f)
}

func responseWithJson(w http.ResponseWriter, data JsonResponse){
	b, err := json.Marshal(data)
 	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
    	return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(b)
}

func fileCreateHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_     = encode(r.FormValue("mimetype"))

	//fmt.Print("Creating " + name + "...")

	os.Mkdir("data", 0777)	

	n := name

	dirPath := "data/" + encode(n)

	var err error

	for i := 0; i < 20; i++ {
		err = os.Mkdir(dirPath, 0777)

		if err == nil { break; }
		
		n = strconv.Itoa(i) + "." + name

		dirPath = "data/" + encode(n)
	}

	fmt.Println(n)

	var response JsonResponse
	if err != nil {
		response = OsProblemsResponse { err.Error() }
	} else {
		response = FileIdResponse { n }
	}
	responseWithJson(w, response)
}

type FormFileProblem     struct { FormFileProblem     string }
type MkDirProblem        struct { MkDirProblem        string }
type FileCreationProblem struct { FileCreationProblem string }
type CopyingProblem      struct { CopyingProblem      string }

func filePartCreateHandler(w http.ResponseWriter, r *http.Request) {
	fid   := encode(r.FormValue("fileId"))
	pid   := encode(r.FormValue("partId"))
	input := r.FormValue("partContent")
	
	// err := os.Mkdir ("data/" + fid, 0777)
	out, err := os.Create("data/" + fid + "/" + pid)
	if err != nil {
		responseWithJson(w, Errorneous { FileCreationProblem { err.Error() } })
		return
	}

	// fmt.Print(" writing...")
	defer out.Close()

	out.Write([]byte(input))

	// fmt.Println(" done.")

	if err != nil {
		responseWithJson(w, Errorneous { CopyingProblem { err.Error() } })
		return
	}
	responseWithJson(w, FileIdResponse { decode(pid) })
}

func fileListHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir("./data")

	var response JsonResponse
	if err != nil {
		response = OsProblemsResponse { err.Error() }
	} else {
		fileResps := make([]FileResponse, 0, 1)

		for _, file := range files {
			id := decode(file.Name())
			fileResps = append(fileResps, FileResponse{id, id})
		}

		response = FileListResponse {fileResps}
	}
	responseWithJson(w, response)
}

func fileGetHandler(w http.ResponseWriter, r *http.Request) {
	fid := encode(r.FormValue("id"))

  	// fmt.Println("GET params:", r.URL);  

	parts, err := ioutil.ReadDir("data/" + fid + "/")

	var response JsonResponse
	if err != nil {
		response = OsProblemsResponse { err.Error() }
	} else {
		partCount := 0
		totalSize := int64(0)
		
		for _, part := range parts {
			if !strings.HasPrefix(part.Name(), ".") {
				fmt.Println(part.Name())
				totalSize += part.Size()
				partCount += 1
			}
		}

		fid = decode(fid)

		response = FileGetResponse { fid, fid, totalSize, partCount }
	}
	responseWithJson(w, response)
}

func filePartGetHandler(w http.ResponseWriter, r *http.Request) {
	fid := encode(r.FormValue("fileId"))
	pid := encode(r.FormValue("partNum"))

  	// fmt.Println("GET params:", r.URL);  

	file, err := os.Open("data/" + fid + "/" + pid)
	defer file.Close()

	if err == nil {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(w, file)
	} else {
		responseWithJson(w, Errorneous { err.Error() })
	}
}

type Errorneous struct {
	Error interface{}
}

type ParamsMissing struct {
	Params []string
}

func shouldHaveParams(handler func(w http.ResponseWriter, r *http.Request), params ...string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var missing []string = make([]string, 0, 10)
		for _, param := range params {
			_, _, err := r.FormFile(param)
			if "" == r.FormValue(param) && err != nil {
				missing = append(missing, param)
			}
		}

		if len(missing) > 0 {
			responseWithJson(w, Errorneous { ParamsMissing { missing } })
		} else {
			handler(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/download/", downloadHandler)
	http.HandleFunc("/file/list/", fileListHandler)
	http.HandleFunc("/file/create/", 
		shouldHaveParams(fileCreateHandler, 
			"name", "mimetype"))
	
	http.HandleFunc("/file/part/create/", 
		shouldHaveParams(filePartCreateHandler,
			"fileId", "partId", "partContent"))
	
	http.HandleFunc("/file/get/", 
		shouldHaveParams(fileGetHandler,
			"id"))

	http.HandleFunc("/file/part/get/", 
		shouldHaveParams(filePartGetHandler,
			"fileId", "partNum"))

	http.ListenAndServe(":8092", nil)
}


type Part struct {
	FileId int
	Hash   string
}

type File struct {
	Name  	 string
	Mimetype string
}

type JsonResponse interface {}

type FileIdResponse struct {
	Id string
}

type FileResponse struct {
	Id string
	Name string
}

type FileGetResponse struct {
	Id 			string
	Name 		string
	Size 		int64
	NumOfParts 	int
}

type FileListResponse struct {
	Files []FileResponse
}

type OsProblemsResponse struct {
	Error string
}

type BoolResponse struct {
	Status bool
}
