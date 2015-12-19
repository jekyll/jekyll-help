# [Combine/uglify javascript? No plugins …](https://github.com/jekyll/jekyll-help/issues/162)

> state: **closed** opened by: **** on: ****

Hi,

I&#x27;m using the &#x60;gh-pages&#x60; gem.

I&#x27;m finally exploring the SCSS and Cofeescript capabilities of Jekyll (been using Grunt for this up until now).

SCSS workflow is great, as I can use Sass to combine/minify my CSS (without plugins).

I&#x27;ve not used Coffeescript much, but it&#x27;s looking like it doesn&#x27;t have a feature to combine/uglify multiple JS files into one, no?

The best solution I&#x27;ve seen is to put my JS files in the &#x60;_includes&#x60; folder and use &#x60;{% include reset.css %}&#x60; (for example) to combine JS files into one (no uglification though).

### Questions:

1. What&#x27;s the best workflow to combine and uglify JS files use Jekyll without plugins (i.e., &#x60;gh-pages&#x60; gem compatible solution)?
1. Are there future plans to allow Jekyll to combine and uglify JS files like Sass does with CSS?

Thanks!


### Comments

---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/162#issuecomment-57895521) on: **Friday, October 03, 2014**

&gt; What&#x27;s the best workflow to combine and uglify JS files use Jekyll without plugins (i.e., &#x60;gh-pages&#x60; gem compatible solution)?

Ah, this comment sounds promising:

[Support a new &#x60;relative_include&#x60; tag #2870](https://github.com/jekyll/jekyll/pull/2870#issuecomment-54758772)

I guess I wouldn&#x27;t mind using the &#x60;include&#x60; technique if I could relatively include JS files from my assets folder (vs. having to call JS files from &#x60;_includes&#x60;).

Now I just have to wait several months until GitHub updates Jekyll! :laughing: 

I&#x27;d still like to see uglification of JS though. Anyone know if JS uglification is on the roadmap for &#x60;gh-pages&#x60; non-plugin workflow?
---
> from: [**mhulse**](https://github.com/jekyll/jekyll-help/issues/162#issuecomment-57896844) on: **Saturday, October 04, 2014**

&gt; Now I just have to wait several months until GitHub updates Jekyll! :laughing:

Looks like relative includes already are supported in the &#x60;gh-pages&#x60; gem (thanks again @gjtorikian).

It works via &#x60;gh-pages&#x60; v28. :+1: 

For example, here&#x27;s &#x60;site.coffee&#x60;:

&#x60;&#x60;&#x60;coffee
---
---

{% include_relative _partials/foo.js %}
{% include_relative _partials/bar.js %}

number = 42
&#x60;&#x60;&#x60;

And here&#x27;s &#x60;foo.js&#x60;:

&#x60;&#x60;&#x60;coffee
console.log &#x27;foo&#x27;
&#x60;&#x60;&#x60;

… and &#x60;bar.js&#x60;:

&#x60;&#x60;&#x60;coffee
console.log &#x27;bar&#x27;
&#x60;&#x60;&#x60;

Here&#x27;s the combined and preprocessed &#x60;site.js&#x60;:

&#x60;&#x60;&#x60;js
(function() {
  var number;

  console.log(&#x27;foo&#x27;);

  console.log(&#x27;bar&#x27;);

  number = 42;

}).call(this);
&#x60;&#x60;&#x60;

Cool!!!!

Now, if only there was a non-plugin way to uglify JS code … Perhaps liquid has something I can utilize?
