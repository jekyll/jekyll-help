# [Github Pages sees required YAML frontmatter in .scss as syntax error](https://github.com/jekyll/jekyll-help/issues/217)

> state: **closed** opened by: **** on: ****

Hi there-

I&#x27;m trying to set up a personal Jekyll site with GH pages at [egardner.github.io](egardner.github.io). [Source code here](https://github.com/egardner/egardner.github.io).

When building the site, I keep getting error messages from Github complaining of syntax errors in my &#x60;main.scss&#x60; file:

&#x60;&#x60;&#x60;
The file &#x60;assets/styles/main.scss&#x60; contains syntax errors. 
For more information, see https://help.github.com/articles/page-build-failed-markdown-errors.
&#x60;&#x60;&#x60;

As far as I can tell, the error is caused by the two lines of YAML front-matter (&#x60;---&#x60;) at the beginning of the file. When I remove those lines and commit the site builds successfully, but all my styles break because the stylesheet SCSS is no longer being converted or processed by Jekyll.

Running &#x60;bundle exec jekyll serve&#x60; (the &#x60;jekyll serve&#x60; command won&#x27;t work by itself for me, for some reason) works fine locally, all the styles and assets show up. Only my &#x60;main.scss&#x60; file contains the frontmatter, all other sass files live in the Sass directory that I&#x27;ve specified in &#x60;config.yml&#x60;.

It seems strange that GH Pages sees a syntax error where I&#x27;m following Jekyll&#x27;s guidelines exactly, as far as I can tell. Sass is supposed to be supported by the latest version of gh-pages.

I imagine many people here are using GH pages for deployment – am I doing something wrong?

Thanks,
Eric


For reference, here&#x27;s my &#x60;main.scss&#x60; file:
&#x60;&#x60;&#x60;
---
---
/* Bourbon, Bitters, &amp; Neat
 * ---------------------------------------------------------------------------*/
@import &quot;lib/normalize&quot;;
@import &quot;lib/bourbon/bourbon&quot;;
@import &quot;base/grid-settings&quot;;
@import &quot;lib/neat/neat&quot;;

/* Fonts
 * ---------------------------------------------------------------------------*/
// Merriweather
@include font-face(&quot;Merriweather&quot;, &#x27;/assets/fonts/merriweather-light&#x27;, normal);
@include font-face(&quot;Merriweather&quot;, &#x27;/assets/fonts/merriweather-lightitalic&#x27;, normal, italic);
@include font-face(&quot;Merriweather&quot;, &#x27;/assets/fonts/merriweather-bold&#x27;, bold);
@include font-face(&quot;Merriweather&quot;, &#x27;/assets/fonts/merriweather-bolditalic&#x27;, bold, italic);
// Lato
@include font-face(&quot;Lato&quot;, &#x27;/assets/fonts/lato-light&#x27;, normal);
@include font-face(&quot;Lato&quot;, &#x27;/assets/fonts/lato-lightitalic&#x27;, normal, italic);
@include font-face(&quot;Lato&quot;, &#x27;/assets/fonts/lato-black&#x27;, bold);
@include font-face(&quot;Lato&quot;, &#x27;/assets/fonts/lato-blackitalic&#x27;, bold, italic);
// League Spartan
@include font-face(&quot;League Spartan&quot;, &#x27;/assets/fonts/leaguespartan-bold&#x27;, bold);

/* Globals
 * ---------------------------------------------------------------------------*/
@import &quot;base/base&quot;;


/* Layouts
 * ---------------------------------------------------------------------------*/
&#x60;&#x60;&#x60;

### Comments

---
> from: [**egardner**](https://github.com/jekyll/jekyll-help/issues/217#issuecomment-67997380) on: **Tuesday, December 23, 2014**

Okay, actually this was not the issue at all!
Turns out I had added most of my vendor Sass libraries to my &#x60;.gitignore&#x60; on the assumption they would be installed locally. But of course GH pages doesn&#x27;t work this way, all the dependencies need to be included. Somewhat misleading error message but otherwise everything works properly.

---
> from: [**troyswanson**](https://github.com/jekyll/jekyll-help/issues/217#issuecomment-68660650) on: **Sunday, January 04, 2015**

:boom: Glad you got it sorted.
