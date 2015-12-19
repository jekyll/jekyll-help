# [How to generate particular file or folder ?](https://github.com/jekyll/jekyll-help/issues/123)

> state: **closed** opened by: **** on: ****

When I build a page, I need to see results quickly in the browser.
How do I generate only particular page or folder?

Example &#x60;# jekyll server -w  &quot;other/folder/&quot;&#x60;

Thanks!
[My site ](https://github.com/devfuria/www.devfuria.com.br)

### Comments

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/123#issuecomment-52238911) on: **Thursday, August 14, 2014**

In terminal, &#x60;cd&#x60; into your source directory, and use &#x60;jekyll serve -w&#x60;. This will build your site and start a web server on your machine that you can use to test with.
---
> from: [**flaviomicheletti**](https://github.com/jekyll/jekyll-help/issues/123#issuecomment-52239930) on: **Thursday, August 14, 2014**

But this will generate the entire site.
Only a folder, a subfolder or file can not I?
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/123#issuecomment-52240117) on: **Thursday, August 14, 2014**

No, you can only generate the entire site.
---
> from: [**flaviomicheletti**](https://github.com/jekyll/jekyll-help/issues/123#issuecomment-52240267) on: **Thursday, August 14, 2014**

Ok, thanks!
