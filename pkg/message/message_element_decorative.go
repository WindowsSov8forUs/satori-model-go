package message

import "golang.org/x/net/html"

type MessageElementStrong struct {
	*childrenMessageElement
}

func (e *MessageElementStrong) Tag() string {
	return "b"
}
func (e *MessageElementStrong) Alias() []string {
	return []string{"strong"}
}

func (e *MessageElementStrong) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageElementEm struct {
	*childrenMessageElement
}

func (e *MessageElementEm) Tag() string {
	return "i"
}
func (e *MessageElementEm) Alias() []string {
	return []string{"em"}
}

func (e *MessageElementEm) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageElementIns struct {
	*childrenMessageElement
}

func (e *MessageElementIns) Tag() string {
	return "s"
}
func (e *MessageElementIns) Alias() []string {
	return []string{"ins"}
}

func (e *MessageElementIns) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageElementSpl struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageElementSpl) Tag() string {
	return "spl"
}

func (e *MessageElementSpl) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageELementCode struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageELementCode) Tag() string {
	return "code"
}

func (e *MessageELementCode) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageELementSup struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageELementSup) Tag() string {
	return "sup"
}

func (e *MessageELementSup) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageELementSub struct {
	*childrenMessageElement
	*noAliasMessageElement
}

func (e *MessageELementSub) Tag() string {
	return "sub"
}

func (e *MessageELementSub) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

func init() {
	regsiterParserMulti([]string{"b", "strong"}, func(n *html.Node) (MessageELement, error) {
		var children []MessageELement
		err := parseHtmlChildrenNode(n, func(e MessageELement) {
			children = append(children, e)
		})
		if err != nil {
			return nil, err
		}
		return &MessageElementStrong{
			&childrenMessageElement{
				Children: children,
			},
		}, nil
	})
	regsiterParserMulti([]string{"i", "em"}, func(n *html.Node) (MessageELement, error) {
		return nil, nil
	})
	regsiterParserMulti([]string{"s", "ins"}, func(n *html.Node) (MessageELement, error) {
		return nil, nil
	})
	regsiterParser("spl", func(n *html.Node) (MessageELement, error) {
		return nil, nil
	})
	regsiterParser("code", func(n *html.Node) (MessageELement, error) {
		return nil, nil
	})
	regsiterParser("sup", func(n *html.Node) (MessageELement, error) {
		return nil, nil
	})

}