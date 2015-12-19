# [Jekyll doesn&#x27;t respect line breaks](https://github.com/jekyll/jekyll-help/issues/171)

> state: **open** opened by: **** on: ****

The Markdown syntax documentation on Daring Fireball says:

&quot;When you do want to insert a &lt;br /&gt; break tag using Markdown, you end a line with two or more spaces, then type return.&quot;

Doesn&#x27;t work.

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/171#issuecomment-59081161) on: **Tuesday, October 14, 2014**

hi @weilawei, what markdown parser are you using?
---
> from: [**weilawei**](https://github.com/jekyll/jekyll-help/issues/171#issuecomment-59118599) on: **Tuesday, October 14, 2014**

I&#x27;m using kramdown, and I&#x27;ve filed a bug against that. Perhaps this one should be closed in favor of that one?
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/171#issuecomment-59119969) on: **Tuesday, October 14, 2014**

@weilawei maybe you can post your repo? I&#x27;ve never had a problem with this and I&#x27;m wondering if you have your editor set to remove spaces on save, or something like that.

_as an aside, note with Kramdown, you can end a line with two backslashes to create a line break_
