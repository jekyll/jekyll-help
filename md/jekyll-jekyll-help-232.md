# [A file was included in &#x60;_layouts/default.html&#x60; that is a symlink or does not exist in your &#x60;_includes&#x60; directory](https://github.com/jekyll/jekyll-help/issues/232)

> state: **open** opened by: **** on: ****

Hi,

Getting an error when I try to deploy to github pages:

A file was included in &#x60;_layouts/default.html&#x60; that is a symlink or does not exist in your &#x60;_includes&#x60; directory.

I can&#x27;t get around this error after many hours.  The source for my site is on github ( https://github.com/russt/russt.github.io ) and the destination is the default ( http://russt.github.io/ ).  I am just generating the default site using &quot;jekyll new&quot; with a few changes to the _config.yaml file.  I have tried both 2.4.0 and 2.5.2 versions of jekyll

Site looks fine locally, so I assume this has to do with how github is generating the site.

Any clues?  I realize that this could be a github issue, but I want to make sure that I&#x27;m not missing something basic.

thanks!
-Russ


### Comments

---
> from: [**russt**](https://github.com/jekyll/jekyll-help/issues/232#issuecomment-69103295) on: **Wednesday, January 07, 2015**

okay, this is a good one - I routinely git ignore &quot;foo&quot; files, but in this case it matched the footer.html pattern.  so indeed the _include file really was missing.

too bad the error message is not more specific about which file is missing!

