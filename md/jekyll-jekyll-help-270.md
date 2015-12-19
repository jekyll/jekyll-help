# [defaults permalink in config aren&#x27;t working](https://github.com/jekyll/jekyll-help/issues/270)

> state: **closed** opened by: **** on: ****

hello,

i tried to use the following defaults in my _config.yml

    defaults:
      -
        scope:
          path: &quot;&quot;
          type: &quot;posts&quot;
        values:
          permalink: &quot;/blog/:year/:month/:day/:title&quot;

but my links to the post become http://localhost:4000/blog/:year/:month/:month/:day/:title. i&#x27;ve tried without the quotes as well. without the defaults, the custom permalink works but within defaults, that&#x27;s the result i get. am i missing something? isn&#x27;t the above suppose to be &quot;all posts shall use the following permalink&quot;? but it isn&#x27;t working. i&#x27;m running jekyll 2.4.0.

thanks.

### Comments

---
> from: [**w1n78**](https://github.com/jekyll/jekyll-help/issues/270#issuecomment-73822746) on: **Tuesday, February 10, 2015**

sorry just realized, 2.4.0 isn&#x27;t the latest version. i upgraded to 2.5.3 and it fixed it. thanks for an awesome tool.
