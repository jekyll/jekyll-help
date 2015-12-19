# [Need CRLF preservation on Windows](https://github.com/jekyll/jekyll-help/issues/196)

> state: **open** opened by: **** on: ****

Hopefully the bad W word doesn&#x27;t leave me without feedback.

Jekyll 2.5.1 and Kramdown currently turn CRLF line endings in LF,
this would be fantastic, except I am using features of IE that require CRLF line endings.
[Specifically &#x27;Mark of the web&#x27;]

You would be correct to say, &#x27;gross&#x27;, but I am required to find a solution anyway]


Is there a suggested, or &#x27;best-way&#x27; to do this?

I need to make sure my Jekyll output on Windows either preserves CRLF [not currently],
or converts line-endings to CRLF.

Thanks in advance





### Comments

---
> from: [**camkego**](https://github.com/jekyll/jekyll-help/issues/196#issuecomment-63420541) on: **Monday, November 17, 2014**

writing a plugin overwriting Jekyll::Convertible#write would be a solution, but I am using Collections, which are not affected by Convertible#write, as a result I wrote issue #197 

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/196#issuecomment-63542017) on: **Tuesday, November 18, 2014**

Use &#x60;Document#write&#x60; for collections, or patch the &#x60;Renderer#transform&#x60; method.
