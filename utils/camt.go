package utils

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/beevik/etree"
)

type CamtDocument struct {
	camtDoc          *etree.Document
	cstmrCdtTrfInitn *etree.Element
}

func NewCamtDocument() *CamtDocument {
	camtDoc := etree.NewDocument()
	camtDoc.CreateProcInst("xml", `version="1.0" encoding="UTF-8"`)
	document := camtDoc.CreateElement("Document")
	document.CreateAttr("xmlns", "urn:iso:std:iso:20022:tech:xsd:pain.001.001.03")
	document.CreateAttr("xmlns:xsi", "http://www.w3.org/2001/XMLSchema-instance")
	document.CreateAttr("xsi:schemaLocation", "urn:iso:std:iso:20022:tech:xsd:pain.001.001.03 pain.001.001.03_GBIC_2.xsd")

	cstmrCdtTrfInitn := document.CreateElement("CstmrCdtTrfInitn")

	return &CamtDocument{camtDoc, cstmrCdtTrfInitn}
}

func (c *CamtDocument) AddHeaders(totalAmount float64, nOfTxs int) {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	tm := time.Now()

	GrpHdr := c.cstmrCdtTrfInitn.CreateElement("GrpHdr")
	MsgId := GrpHdr.CreateElement("MsgId")
	MsgId.CreateCharData(fmt.Sprintf("0891220180122140%d", rd.Intn(999999)))

	CreDtTm := GrpHdr.CreateElement("CreDtTm")
	CreDtTm.CreateCharData(tm.Format("2006-01-02T03:04:05Z"))

	NbOfTxs := GrpHdr.CreateElement("NbOfTxs")
	NbOfTxs.CreateCharData(fmt.Sprintf("%d", nOfTxs))

	CtrlSum := GrpHdr.CreateElement("CtrlSum")
	CtrlSum.CreateCharData(fmt.Sprintf("%.2f", totalAmount))

	InitgPty := GrpHdr.CreateElement("InitgPty")
	Nm := InitgPty.CreateElement("Nm")
	Nm.CreateCharData(config.GetString("camt.company"))

	PmtInf := c.cstmrCdtTrfInitn.CreateElement("PmtInf")

	PmtInfId := PmtInf.CreateElement("PmtInfId")
	PmtInfId.CreateCharData(config.GetString("camt.CamtPmtInfId"))

	PmtMtd := PmtInf.CreateElement("PmtMtd")
	PmtMtd.CreateCharData(config.GetString("camt.CamtPmtMtd"))

	PmtInf.AddChild(NbOfTxs.Copy())
	PmtInf.AddChild(CtrlSum.Copy())

	PmtTpInf := PmtInf.CreateElement("PmtTpInf")
	SvcLvl := PmtTpInf.CreateElement("SvcLvl")
	Cd := SvcLvl.CreateElement("Cd")
	Cd.CreateCharData(config.GetString("camt.sepa"))

	ReqdExctnDt := PmtInf.CreateElement("ReqdExctnDt")
	ReqdExctnDt.CreateCharData(tm.Format("2006-01-02"))

	Dbtr := PmtInf.CreateElement("Dbtr")
	Dbtr.AddChild(Nm.Copy())

	DbtrAcct := PmtInf.CreateElement("DbtrAcct")
	Id := DbtrAcct.CreateElement("Id")
	IBAN := Id.CreateElement("IBAN")
	IBAN.CreateCharData(config.GetString("camt.iban"))

	DbtrAgt := PmtInf.CreateElement("DbtrAgt")
	FinInstnId := DbtrAgt.CreateElement("FinInstnId")
	BIC := FinInstnId.CreateElement("BIC")
	BIC.CreateCharData(config.GetString("camt.bic"))

	ChrgBr := PmtInf.CreateElement("ChrgBr")
	ChrgBr.CreateCharData(config.GetString("camt.CamtChrgBr"))
}

func (c *CamtDocument) AddTransactionData(transactionData map[string]*Transaction) {
	for _, data := range transactionData {
		CdtTrfTxInf := c.cstmrCdtTrfInitn.FindElement("PmtInf").CreateElement("CdtTrfTxInf")
		PmtId := CdtTrfTxInf.CreateElement("PmtId")
		EndToEndId := PmtId.CreateElement("PmtId")
		EndToEndId.CreateCharData(config.GetString("camt.CamtEndToEnd") + data.endToEnd)

		Amt := CdtTrfTxInf.CreateElement("Amt")
		InstdAmt := Amt.CreateElement("InstdAmt")
		InstdAmt.CreateAttr("Ccy", "EUR")
		InstdAmt.CreateCharData(data.amount)

		CdtrAgt := CdtTrfTxInf.CreateElement("CdtrAgt")
		FinInstnId := CdtrAgt.CreateElement("FinInstnId")
		BIC := FinInstnId.CreateElement("BIC")
		BIC.CreateCharData(data.bic)

		Cdtr := CdtTrfTxInf.CreateElement("Cdtr")
		Nm := Cdtr.CreateElement("Nm")
		Nm.CreateCharData(data.bankName)

		CdtrAcct := CdtTrfTxInf.CreateElement("CdtrAcct")
		Id := CdtrAcct.CreateElement("Id")
		IBAN := Id.CreateElement("IBAN")
		IBAN.CreateCharData(data.iban)

		Purp := CdtTrfTxInf.CreateElement("Purp")
		Cd := Purp.CreateElement("Purp")
		Cd.CreateCharData(config.GetString("camt.CamtCd"))

		RmtInf := CdtTrfTxInf.CreateElement("RmtInf")
		Ustrd := RmtInf.CreateElement("Ustrd")
		Ustrd.CreateCharData(config.GetString("camt.CamtRef"))
	}
}

func (c *CamtDocument) PrintDocument() {
	c.camtDoc.Indent(2)
	c.camtDoc.WriteTo(os.Stdout)
}
