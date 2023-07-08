package codegen

/*

//// grammar:

Expression = Expression PLUS Expression    @left(1)   #Plus
           | Expression MINUS Expression   @left(1)   #Minus
					 | Expression TIMES Expression   @left(2)   #Times
					 | Expression DIV Expression     @left(2)   #Div
					 | Expression MOD Expression     @left(3)   #Mod
					 | Expression POWER Expression   @right(4)  #Power .

//// user:

type Parser struct {
	loxParser
}

func (p *Parser) reduceExpressionPlus(e1 ast.Expr, e3 ast.Expr) ast.Expr {
	// ...
}

//// generated:

















*/
