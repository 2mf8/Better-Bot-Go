package markdown

import (
	"fmt"
	"strings"
)

type MarkDown struct {
	Str string
}

func Builder() *MarkDown {
	return &MarkDown{
		Str: "",
	}
}

/* func (r *MarkDown) MqqApi(content string) *MarkDown {
	str := fmt.Sprintf("\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=false)", content, url.PathEscape(content))
	r.Str += str
	return r
}

func (r *MarkDown) MqqApiAuto(content string) *MarkDown {
	str := fmt.Sprintf("\n[%s](mqqapi://aio/inlinecmd?command=%s&reply=false&enter=true)", content, url.PathEscape(content))
	r.Str += str
	return r
}

func (r *MarkDown) MqqApiAt(nickname string, tinyId uint64) *MarkDown {
	str := fmt.Sprintf("\n[@%s](mqqapi://markdown/mention?at_type=1&at_tinyid=%v)", nickname, tinyId)
	r.Str += str
	return r
}

func (r *MarkDown) MqqApiAtToProfile(nickname string, tinyId uint64) *MarkDown {
	str := fmt.Sprintf("\n[@%s](mqqapi://card/show_pslcard?src_type=internal&version=1&uin=%v&crad_type=friend&source=qrcode)", nickname, tinyId)
	r.Str = str
	return r
} */

// 客户端展示为： /回车指令 用户可点击的标签，群聊和文字子频道不支持该能力。
func (r *MarkDown) CmdEnter(text, show, isTrue string) *MarkDown {
	str := fmt.Sprintf("<qqbot-cmd-enter text=\"%s\" show=\"%s\" reference=\"%s\" />", text, show, isTrue)
	r.Str += str
	return r
}

// 客户端展示为： /参数指令 用户可点击的标签
// isTrue 为 true 或 false
func (r *MarkDown) CmdInput(text, show, isTrue string) *MarkDown {
	str := fmt.Sprintf("<qqbot-cmd-input text=\"%s\" show=\"%s\" reference=\"%s\" />", text, show, isTrue)
	r.Str += str
	return r
}

func (r *MarkDown) Url(name, webUrl string) *MarkDown {
	str := fmt.Sprintf("\n[🔗%s](%s)", name, webUrl)
	r.Str += str
	return r
}

func (r *MarkDown) H1(content string) *MarkDown {
	str := fmt.Sprintf("\n# %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H2(content string) *MarkDown {
	str := fmt.Sprintf("\n## %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) H3(content string) *MarkDown {
	str := fmt.Sprintf("\n### %s", content)
	r.Str += str
	return r
}

func (r *MarkDown) DeleteLine(content string) *MarkDown {
	str := fmt.Sprintf("~~%s~~", content)
	r.Str += str
	return r
}

func (r *MarkDown) Bold(content string) *MarkDown {
	str := fmt.Sprintf("**%s** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) Italic(content string) *MarkDown {
	str := fmt.Sprintf("*%s* ", content)
	r.Str += str
	return r
}

func (r *MarkDown) ItalicBold(content string) *MarkDown {
	str := fmt.Sprintf("***%s*** ", content)
	r.Str += str
	return r
}

func (r *MarkDown) BlockReference(content string) *MarkDown {
	str := fmt.Sprintf("\n> %s\n", content)
	r.Str += str
	return r
}

func (r *MarkDown) Image(text, url string, width, height int) *MarkDown {
	str := fmt.Sprintf("![%s #%vpx #%vpx](%s)", text, width, height, url)
	r.Str += str
	return r
}

func (r *MarkDown) DividerLine() *MarkDown {
	str := "\n ---\n"
	r.Str += str
	r.Str = strings.TrimSpace(r.Str)
	return r
}

func (r *MarkDown) Text(content string) *MarkDown {
	r.Str += content
	return r
}

func (r *MarkDown) NewLine() *MarkDown {
	r.Str += "\n"
	return r
}

func (r *MarkDown) Code(content string) *MarkDown {
	c := strings.ReplaceAll(strings.ReplaceAll(content, "\t", "\t"), "\n", "\n")
	str := fmt.Sprintf("\n```\n%s\n```\n", c)
	r.Str += str
	return r
}
