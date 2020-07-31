# Relevant for Reddit Digest
---

{{range $key, $list := .masterMap}}
## {{$key}}

{{range $comment := $list }}

{{$comment}}
&nbsp;
{{end}}
---
{{end}}