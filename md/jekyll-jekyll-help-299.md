# [An image wrapped inside div does not display](https://github.com/jekyll/jekyll-help/pull/299)

> state: **closed** opened by: **** on: ****

An image wrapped inside a div does not show up in the html file created. The element #my_photo is styled as follows: 

#my_photo
{
  float: right;
  padding-left: 50px;
}

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/pull/299#issuecomment-100435673) on: **Friday, May 08, 2015**

This is not a Jekyll issue. When you use Markdown or HTML correctly to show an image and have no CSS applied that will hide the image, the image will show up. Jekyll doesn&#x27;t interfere with that.

Also I don&#x27;t understand what this pull request is doing here.
---
> from: [**jaybe-jekyll**](https://github.com/jekyll/jekyll-help/pull/299#issuecomment-100435818) on: **Friday, May 08, 2015**

Depending on the markdown renderer and options, sometimes markdown is not rendered within html blocks.

Does the markdown display versus the image?

Try separating the html and markdown with a newline/whitespace.

Also try without the html block.
---
> from: [**knsam**](https://github.com/jekyll/jekyll-help/pull/299#issuecomment-100543073) on: **Saturday, May 09, 2015**

Hey, first off, this must have been an issue, I realised it only after I posted here. So, sorry! 

@jaybe-jekyll Yes, the markdown code shows up instead of the image. 
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/pull/299#issuecomment-100550965) on: **Saturday, May 09, 2015**

Then it might be the issue. Consider sometthing like this:

&#x60;&#x60;&#x60;html
&lt;p&gt;Nice, but _sneaky_ paragraph.&lt;/p&gt;
&#x60;&#x60;&#x60;

This will render a paragraph, but the &#x60;sneaky&#x60; won&#x27;t be put inside &#x60;em&#x60; tags, but as is (i.e. &#x60;_sneaky_ &#x60;, not &#x60;&lt;em&gt;sneaky&lt;/em&gt;&#x60;). The problem is you can&#x27;t mix HTML and Markdown in the same line.

This will work:

&#x60;&#x60;&#x60;md
Nice, but  _not so sneaky_ paragraph.
&#x60;&#x60;&#x60;

… which is the same as this:

&#x60;&#x60;&#x60;html
&lt;p&gt;Nice, but  &lt;em&gt;not so sneaky&lt;/em&gt; paragraph.&lt;/p&gt;
&#x60;&#x60;&#x60;

If this doesn&#x27;t help you, it&#x27;d really help to have an example file with Markdown that is not processed correctly.
---
> from: [**knsam**](https://github.com/jekyll/jekyll-help/pull/299#issuecomment-100565152) on: **Saturday, May 09, 2015**

@kleinfreund Thank you very much! That indeed clarifies my problem. And, I shall so close this issue. 

Thank you for putting up with my newbie question. 
