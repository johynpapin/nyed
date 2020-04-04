package highlight

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/state"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

type Highlighter struct {
	parser *sitter.Parser
	input  *sitter.Input
	Buffer *state.Buffer
}

func NewHighlighter(buffer *state.Buffer) *Highlighter {
	highlighter := &Highlighter{
		input: &sitter.Input{
			Encoding: sitter.InputEncodingUTF8,
		},
		Buffer: buffer,
	}

	highlighter.input.Read = highlighter.readFunc

	return highlighter
}

func (highlighter *Highlighter) readFunc(offset uint32, position sitter.Point) []byte {
	if int(position.Row) >= highlighter.Buffer.LineArray.Len() {
		return nil
	}

	return append(highlighter.Buffer.LineArray.Line(int(position.Row)).Slice(int(position.Column)), '\n')
}

func (highlighter *Highlighter) Init() {
	highlighter.parser = sitter.NewParser()
	highlighter.parser.SetLanguage(golang.GetLanguage())
}

func (highlighter *Highlighter) ReHighlight() error {
	return highlighter.update()
}

func (highlighter *Highlighter) update() error {
	highlighter.Buffer.Tree = highlighter.parser.ParseInput(highlighter.Buffer.Tree, *highlighter.input)

	query, err := sitter.NewQuery([]byte(GOLANG_HIGHLIGHT_SCM), golang.GetLanguage())
	if err != nil {
		return err
	}

	queryCursor := sitter.NewQueryCursor()

	queryCursor.Exec(query, highlighter.Buffer.Tree.RootNode())

	highlighter.Buffer.LineArray.ClearStyles()

	for {
		queryMatch, ok := queryCursor.NextMatch()
		if !ok {
			break
		}

		for _, capture := range queryMatch.Captures {
			highlighter.Buffer.LineArray.Line(int(capture.Node.StartPoint().Row)).Styles[int(capture.Node.StartPoint().Column)] = tcell.StyleDefault.Foreground(tcell.NewHexColor(0x7f9fbe))
			highlighter.Buffer.LineArray.Line(int(capture.Node.EndPoint().Row)).Styles[int(capture.Node.EndPoint().Column)] = tcell.StyleDefault
		}
	}

	return nil
}
