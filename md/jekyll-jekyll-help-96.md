# [Disable the highlighter](https://github.com/jekyll/jekyll-help/issues/96)

> state: **closed** opened by: **** on: ****

To disable the highlighter should I set

&#x60;&#x60;&#x60;yaml
highlighter: null
&#x60;&#x60;&#x60;

in the config?

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/96#issuecomment-48844167) on: **Sunday, July 13, 2014**

The only source for &#x60;null&#x60; I found [is this](https://github.com/jekyll/jekyll/blob/master/lib/jekyll/configuration.rb#L231):

&#x60;&#x60;&#x60;
    if config.has_key? &#x27;pygments&#x27;
        Jekyll.logger.warn &quot;Deprecation:&quot;, &quot;The &#x27;pygments&#x27; configuration option&quot; +
        &quot; has been renamed to &#x27;highlighter&#x27;. Please update your&quot; +
        &quot; config file accordingly. The allowed values are &#x27;rouge&quot;+
        &quot;&#x27;pygments&#x27; or null&quot;
---
> from: [**penibelst**](https://github.com/jekyll/jekyll-help/issues/96#issuecomment-48844718) on: **Sunday, July 13, 2014**

The Yaml [spec](http://yaml.org/type/null.html) says: 

&gt; A null value is used to indicate the lack of a value.

I assume this is what the warning means. Thank you, @kleinfreund.
