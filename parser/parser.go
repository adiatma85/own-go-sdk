package parser

import "github.com/adiatma85/own-go-sdk/log"

type Parser interface {
	JSONParser() JSONInterface
}

type parser struct {
	json JSONInterface
}

type Options struct {
	JSONOptions JSONOptions
}

func InitParser(log log.Interface, opt Options) Parser {
	return &parser{
		json: initJSON(opt.JSONOptions, log),
	}
}

func (p *parser) JSONParser() JSONInterface {
	return p.json
}
