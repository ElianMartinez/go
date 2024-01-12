package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func mai1n() {
    // URL del servicio web
    url := "https://ecf.dgii.gov.do/testecf/autenticacion/api/Autenticacion/ValidarSemilla"

    // Abrir el archivo XML
    file, err := os.Open("test_firmado.xml")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer file.Close()

    // Preparar un formulario multipart
    body := &bytes.Buffer{}
    writer := multipart.NewWriter(body)
    part, err := writer.CreateFormFile("xml", "response_1705031384269.xml")
    if err != nil {
        fmt.Println("Error al crear form file:", err)
        return
    }
    _, err = io.Copy(part, file)
    if err != nil {
        fmt.Println("Error al copiar datos al form file:", err)
        return
    }
    writer.Close()

    // Crear la solicitud
    req, err := http.NewRequest("POST", url, body)
    if err != nil {
        fmt.Println("Error al crear la solicitud:", err)
        return
    }

    // Agregar encabezados necesarios
    req.Header.Set("Content-Type", writer.FormDataContentType())
    req.Header.Set("accept", "application/json")

    // Enviar la solicitud
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error al enviar la solicitud:", err)
        return
    }
    defer resp.Body.Close()

    // Leer la respuesta
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error al leer la respuesta:", err)
        return
    }

    // Imprimir la respuesta
    fmt.Println("Respuesta del servidor:", string(respBody))
}

