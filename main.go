package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/beevik/etree"
	"golang.org/x/crypto/pkcs12"
)

// cargarClavePrivada carga una clave privada RSA desde un archivo PEM.
func cargarDesdeP12(rutaArchivo, password string) (*rsa.PrivateKey, string, error) {
	data, err := os.ReadFile(rutaArchivo)
	if err != nil {
		return nil, "", err
	}

	// Decodificar el archivo .p12
	priv, cert, err := pkcs12.Decode(data, password)
	if err != nil {
		return nil, "", err
	}

	privateKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return nil, "", errors.New("la clave no es del tipo RSA")
	}

	certificadoBase64 := base64.StdEncoding.EncodeToString(cert.Raw)

	return privateKey, certificadoBase64, nil
}

// firmarXML crea una firma digital para un documento XML.
func firmarXML(xmlData []byte, rutaP12, password string) (string, error) {

	privateKey, certificadoBase64, err := cargarDesdeP12(rutaP12, password)
	if err != nil {
		return "", err
	}

	doc := etree.NewDocument()
	err = doc.ReadFromBytes(xmlData)
	if err != nil {
		return "", err
	}

	// Calcula el DigestValue del contenido XML
	digestHasher := sha256.New()
	_, err = digestHasher.Write(xmlData) // xmlData es el contenido XML a firmar
	if err != nil {
		return "", err
	}

	digestValueBase64 := base64.StdEncoding.EncodeToString(digestHasher.Sum(nil))

	// Calcula la firma (SignatureValue)
	firmaHasher := sha256.New()
	_, err = firmaHasher.Write([]byte(xmlData)) // O el contenido específico que necesita ser firmado
	if err != nil {
		return "", err
	}

	firma, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, firmaHasher.Sum(nil))
	if err != nil {
		return "", err
	}

	firmaBase64 := base64.StdEncoding.EncodeToString(firma)

	// Construir estructura XML de la firma
	firmaXML := fmt.Sprintf(`<Signature xmlns="http://www.w3.org/2000/09/xmldsig#">
<SignedInfo>
<CanonicalizationMethod Algorithm="http://www.w3.org/TR/2001/REC-xml-c14n-20010315" />
<SignatureMethod Algorithm="http://www.w3.org/2001/04/xmldsig-more#rsa-sha256" />
<Reference URI="">
<Transforms>
<Transform Algorithm="http://www.w3.org/2000/09/xmldsig#enveloped-signature" />
</Transforms>
<DigestMethod Algorithm="http://www.w3.org/2001/04/xmlenc#sha256" />
<DigestValue>%s</DigestValue>
</Reference>
</SignedInfo>
<SignatureValue>%s</SignatureValue>
<KeyInfo>
<X509Data>
<X509Certificate>%s</X509Certificate>
</X509Data>
</KeyInfo>
</Signature>`, digestValueBase64, firmaBase64, certificadoBase64)

	firmaElement := etree.NewDocument()
	if err := firmaElement.ReadFromString(firmaXML); err != nil {
		return "", err
	}

	root := doc.Root()
	root.AddChild(firmaElement.Root())

	// Convertir root to string

	xmlFirmado, err := doc.WriteToString()
	if err != nil {
		return "", err
	}

	return xmlFirmado, nil
}

func main() {
	xmlData, err := os.ReadFile("xml.xml")
	if err != nil {
		panic(err)
	}

	rutaP12 := "4299954_identity.p12"
	password := "85cta1070"
	startTime := time.Now()
	xmlFirmado, err := firmarXML(xmlData, rutaP12, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tiempo de ejecución: ", time.Since(startTime).Seconds())

	err = os.WriteFile("test_firmado.xml", []byte(xmlFirmado), 0644)
	if err != nil {
		panic(err)
	}

}
