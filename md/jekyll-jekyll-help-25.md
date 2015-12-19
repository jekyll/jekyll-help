# [Better Objective-C syntax highlighting?](https://github.com/jekyll/jekyll-help/issues/25)

> state: **closed** opened by: **** on: ****

I know Jekyll supports Objective-C syntax highlighting now, but the Pygments implementation very bland. Is there a better way to do it? I&#x27;ve searched for hours, and short of writing my own (which I have absolutely no idea how to do), I&#x27;m beginning to believe this is the only option. Beggars can&#x27;t be choosers, so if there&#x27;s nothing else I&#x27;ll be fine, but it never hurts to check.

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/25#issuecomment-41587514) on: **Monday, April 28, 2014**

Are you looking for a CSS document which will make your Obj-C look better? Or for a better processor?
---
> from: [**forgot**](https://github.com/jekyll/jekyll-help/issues/25#issuecomment-41587694) on: **Monday, April 28, 2014**

A CSS document. Sorry for the lack of clarification...

(The &quot;writing my own&quot; comment was in reference to a Pygments lexer. I have no problem with Pygments.)
---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/25#issuecomment-41606855) on: **Monday, April 28, 2014**

What are you using already? I know I changed mine up a little in the name of getting it to fit in with Solarized, but I can&#x27;t remember how much. Either way, my copy of the syntax stylesheet is pretty well commented as a starting point:

https://bitbucket.org/matthew-scharley/matt.scharley.me/src/d5531861bf9054537dd2fecacd873227ec1142c9/src/_sass/?at=master

You&#x27;re looking at &#x60;_solarized.scss&#x60; for the colour definitions and &#x60;_syntax.scss&#x60; for the actual highlighting code.
