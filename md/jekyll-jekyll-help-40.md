# [{% raw %} {% endraw %} tags not being respected](https://github.com/jekyll/jekyll-help/issues/40)

> state: **closed** opened by: **** on: ****

I can&#x27;t figure this out - for some reason the stuff I put inside of {%raw%}{%endraw%} tags is still being parsed/evaluated by Liquid.

So if I make a .md file with just this for content:
&#x60;&#x60;&#x60;&#x60;
Start
{% raw %}
{{something}}
{% endraw %}
End
&#x60;&#x60;&#x60;&#x60;

My generated HTML just looks like:
Start
End

Same goes for:

&#x60;&#x60;&#x60;&#x60;
Start
{% raw %}
{% sometag %}
{% endraw %}
End
&#x60;&#x60;&#x60;&#x60;

If sometag is a real tag, then Liquid evaluates it and replaces it with whatever that tag gets replaced with, and if it&#x27;s not then jekyll build tells me it doesn&#x27;t know that the tag &#x27;something&#x27;

Been stuck on this for a while and really can&#x27;t figure it out.

**_config.yml:**
encoding: utf-8
markdown: redcarpet 
highlighter: highlighter 
redcarpet:
  extensions: [&#x27;highlight&#x27;,&#x27;strikethrough&#x27;]
...

**relevant gem versions:**
jekyll 2.0.2
liquid 2.5.5
kramdown 1.3.3
redcarpet 3.1.1

### Comments

---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/40#issuecomment-42689886) on: **Friday, May 09, 2014**

Hi @Daniel0524, was this happening in an earlier version of Jekyll? It seems weird because liquid hasn&#x27;t been touched.
---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/40#issuecomment-42798727) on: **Sunday, May 11, 2014**

Works on my end both with _kramdown_ and _redcarpet_ as mardown engine. Only difference in relevant gem versions would be Jekyll v2.0.3. Shouldn&#x27;t be it.
---
> from: [**Daniel0524**](https://github.com/jekyll/jekyll-help/issues/40#issuecomment-42827715) on: **Monday, May 12, 2014**

I&#x27;ll make a new project later today and see if I&#x27;m still running into the same problems.
