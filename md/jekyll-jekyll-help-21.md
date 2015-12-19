# [Theme settings in config or data?](https://github.com/jekyll/jekyll-help/issues/21)

> state: **closed** opened by: **** on: ****

I write a customizable theme for Jekyll. Call it &#x60;theme&#x60;. I see two places where the user can place her settings:

1. &#x60;_config.yml&#x60;
    
    &#x60;&#x60;&#x60;yml
    theme:
      option1: value1
      option2: value2
    &#x60;&#x60;&#x60;
    Then I access the settings via &#x60;site.theme&#x60;.
1. &#x60;_data/theme.yml&#x60;
    
    &#x60;&#x60;&#x60;yml
    option1: value1
    option2: value2
    &#x60;&#x60;&#x60;
    Then I access the settings via &#x60;site.data.theme&#x60;.

Which one is better and why?

### Comments

---
> from: [**mscharley**](https://github.com/jekyll/jekyll-help/issues/21#issuecomment-40973307) on: **Monday, April 21, 2014**

For my plugins I&#x27;ve been favouring option #2, simply because it&#x27;s
segregated. It&#x27;s easier to ask users for their _data/config.yml file and
get just the relevant settings. It&#x27;s also an organisational tool. Just with
jekyll plus a few other settings, my _settings.yml file is getting quite
long as it is.


On 22 April 2014 06:24, Anatol Broder &lt;notifications@github.com&gt; wrote:

&gt; I write a customizable theme for Jekyll. Call it theme. I see two places
&gt; where the user can place her settings:
&gt;
&gt;    1.
&gt;
&gt;    _config.yml
&gt;
&gt;    theme:
&gt;      option1: value1
&gt;      option2: value2
&gt;
&gt;    Then I access the settings via site.theme.
&gt;     2.
&gt;
&gt;    _data/theme.yml
&gt;
&gt;    option1: value1option2: value2
&gt;
&gt;    Then I access the settings via site.data.theme.
&gt;
&gt; Which one is better and why?
&gt;
&gt; —
&gt; Reply to this email directly or view it on GitHub&lt;https://github.com/jekyll/help/issues/21&gt;
&gt; .
&gt;
---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/21#issuecomment-40973326) on: **Monday, April 21, 2014**

I would say &#x60;_config.yml&#x60;, since you&#x27;re configuring the theme. Use data for real datasets like [government organizations that use GitHub](https://github.com/github/government.github.com/blob/gh-pages/_data/organizations.yml).
---
> from: [**albertogg**](https://github.com/jekyll/jekyll-help/issues/21#issuecomment-40973846) on: **Monday, April 21, 2014**

I&#x27;m bit more inclined to the &#x60;_config.yml&#x60;; both ways are good and functionally the same, but those options seems to be more like configuration settings.
---
> from: [**budparr**](https://github.com/jekyll/jekyll-help/issues/21#issuecomment-40976046) on: **Monday, April 21, 2014**

I use _data/settings.yaml for anything I want to expose to non-technical people, and config for everything else. 
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/21#issuecomment-41127433) on: **Tuesday, April 22, 2014**

Thank you for opinions. I will use &#x60;_config.yml&#x60; because:

* the file already exists in every installation
* &#x60;_data&#x60; should contain user content only

