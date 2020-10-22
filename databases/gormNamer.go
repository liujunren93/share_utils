package databases

import (
	"crypto/sha1"
	"fmt"
	"github.com/jinzhu/inflection"
	"gorm.io/gorm/schema"
	"strings"
	"sync"
	"unicode/utf8"
)

var (
	smap                      sync.Map
	commonInitialismsReplacer *strings.Replacer
	commonInitialisms         = []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "TTL", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	defaultNamingStrategy = NamingStrategy{
		SingularTable: false,
		TrimStr:       "model",
	}
)
func init() {
	var commonInitialismsForReplacer []string
	for _, initialism := range commonInitialisms {
		commonInitialismsForReplacer = append(commonInitialismsForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
	}
	commonInitialismsReplacer = strings.NewReplacer(commonInitialismsForReplacer...)
}

type NamingStrategy struct {
	TablePrefix   string
	SingularTable bool
	TrimStr       string
}

// TableName convert string to table name
func (ns NamingStrategy) TableName(table string) string {

	if ns.TrimStr != "" {
		index := strings.LastIndex(table,  ns.TrimStr)
		if index>=0 {
			table = table[:index]
		}

	}
	if ns.SingularTable {

		fmt.Println(table)
		return ns.TablePrefix + toDBName(table)

	}
	return ns.TablePrefix + inflection.Plural(toDBName(table))

}

// ColumnName convert string to column name
func (ns NamingStrategy) ColumnName(table, column string) string {
	return toDBName(column)
}

// JoinTableName convert string to join table name
func (ns NamingStrategy) JoinTableName(str string) string {
	if strings.ToLower(str) == str {
		return str
	}

	if ns.SingularTable {
		return ns.TablePrefix + toDBName(str)
	}
	return ns.TablePrefix + inflection.Plural(toDBName(str))
}

// RelationshipFKName generate fk name for relation
func (ns NamingStrategy) RelationshipFKName(rel schema.Relationship) string {
	return fmt.Sprintf("fk_%s_%s", rel.Schema.Table, toDBName(rel.Name))
}

// CheckerName generate checker name
func (ns NamingStrategy) CheckerName(table, column string) string {
	return fmt.Sprintf("chk_%s_%s", table, column)
}

// IndexName generate index name
func (ns NamingStrategy) IndexName(table, column string) string {
	idxName := fmt.Sprintf("idx_%v_%v", table, toDBName(column))

	if utf8.RuneCountInString(idxName) > 64 {
		h := sha1.New()
		h.Write([]byte(idxName))
		bs := h.Sum(nil)

		idxName = fmt.Sprintf("idx%v%v", table, column)[0:56] + string(bs)[:8]
	}
	return idxName
}
func toDBName(name string) string {
	if name == "" {
		return ""
	} else if v, ok := smap.Load(name); ok {
		return fmt.Sprint(v)
	}

	var (
		value                          = commonInitialismsReplacer.Replace(name)
		buf                            strings.Builder
		lastCase, nextCase, nextNumber bool // upper case == true
		curCase                        = value[0] <= 'Z' && value[0] >= 'A'
	)

	for i, v := range value[:len(value)-1] {
		nextCase = value[i+1] <= 'Z' && value[i+1] >= 'A'
		nextNumber = value[i+1] >= '0' && value[i+1] <= '9'

		if curCase {
			if lastCase && (nextCase || nextNumber) {
				buf.WriteRune(v + 32)
			} else {
				if i > 0 && value[i-1] != '_' && value[i+1] != '_' {
					buf.WriteByte('_')
				}
				buf.WriteRune(v + 32)
			}
		} else {
			buf.WriteRune(v)
		}

		lastCase = curCase
		curCase = nextCase
	}

	if curCase {
		if !lastCase && len(value) > 1 {
			buf.WriteByte('_')
		}
		buf.WriteByte(value[len(value)-1] + 32)
	} else {
		buf.WriteByte(value[len(value)-1])
	}

	return buf.String()
}
