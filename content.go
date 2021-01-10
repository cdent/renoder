package renoder

import (
    //"strings"
)

const (
    Full = iota
    Folded
    Children
)

type Content struct {
    Node
    Text string
}

type OutputOptions struct {
    Format string
}

func NewContent() *Content {
    node := NewNode()
    content := Content{}
    content.Node = node
    return &content
}

func (c *Content) Output(options ...OutputOptions) string {
    // TODO: user a string builder?
    var output string

    if c == nil {
        return ""
    }

    output += c.Text + "\n"

    child := c.GetChild()
    if child != nil {
        output += child.(*Content).Output()
    }
    sibling := c.GetSibling()
    if sibling != nil {
        output += sibling.(*Content).Output()
    }
    return output
}
