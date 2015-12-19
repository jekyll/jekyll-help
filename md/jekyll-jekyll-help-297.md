# [Using config defaults for my than _posts](https://github.com/jekyll/jekyll-help/issues/297)

> state: **open** opened by: **** on: ****

For some reason, I can only get my Jekyll defaults to work, when &#x60;path: “”&#x60;; if I do anything else, it breaks cf. hafniatimes/issues#13. And by break, I mean that the defaults just don’t seem to apply. Currently, I have to either define the values explicitly in the front matter or hardcode the permalinks.

The config file is [here][config], but here is the relevant part:

&#x60;&#x60;&#x60;yaml
defaults:
    -
        scope:
            path: “” # breaks if &#x60;path: _posts&#x60;
            type: &quot;posts&quot;
        values: &amp;en_defaults
            layout: &quot;post&quot;
            author: &quot;ndarville&quot;
            lang: &quot;en&quot;
            categories: &quot;articles&quot;
    -
        scope: # busted
            path: &quot;_posts/da&quot;
        values: &amp;da_defaults
            lang: &quot;da&quot;
            categories:
                - &quot;da&quot;
                - &quot;articles&quot;
#   - # Redundant when top scope path is &quot;&quot;
#       scope: # busted
#           path: &quot;_drafts&quot;
#       values: *en_defaults
    -
        scope: # busted
            path: &quot;_drafts/da&quot;
        values: *da_defaults
&#x60;&#x60;&#x60;

The top-most scope works—but breaks if I use &#x60;_posts&#x60; (or anything) instead of &#x60;””&#x60; as the path. Maybe I’m misunderstanding how the defaults work.

The idea is to have two types of posts:

* English posts/drafts where &#x60;lang: en&#x60; and &#x60;categories: articles&#x60;
* Danish posts/drafts where &#x60;lang: da&#x60; and &#x60;categories: [da, articles]&#x60;

Other example configs haven’t solved it for me either, so maybe someone here can figure out what’s up.


[config]: https://github.com/hafniatimes/hafniatimes.github.io/blob/e33b5c492b4e2c16942a98fd214cc63310342282/_config.yml#L29-L54

### Comments

---
> from: [**kleinfreund**](https://github.com/jekyll/jekyll-help/issues/297#issuecomment-99401600) on: **Wednesday, May 06, 2015**

Have a look at my defaults: https://github.com/kleinfreund/kleinfreund.github.io/blob/master/_config.yml

I come back to you, when I&#x27;m at home. 
