# [Error with SCSS generation in Jekyll 2.0](https://github.com/jekyll/jekyll-help/issues/41)

> state: **closed** opened by: **** on: ****

I have my project laid out like so:

&#x60;&#x60;&#x60;
-Project
  -css
    -import.scss
    -_sass/
      main.scss
&#x60;&#x60;&#x60;

The contents of import.scss are:
&#x60;&#x60;&#x60;
---
---

@import &quot;main.scss&quot;;
&#x60;&#x60;&#x60;

What I expected to happen was for main.scss to be imported into import.scss, then, import.scss would compile to import.css within the genrated _site/ directory.

Instead, I get the following error
&#x60;&#x60;&#x60;
Conversion error: There was an error converting &#x27;css/import.scss&#x27;.
jekyll 2.0.3 | Error:  Invalid CSS after &quot;-&quot;: expected number or function, was &quot;--&quot;
&#x60;&#x60;&#x60;

I&#x27;m guessing it&#x27;s complaining about the YAML front-matter at the top of import.scss, but jekyll won&#x27;t convert this if I remove the front-matter, will it?

### Comments

---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/41#issuecomment-42753582) on: **Saturday, May 10, 2014**

Only files you want to go directly from SCSS to CSS do you add the YAML front-matter to. If it&#x27;s a partial, don&#x27;t add the front-matter. This may be unclear in the docs.

Can you run jekyll with &#x60;--trace&#x60; and post that output here?
---
> from: [**JHSheridan**](https://github.com/jekyll/jekyll-help/issues/41#issuecomment-42756526) on: **Saturday, May 10, 2014**

Looks as if it was just a misunderstanding on my part. 

My question was answered here http://stackoverflow.com/questions/23585587/error-with-scss-generation-in-jekyll-2-0/

Turns out _only_ the files that you want to be generated into the _site directory as .css files should have the YAML front matter. I just removed the YAML front matter from the files within the _sass directory, and everything runs fine now.

Maybe this could be clearer in the docs. That&#x27;s something I could contribute myself, I suppose.
---
> from: [**kostrse**](https://github.com/jekyll/jekyll-help/issues/41#issuecomment-48800777) on: **Friday, July 11, 2014**

Yes, this could be added to the documentation. I just bumped to the same problem.
---
> from: [**parkr**](https://github.com/jekyll/jekyll-help/issues/41#issuecomment-48802726) on: **Friday, July 11, 2014**

@kostrse Would you mind submitting a pull request?
