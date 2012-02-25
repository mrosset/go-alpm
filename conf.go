package alpm

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

// Parsing routines for pacman.conf format.

type PacmanOption uint

const (
	ConfUseSyslog PacmanOption = 1 << iota
	ConfShowSize
	ConfUseDelta
	ConfTotalDownload
	ConfCheckSpace
	ConfVerbosePkgLists
	ConfILoveCandy
)

var optionsMap = map[string]PacmanOption{
	"UseSyslog":       ConfUseSyslog,
	"ShowSize":        ConfShowSize,
	"UseDelta":        ConfUseDelta,
	"TotalDownload":   ConfTotalDownload,
	"CheckSpace":      ConfCheckSpace,
	"VerbosePkgLists": ConfVerbosePkgLists,
	"ILoveCandy":      ConfILoveCandy,
}

type PacmanConfig struct {
	CacheDir     []string
	HoldPkg      []string
	SyncFirst    []string
	IgnoreGroup  []string
	IgnorePkg    []string
	NoExtract    []string
	NoUpgrade    []string
	RootDir      string
	DBPath       string
	GPGDir       string
	LogFile      string
	Architecture string
	XferCommand  string
	CleanMethod  string
	SigLevel     SigLevel
	Options      PacmanOption
	Repos        []RepoConfig
}

type RepoConfig struct {
	Name     string
	SigLevel SigLevel
	Servers  []string
}

const (
	tokenSection = iota
	tokenKey
	tokenComment
)

type iniToken struct {
	Type   uint
	Name   string
	Values []string
}

type confReader struct {
	*bufio.Reader
	Lineno uint
}

func newConfReader(r io.Reader) confReader {
	if buf, ok := r.(*bufio.Reader); ok {
		return confReader{buf, 0}
	}
	buf := bufio.NewReader(r)
	return confReader{buf, 0}
}

func (rdr *confReader) ParseLine() (tok iniToken, err error) {
	line, overflow, err := rdr.ReadLine()
	switch {
	case err != nil:
		return
	case overflow:
		err = fmt.Errorf("line %d too long", rdr.Lineno)
		return
	}
	rdr.Lineno++

	line = bytes.TrimSpace(line)
	if len(line) == 0 {
		tok.Type = tokenComment
		return
	}
	switch line[0] {
	case '#':
		tok.Type = tokenComment
		return
	case '[':
		closing := bytes.IndexByte(line, ']')
		if closing < 0 {
			err = fmt.Errorf("missing ']' is section name at line %d", rdr.Lineno)
			return
		}
		tok.Name = string(line[1:closing])
		if closing+1 < len(line) {
			err = fmt.Errorf("trailing characters %q after section name %s",
				line[closing+1:], tok.Name)
			return
		}
		return
	default:
		tok.Type = tokenKey
		if idx := bytes.IndexByte(line, '='); idx >= 0 {
			optname := bytes.TrimSpace(line[:idx])
			values := bytes.Split(line[idx+1:], []byte{' '})
			tok.Name = string(optname)
			tok.Values = make([]string, 0, len(values))
			for _, word := range values {
				word = bytes.TrimSpace(word)
				if len(word) > 0 {
					tok.Values = append(tok.Values, string(word))
				}
			}
		} else {
			// boolean option
			tok.Name = string(line)
			tok.Values = nil
		}
		return
	}
	panic("impossible")
}

func ParseConfig(r io.Reader) (conf PacmanConfig, err error) {
	rdr := newConfReader(r)
	rdrStack := []confReader{rdr}
	confReflect := reflect.ValueOf(&conf).Elem()
	var currentSection string
	var curRepo *RepoConfig
lineloop:
	for {
		line, err := rdr.ParseLine()
		// fmt.Printf("%+v\n", line)
		switch err {
		case io.EOF:
			// pop reader stack.
			l := len(rdrStack)
			if l == 1 {
				return conf, nil
			} else {
				rdr = rdrStack[l-2]
				rdrStack = rdrStack[:l-1]
			}
		default:
			return conf, err
		case nil:
			// Ok.
		}

		switch line.Type {
		case tokenComment:
		case tokenSection:
			currentSection = line.Name
			if currentSection != "options" {
				conf.Repos = append(conf.Repos, RepoConfig{})
				curRepo = &conf.Repos[len(conf.Repos)-1]
				curRepo.Name = line.Name
			}
		case tokenKey:
			switch line.Name {
			case "SigLevel":
				// TODO: implement SigLevel parsing.
				continue lineloop
			case "Server":
				curRepo.Servers = append(curRepo.Servers, line.Values...)
				continue lineloop
			case "Include":
				f, err := os.Open(line.Values[0])
				if err != nil {
					err = fmt.Errorf("error while processing Include directive at line %d: %s",
						rdr.Lineno, err)
					return conf, err
				}
				rdr = newConfReader(f)
				rdrStack = append(rdrStack, rdr)
				continue lineloop
			}

			if currentSection != "options" {
				err = fmt.Errorf("option %s outside of [options] section, at line %d",
					line.Name, rdr.Lineno)
				return conf, err
			}
			// main options.
			if opt, ok := optionsMap[line.Name]; ok {
				// boolean option.
				conf.Options |= opt
			} else {
				// key-value option.
				fld := confReflect.FieldByName(line.Name)
				if !fld.IsValid() {
					err = fmt.Errorf("unknown option at line %d: %s", rdr.Lineno, line.Name)
				}
				switch field_p := fld.Addr().Interface().(type) {
				case *string:
					// single valued option.
					*field_p = strings.Join(line.Values, " ")
				case *[]string:
					//many valued option.
					*field_p = append(*field_p, line.Values...)
				}
			}
		}
	}
	panic("impossible")
}
