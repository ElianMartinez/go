import forge from "node-forge";
import * as xmldom from "xmldom";
//import fs
import fs from "node:fs";

//read xml .p12 file
let xml = fs.readFileSync("xml_sin_firm.xml", "utf8");
//read .p12 file
let p12 = fs.readFileSync("4299954_identity.p12");

let pem = convertToPem(p12, "85cta1070");
let privateKey = forge.pki.privateKeyFromPem(pem.pemKey);

//-----------------------------------------------------------
//Se prepara la encriptacion del contenido del xml
let md = forge.md.sha256.create();
var xmlData2 = new xmldom.DOMParser().parseFromString(xml);
console.log(xml);
//Canolizacion del tag SignedInfo para poder generar correctamente la firma del documento.

var xmlCanolizadoData2 = this.c14nCanonicalization(xmlData2.firstChild, null);
md.update(xmlCanolizadoData2.toString(), "utf8");
//------------------------------------------------------------
/* Se obtiene solo el contenido del en formato encode64
para luego asignar este valor a el tag DigestValue
*/
let messageDigest = forge.util.encode64(md.digest().data);

function convertToPem(p12ArrayBuffer, password) {
  var p12Der = forge.util.createBuffer(p12ArrayBuffer);
  var p12Asn1 = forge.asn1.fromDer(p12Der);
  var p12 = forge.pkcs12.pkcs12FromAsn1(p12Asn1, false, password);
  var pemKey = getKeyFromP12(p12, password);
  var certificate = getCertificateFromP12(p12);
  return {
    pemKey: pemKey,
    pemCertificate: certificate.pemCertificate,
    commonName: certificate.commonName,
  };
}

function getKeyFromP12(p12, password) {
  var keyData = p12.getBags(
    { bagType: forge.pki.oids.pkcs8ShroudedKeyBag },
    password
  );
  var pkcs8Key = keyData[forge.pki.oids.pkcs8ShroudedKeyBag][0];
  if (typeof pkcs8Key === "undefined") {
    pkcs8Key = keyData[forge.pki.oids.keyBag][0];
  }
  if (typeof pkcs8Key === "undefined") {
    throw new Error("Unable to get private key.");
  }
  var pemKey = forge.pki.privateKeyToPem(pkcs8Key.key);
  pemKey = pemKey.replace(/\r\n/g, "");
  return pemKey;
}

function getCertificateFromP12(p12) {
  var certData = p12.getBags({ bagType: forge.pki.oids.certBag });
  var certificate = certData[forge.pki.oids.certBag][0];
  var pemCertificate = forge.pki.certificateToPem(certificate.cert);
  pemCertificate = pemCertificate.replace(/\r\n/g, "");
  var commonName = certificate.cert.subject.attributes[0].value;
  return { pemCertificate: pemCertificate, commonName: commonName };
}

function c14nCanonicalization(node, options) {
  options = options || {};
  var defaultNs = options.defaultNs || "";
  var defaultNsForPrefix = options.defaultNsForPrefix || {};
  var ancestorNamespaces = options.ancestorNamespaces || [];
  var res = c14nCanonicalizationInterno(
    node,
    [],
    defaultNs,
    defaultNsForPrefix,
    ancestorNamespaces
  );
  return res;
}
