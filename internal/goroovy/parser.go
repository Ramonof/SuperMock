package goroovy

import (
	"errors"
)

var noRightCond = errors.New("no right cond")

type Tokenized struct {
	Line  int
	Col   int
	Token Token
	Lit   string
}

type Parser struct {
	tokens     []*Tokenized
	pos        int
	lastReturn string
	variables  map[string]*Variable
}

type Variable struct {
	v any
	t string
}

func NewParser(tokens []*Tokenized) *Parser {
	return &Parser{tokens: tokens, pos: -1, lastReturn: "", variables: make(map[string]*Variable)}
}

func (p *Parser) AddVariable(name string, v string) {
	p.variables[name] = &Variable{v, "string"}
}

func (p *Parser) nextToken() *Tokenized {
	p.pos++
	if p.pos >= len(p.tokens) {
		return nil
	}
	return p.tokens[p.pos]
}

func (p *Parser) moveBack() {
	if p.pos-1 < -1 {
		panic("moved back too much")
	}
	p.pos--
}

// TODO return Response struct
func (p *Parser) ParseTokens() (string, error) {
	for {
		token := p.nextToken()
		if token == nil {
			return "", nil
		}
		switch token.Token {
		case NEWLINE:
		case IF:
			ok, err := p.handleIf()
			if err != nil {
				return "", err
			}
			if !ok {
				p.skipIf()
			}
		case IDENT:
			//TODO
		case RETURN:
			r, err := p.handleReturn()
			if err != nil {
				return "", err
			}
			return r, nil
		default:
			return "", errors.New("illegal token")
		}
	}
}

func (p *Parser) handleReturn() (string, error) {
	token := p.nextToken()
	if token == nil {
		return "", errors.New("empty return token")
	}
	switch token.Token {
	case IDENT:
		v, _ := p.variables[token.Lit]
		return v.v.(string), nil
	}
	return "", errors.New("illegal return token")
}

func (p *Parser) skipIf() {
	//TODO refactor for curly brackets {}
	token := p.nextToken()
	if token == nil {
		return
	}
}

type Conditional struct {
	conditionals []*Conditional
	isTrue       bool
}

func (c *Conditional) IsTrue() bool {
	for _, cond := range c.conditionals {
		if !cond.IsTrue() {
			return false
		}
	}
	return true
}

func (p *Parser) handleIf() (bool, error) {
	conditional, err := p.fillConditional()
	if err != nil {
		return false, err
	}
	return conditional.isTrue, nil
}

type ConditionVariable struct {
	variable any
	varType  Token
}

type Condition struct {
	left     *ConditionVariable
	right    *ConditionVariable
	operator Token
}

func (c *Condition) addVar(v any, t Token) error {
	if c.left == nil {
		c.left = &ConditionVariable{variable: v, varType: t}
		return nil
	} else if c.right == nil {
		c.right = &ConditionVariable{variable: v, varType: t}
		return nil
	}
	return errors.New("bad if")
}

func (c *Condition) addOperator(i Token) error {
	if c.operator == 0 {
		c.operator = i
		return nil
	}
	return errors.New("bad if")
}

func (p *Parser) fillConditional() (*Conditional, error) {
	res := &Conditional{}

	//Opening bracket "("
	token := p.nextToken()
	if token == nil {
		return nil, errors.New("bad if")
	}
	if token.Token != BRACKETOPEN {
		return nil, errors.New("bad if")
	}

	for {
		token = p.nextToken()
		if token == nil {
			return nil, errors.New("conditional must end with )")
		}
		switch token.Token {
		case NEWLINE:
		case BRACKETOPEN:
			p.moveBack()
			conditional, err := p.fillConditional()
			if err != nil {
				return nil, err
			}
			res.isTrue = conditional.isTrue
		case BRACKETCLOSE:
			return res, nil
		default:
			p.moveBack()
			condition, err := p.fillCondition()
			if err != nil {
				return nil, err
			}
			res.isTrue = condition.checkCondition(p)

			//skip all until condition can change
			if !res.isTrue {
			loop1:
				for {
					token = p.nextToken()
					if token == nil {
						return nil, errors.New("conditional must end with )")
					}
					switch token.Token {
					case BRACKETCLOSE:
						return res, nil
					case OR:
						p.moveBack()
						break loop1
					}
				}
			} else {
			loop2:
				for {
					token = p.nextToken()
					if token == nil {
						return nil, errors.New("conditional must end with )")
					}
					switch token.Token {
					case BRACKETCLOSE:
						return res, nil
					case AND:
						p.moveBack()
						break loop2
					}
				}
			}
		}
	}
}

func (c *Condition) checkCondition(p *Parser) bool {
	if c.left == nil {
		return false
	}
	if c.operator == 0 {
		return c.left.variable != 0
	}
	l, r := fillVars(p, c)
	switch c.operator {
	case EQUALS:
		return l.v == r.v
	}
	return false
}

func fillVars(p *Parser, c *Condition) (*Variable, *Variable) {
	var l, r *Variable
	if c.left.varType == IDENT {
		l = p.variables[c.left.variable.(string)]
	}
	if c.right.varType == IDENT {
		r = p.variables[c.right.variable.(string)]
	}
	if c.right.varType == QUOTE {
		r = &Variable{
			v: c.right.variable,
			t: "string",
		}
	}
	return l, r
}

func (p *Parser) fillCondition() (*Condition, error) {
	condition := &Condition{}
	condition, err := p.condLeftVar(condition)
	if err != nil {
		return nil, err
	}

	condition, err = p.condOp(condition)
	if err != nil {
		if errors.Is(err, noRightCond) {
			return condition, nil
		}
		return nil, err
	}

	condition, err = p.condRightVar(condition)
	if err != nil {
		return nil, err
	}

	return condition, nil
}

func (p *Parser) condRightVar(condition *Condition) (*Condition, error) {
	token := p.nextToken()
	if token == nil {
		return nil, errors.New("conditional must end with )")
	}

	switch token.Token {
	case QUOTE:
		err := condition.addVar(token.Lit, token.Token)
		if err != nil {
			return nil, err
		}
	case IDENT:
		err := condition.addVar(token.Lit, token.Token)
		if err != nil {
			return nil, err
		}
	case INT:
		err := condition.addVar(token.Lit, token.Token)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("illegal token")
	}
	return condition, nil
}

func (p *Parser) condOp(condition *Condition) (*Condition, error) {
	token := p.nextToken()
	if token == nil {
		return nil, errors.New("conditional must end with )")
	}

	switch token.Token {
	case EQUALS, MORE, LESS, MOREOREQUAL, LESSOREQUAL:
		err := condition.addOperator(token.Token)
		if err != nil {
			return nil, err
		}
	default:
		p.moveBack()
		return condition, noRightCond
	}
	return condition, nil
}

func (p *Parser) condLeftVar(condition *Condition) (*Condition, error) {
	token := p.nextToken()
	if token == nil {
		return nil, errors.New("conditional must end with )")
	}

	switch token.Token {
	case QUOTE:
		err := condition.addVar(token.Lit, token.Token)
		if err != nil {
			return nil, err
		}
	case IDENT:
		lit, _ := p.parseIdent()
		err := condition.addVar(lit, token.Token)
		if err != nil {
			return nil, err
		}
	case INT:
		err := condition.addVar(token.Lit, token.Token)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("illegal token")
	}
	return condition, nil
}

func (p *Parser) parseIdent() (string, error) {
	p.moveBack()
	res := ""
	for {
		token := p.nextToken()
		if token == nil {
			p.moveBack()
			return res, nil
		}

		switch token.Token {
		case IDENT:
			res += token.Lit
		case DOT:
			res += token.Lit
		default:
			p.moveBack()
			return res, nil
		}
	}
}
