package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"
)

func uploadModel(fileName string) error{

	fmt.Println("localPath=====>:",fileName)

	//upload file by Http
	hash := HttpUpload(fileName, fileName)
	fmt.Println("hash: " + hash)
	if hash != ""{
		 return nil
	}

	err := errors.New("fail to up load")
	return err
}

func HttpUpload(filePath string, localfileName string) string{
	//open the file
	fh, err := os.Open(filePath)
	if err != nil{
		fmt.Println(err);
		return ""
	} else{
		fmt.Println("open file success")
	}
	defer fh.Close()

	//creat file buffer
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", localfileName)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		fmt.Println("error copy file")
		return ""
	}
	//get file type
	contentType := bodyWriter.FormDataContentType()
	err_bodeWriter := bodyWriter.Close()
	if err_bodeWriter != nil {
		fmt.Println("err_bodeWriter:",err_bodeWriter)
	}

	targetUrl := os.Getenv("IPFS_HOST")

	fmt.Println("targetUrl:",targetUrl)
	//send http request
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		fmt.Println("http error:", err)
		return ""
	}

	defer resp.Body.Close()

	//get response's content
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("resp_body, err:",err)
		return ""
	}

	if resp.Status != "200 OK" {
		fmt.Println("resp.Status:",resp.Status)
		return ""
	}
	return string(resp_body)
}

func compress(model string,RSA_key string,HE_key string,ETH_key string) error{
	var files = [...]string{model,RSA_key,HE_key,ETH_key}

	//creat a buffer to save file's content
	buf := new(bytes.Buffer)

	// creat a new writer writing a zip file to zip_file
	zip_file := zip.NewWriter(buf)

	for _, file := range files {

		// get file name with suffix
		filenameWithSuffix := path.Base(file)

		//get file's body
		file_body, file_err := ioutil.ReadFile(file)
		if file_err != nil {
			return file_err
		}

		//add file to zip
		f, err := zip_file.Create(filenameWithSuffix)
		if err != nil {
			return err
		}

		//write the file‘s body to file in zip
		_, err = f.Write([]byte(file_body))
		if err != nil {
			return err
		}
	}

	//close file
	err := zip_file.Close()
	if err != nil {
		return err
	}

	//write the zip buffer to file
	f, err := os.OpenFile(get_file_name(model) + ".zip", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	_,errBuf := buf.WriteTo(f)
	if errBuf != nil {
		fmt.Println(errBuf)
	}
	return errBuf
}

func get_file_name(Path string)string{

	var filenameWithSuffix string
	filenameWithSuffix = path.Base(Path)

	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)
	fmt.Println("filenameOnly =", filenameOnly)

	return filenameOnly
}

func delete_file(Path string){
	file := Path                   //源文件路径
	err := os.Remove(file)               //删除文件test.txt
	if err != nil {
		fmt.Println("delete error:",err)
	} else{
		fmt.Println("successfully deleted")
	}
}