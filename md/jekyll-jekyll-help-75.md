# [Unable to pull content from .yml file in _data folder](https://github.com/jekyll/jekyll-help/issues/75)

> state: **closed** opened by: **** on: ****

I&#x27;ve been trying to set up author bios on posts that pull their content from &#x60;/_data/authors.yml&#x60;. No errors are being generated and the divs, spans, etc. that are supposed to hold the information are being generated, but they&#x27;re empty.

To make sure there wasn&#x27;t an error in what I was doing, I created files to match the example given in the documentation:

In &#x60;/_data/members.yml&#x60;:

    - name: Tom Preston-Werner
      github: mojombo

    - name: Parker Moore
      github: parkr

    - name: Liu Fengyun
      github: liufengyun

with the following in an otherwise empty template:

    &lt;ul&gt;
    {% for member in site.data.members %}
      &lt;li&gt;
        &lt;a href=&quot;https://github.com/{{ member.github }}&quot;&gt;
          {{ member.name }}
        &lt;/a&gt;
      &lt;/li&gt;
    {% endfor %}
    &lt;/ul&gt;

This resulted in the following output:

    &lt;ul&gt;
    
    &lt;/ul&gt;

I&#x27;ve been unable to locate any information addressing such an issue. 

### Comments

---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/75#issuecomment-46490728) on: **Wednesday, June 18, 2014**

This may not be exactly helpful, but FYI, the code you have there works for me.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/75#issuecomment-46496784) on: **Wednesday, June 18, 2014**

What version of Jekyll are you using?
---
> from: [**bdharva**](https://github.com/jekyll/jekyll-help/issues/75#issuecomment-46512785) on: **Wednesday, June 18, 2014**

I&#x27;m an idiot -- thought I&#x27;d updated it more recently than I have. Didn&#x27;t even think to check that. Was still running pre-1.3.0. Thanks.
