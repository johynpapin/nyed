package highlight

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/buffer"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/golang"
)

type Highlighter struct {
	parser *sitter.Parser
	input  *sitter.Input
	buffer *buffer.Buffer
}

func NewHighlighter(buffer *buffer.Buffer) *Highlighter {
	highlighter := &Highlighter{
		input: &sitter.Input{
			Encoding: sitter.InputEncodingUTF8,
		},
		buffer: buffer,
	}

	highlighter.input.Read = highlighter.readFunc

	return highlighter
}

func (highlighter *Highlighter) readFunc(offset uint32, position sitter.Point) []byte {
	if int(position.Row) >= highlighter.buffer.LineArray.Len() {
		return nil
	}

	return append(highlighter.buffer.LineArray.Line(int(position.Row)).Slice(int(position.Column)), '\n')
}

func (highlighter *Highlighter) Init() error {
	highlighter.parser = sitter.NewParser()
	highlighter.parser.SetLanguage(golang.GetLanguage())

	return highlighter.update(nil)
}

func (highlighter *Highlighter) OnEdit() error {
	return highlighter.update(nil)
}

func (highlighter *Highlighter) update(tree *sitter.Tree) error {
	tree = highlighter.parser.ParseInput(tree, *highlighter.input)

	query, err := sitter.NewQuery([]byte(GOLANG_HIGHLIGHT_SCM), golang.GetLanguage())
	if err != nil {
		return err
	}

	queryCursor := sitter.NewQueryCursor()

	queryCursor.Exec(query, tree.RootNode())

	highlighter.buffer.LineArray.ClearColors()

	for {
		queryMatch, ok := queryCursor.NextMatch()
		if !ok {
			break
		}

		for _, capture := range queryMatch.Captures {
			highlighter.buffer.LineArray.Line(int(capture.Node.StartPoint().Row)).Colors[int(capture.Node.StartPoint().Column)] = tcell.NewHexColor(0x7f9fbe)
			highlighter.buffer.LineArray.Line(int(capture.Node.EndPoint().Row)).Colors[int(capture.Node.EndPoint().Column)] = tcell.ColorWhite
		}
	}

	return nil
}
