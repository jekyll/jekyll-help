# [&quot;undefined method &#x27;frontmatter_defaults&#x27;&quot; ](https://github.com/jekyll/jekyll-help/issues/275)

> state: **closed** opened by: **** on: ****

Hi!

I&#x27;m trying to build the jekyll-based manual for a program called Ardour. [Here is the github repo for the manual](https://github.com/Ardour/manual). 

I&#x27;m using jekyll 2.5.3 installed via gem on an Ubuntu 14.10 machine. However, I&#x27;m running into various problems. First, it couldn&#x27;t find the plugins even though they&#x27;re in the suppost default directory &quot;_plugins&quot;; I had to set it explicitly. Then I ran into a problem &quot;undefined method &#x60;frontmatter_defaults&#x27; for nil:NilClass&quot; - please see [this bug report](http://tracker.ardour.org/view.php?id=6167) at the Ardour bug tracker for details. Any clues what&#x27;s going on?  

### Comments

---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/275#issuecomment-75461072) on: **Sunday, February 22, 2015**

Regarding your first question... by default, jekyll expects to find your &#x60;_plugins&#x60; folder inside your &#x60;source&#x60; folder (alongside &#x60;_layouts&#x60;, &#x60;_posts&#x60;, etc).
---
> from: [**skagedal**](https://github.com/jekyll/jekyll-help/issues/275#issuecomment-75461463) on: **Sunday, February 22, 2015**

Allright, thanks! So that problem makes sense then. 
---
> from: [**willnorris**](https://github.com/jekyll/jekyll-help/issues/275#issuecomment-75467546) on: **Sunday, February 22, 2015**

So, as alluded to in the error message, &#x60;@site&#x60; is nil when called [here](https://github.com/Ardour/manual/blob/0b6ad9326ccdaec578164df288c8ce6321ca67a0/_plugins/manual.rb#L72), which is ultimately what&#x27;s throwing the error.  This &#x60;manual.rb&#x60; plugin is really convoluted... it would probably be much better to set these up as standard jekyll pages and come up with a simpler approach to build the TOC (there are a number of TOC plugins out there)
---
> from: [**skagedal**](https://github.com/jekyll/jekyll-help/issues/275#issuecomment-75534923) on: **Monday, February 23, 2015**

Thank you, Will! This helped me craft up a hack (see [Ardour bug report](http://tracker.ardour.org/view.php?id=6167) in case you would be interested). Now of course the next problem surfaced... :) And yes, it seems to me as well there would be simpler ways of doing this closer to the standard jekyll scenario, thanks for confirming that. The TOC plugins I can find when googling however is for building a TOC of a single page, not for a hierarchy of pages.... But anyway, I guess I can close this issue. 

To think out loud, my choices at this point seem to be: 1. try to see if I can get the same old version of jekyll as the Ardour devs are using; 2. keeping trying with some ugly hacks; 3. actually learn Ruby. :) 
