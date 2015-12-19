# [Depreciation issue with pygments](https://github.com/jekyll/jekyll-help/issues/48)

> state: **closed** opened by: **** on: ****

Hi. I know jekyll for three day now and it seems to work fine so far. I get into it and I love the possibilities it provides to use jekyll in combination with Github pages, R, R markdown, and RStudio (see http://magnusmetz.github.io/). Just one (hopefully last) question: Everytime I call 

~~~
jekyll serve
~~~

the command does its job and parses the site but I also get a message, which confuses me: 

~~~
 Deprecation: The &#x27;pygments&#x27; configuration option has been renamed to &#x27;highlighter&#x27;. Please update your config file accordingly. The allowed values are &#x27;rouge&#x27;, &#x27;pygments&#x27; or null.
~~~

What does that mean? Can you help interpret that and tell me what to do? I just want to understand what is going on. Thanks in advance for further help.

Yours,
Magnus

### Comments

---
> from: [**fabianrbz**](https://github.com/jekyll/jekyll-help/issues/48#issuecomment-43028812) on: **Tuesday, May 13, 2014**

@magnusmetz you have to change your config file,
[this](https://github.com/magnusmetz/magnusmetz.github.io/blob/master/_config.yml#L7) line in particular. Replace &#x60;&#x60;&#x60;pygments&#x60;&#x60;&#x60; with &#x60;&#x60;&#x60;highlighter&#x60;&#x60;&#x60; and set the value you want, check the [docs](http://jekyllrb.com/docs/templates/) for the allowed values
---
> from: [**magnusmetz**](https://github.com/jekyll/jekyll-help/issues/48#issuecomment-43030191) on: **Tuesday, May 13, 2014**

Thanks. I had something like this in mind, but was not sure where to set the commands. Works great now.
