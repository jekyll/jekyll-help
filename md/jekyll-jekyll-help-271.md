# [How do you properly enable syntax highlighting for Objective-C code?](https://github.com/jekyll/jekyll-help/issues/271)

> state: **open** opened by: **** on: ****

I&#x27;ve been trying to input Objective-C code into Jekyll but am not able to properly get the syntax highlighting working.

None of the following seem to do anything:

&#x60;&#x60;&#x60;objective-c
&#x60;&#x60;&#x60;objc
&#x60;&#x60;&#x60;objectivec

How do you let Jekyll know about objective-c code?

### Comments

---
> from: [**ronak2121**](https://github.com/jekyll/jekyll-help/issues/271#issuecomment-73934743) on: **Wednesday, February 11, 2015**

To clarify, it looks like it doesn&#x27;t detect comments, 

&#x60;&#x60;&#x60;objc
@interface, @end, @implementation objective-c constructs.
&#x60;&#x60;&#x60;
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/271#issuecomment-73946645) on: **Wednesday, February 11, 2015**

@ronak2121 On a side note, it would be helpful to format &#x60;@interface&#x60;, &#x60;@end&#x60; and &#x60;@implementation&#x60; as code, because you just notified these people about this issue. They probably don&#x27;t have anything to do with it.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/271#issuecomment-73949013) on: **Wednesday, February 11, 2015**

@ronak2121 Did you manage to resolve this issue? It would help future visitors to tell them what you needed to do.

Otherwise, feel free to reopen this one.
---
> from: [**ronak2121**](https://github.com/jekyll/jekyll-help/issues/271#issuecomment-73971202) on: **Wednesday, February 11, 2015**

Well I thought that I could fix this by using:

{% highlight objective-c %}

{% endhighlight %}

but that doesn&#x27;t format anything really.

Do I have to add any css or Javascript to the page to enable this?
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/271#issuecomment-74029515) on: **Wednesday, February 11, 2015**

Do you have a highlighter set in your config? 
