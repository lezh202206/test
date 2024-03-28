package main

import (
	"fmt"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func process(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	defer file.Close()
	f, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	f.Path = "Book1.xlsx"
	f.NewSheet("NewSheet")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Path))
	w.Header().Set("Content-Type", req.Header.Get("Content-Type"))
	if err := f.Write(w); err != nil {
		fmt.Fprint(w, err.Error())
	}
}

func main() {
	http.HandleFunc("/process", process)
	http.ListenAndServe(":8090", nil)
}
