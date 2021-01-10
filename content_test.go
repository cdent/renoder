package renoder

import (
    "testing"
)

var expectedOutput = `# Hello
This is a paragraph
# Next Section
`

func TestContentOutput(t *testing.T) {
    c1 := NewContent()
    c1.Text = "# Hello"
    c2 := NewContent()
    c2.Text = "This is a paragraph"
    c3 := NewContent()
    c3.Text = "# Next Section"

    c1.InsertSibling(c3)
    c1.InsertChild(c2)

    output := c1.Output()

    if output != expectedOutput {
        t.Errorf("Produced unexpected output. Got: %v\n\nWanted: %v\n\n", output, expectedOutput)
    }
}
