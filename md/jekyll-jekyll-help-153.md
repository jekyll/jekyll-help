# [How to create multi-lang article page? maybe i18n related...](https://github.com/jekyll/jekyll-help/issues/153)

> state: **open** opened by: **** on: ****

I want to create a page which may both includes english and chinese version, which may have in the following layout:

__Title__
__English Content On The Left__  ------  __Chinese Version on Right___
__Date__

And the content should comes from:
_posts/2014-09-18-test.en.md -- for english
_posts/2014-09-18-test.cn.md -- for chinese

And if post only have one version, it should also work:
_posts/2014-09-18-test.md

So, what is the best way to achieve this?

### Comments

---
> from: [**joewils**](https://github.com/jekyll/jekyll-help/issues/153#issuecomment-56040120) on: **Thursday, September 18, 2014**

## _layouts/post.html

&#x60;&#x60;&#x60;liquid
{% capture path %}{{page.path}}{% endcapture %}
{% capture en_path %}{{ path | replace:&#x27;.md&#x27;,&#x27;.en.md&#x27; }}{% endcapture %}
{% capture cn_path %}{{ path | replace:&#x27;.md&#x27;,&#x27;.cn.md&#x27; }}{% endcapture %}
&lt;div class=&quot;en&quot;&gt;
    &lt;h1&gt;{{en_path}}&lt;/h1&gt;
    {% include {{en_path}} %}
&lt;/div&gt;
&lt;div class=&quot;cn&quot;&gt;
    &lt;h1&gt;{{cn_path}}&lt;/h1&gt;
    {% include {{cn_path}} %}
&lt;/div&gt;
&lt;div class=&quot;default&quot;&gt;
{{content}}
&lt;/div&gt;
&#x60;&#x60;&#x60;
## 3 Files for every post

* _posts/2014-09-18-test.md
* _posts/2014-09-18-test.en.md
* _posts/2014-09-18-test.cn.md

---
> from: [**yarcowang**](https://github.com/jekyll/jekyll-help/issues/153#issuecomment-56132985) on: **Thursday, September 18, 2014**

Get it...Thank you...
---
> from: [**yarcowang**](https://github.com/jekyll/jekyll-help/issues/153#issuecomment-56150759) on: **Friday, September 19, 2014**

Hi, sorry, but i got such error when running jekyll.

&#x60;&#x60;&#x60;
 Liquid Exception: Included file &#x27;/Users/yarco/Sites/_personal_/bbish/_includes/_posts/2014-09-18-test.en.md&#x27; not found in _layouts/post.html
error: Included file &#x27;/Users/yarco/Sites/_personal_/bbish/_includes/_posts/2014-09-18-test.en.md&#x27; not found. Use --trace to view backtrace
&#x60;&#x60;&#x60;

Can i add some _if_ statement there to check file exists?
&#x60;&#x60;&#x60;
{% if exists en_path %}
{% include {{en_path}} %}
{% endif %}
&#x60;&#x60;&#x60;
---
> from: [**iamsebastian**](https://github.com/jekyll/jekyll-help/issues/153#issuecomment-69649151) on: **Monday, January 12, 2015**

Maybe this would not work:
http://ecommerce.shopify.com/c/ecommerce-design/t/testing-if-a-file-exists-29624
