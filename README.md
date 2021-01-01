
Renoder is an experiment following several in the past, including
[PurpleWiki](https://github.com/eekim/purplewiki), to model a body of text as a
tree of uniquely identified nodes that may be transcluded to other trees.

This experiment differs from the others by being written in golang instead of
Python or Perl, to see if the CSP model provided by go has any impact on the
operations that might be possible with the data.

The basic datum is a binary tree with some enforced semantics:

```
          <root>
            /\
           /  \
          /    \
   <sibling>  <child>
```

(or maybe sibling on the right, it depends on how things play out, but the idea
is to be able to rotate the descending external lines to horizontal or vertical
and for that to mean _something_).

# Presentation

Any node can be an entry point. Any node can be retrieved in isolation or as a
tree. It can be rendered to an HTML document or provided a more
pre-presentation format for client-side rendering. A retrieved tree can be ids
or ids + content.
