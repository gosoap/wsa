package wsa

import "encoding/xml"

const Namespace = "http://www.w3.org/2005/08/addressing"

type To struct {
	XMLName xml.Name `xml:"wsa:To"`
	NS      string   `xml:"xmlns:wsa,attr"`

	Value string `xml:",chardata"`
}

type ReplyTo struct {
	XMLName xml.Name `xml:"wsa:ReplyTo"`
	NS      string   `xml:"xmlns:wsa,attr"`

	Address *Address
}

type Address struct {
	XMLName xml.Name `xml:"wsa:Address"`

	Value string `xml:",chardata"`
}

type Action struct {
	XMLName xml.Name `xml:"wsa:Action"`
	NS      string   `xml:"xmlns:wsa,attr"`

	Value string `xml:",chardata"`
}

type MessageID struct {
	XMLName xml.Name `xml:"wsa:MessageID"`
	NS      string   `xml:"xmlns:wsa,attr"`

	Value string `xml:",chardata"`
}
