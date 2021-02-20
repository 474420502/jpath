package jpath

import (
	"fmt"
)

// Object 对象
type object struct {
	Start int
	End   int
}

func (o *object) Get(content []rune) []rune {
	return content[o.Start:o.End]
}

func (o *object) GetString(content []rune) string {
	return string(content[o.Start:o.End])
}

// Path 路径
type Path struct {
	Index  object // 整个路径解析的范围
	Target object // 查找的对象

	Type int // 1 - 2. find all or one target.(move next)  -1. back (move prev) 0. root

	Condition func(cur []byte) bool

	Prev *Path
	Next *Path
}

func skipSpace(content []rune, i *int) {
	for *i < len(content) && content[*i] == ' ' {
		*i++
	}
}

func getCondition(content []rune, i *int) (head *object) {

	n := *i

	head = &object{Start: n}
	defer func() {
		head.End = n
		*i = n
	}()

	for ; n < len(content); n++ {
		c := content[n]
		switch c {
		case ']':
			n++
			return
		case '\\':
			n++
			continue
		default:
			continue
		}
	}

	return
}

func getConditions(content []rune, i *int) (condslist [][]*object) { // 每个[]*object条件都是 or 关系
	n := *i

	var conds []*object

	defer func() {
		condslist = append(condslist, conds)
		*i = n
	}()

	skipSpace(content, &n)
	conds = append(conds, getCondition(content, &n))
	skipSpace(content, &n)

	for n < len(content) {
		c := content[n]

		switch c {
		case '|':
			condslist = append(condslist, conds)
			conds = []*object{}
			break
		case '&':
			break
		case '[':
			conds = append(conds, getCondition(content, &n))
			skipSpace(content, &n)
			continue // 避免n++
		case '/':
			return
		}
		n++
	}

	return
}

func getTarget(content []rune, i *int) (head *object) {

	skipSpace(content, i)
	n := *i

	head = &object{Start: n}
	defer func() {
		head.End = n
		*i = n
	}()

	for ; n < len(content); n++ {
		c := content[n]
		switch c {
		case '[':
			return
		case ' ':
			return
		case '/':
			return
		default:
			continue
		}
	}

	return
}

func headHandler(src []rune) (content []rune) {

	defer func() {
		content = append(content, ' ')
	}()

	for i := 0; i < len(src); i++ {
		c := src[i]
		switch c {
		case ' ':
			continue
		case '/':
			content = append(content)
			content = append(content, src[i:]...)
			return
		case '[':
			content = append(content, '/', '.')
			content = append(content, src[i:]...)
			return
		default:
			content = append(content, '/')
			content = append(content, src[i:]...)
			return
		}
	}

	return
}

// Parse 解析操作路径
func Parse(src []rune) *Path {

	content := headHandler(src)

	cur := &Path{Type: 0}
	for i := 0; i < len(content); i++ {
		c0 := content[i]
		switch c0 {
		case '/':
			p := &Path{}
			if content[i+1] == '/' { // 类型判断
				p.Type = 1
			} else {
				p.Type = 2
			}

			skipSpace(content, &i)

			cur.Next = p

		case ' ':
			continue
		default:
			panic(fmt.Errorf("error %b", c0))
		}
	}

	return nil
}
